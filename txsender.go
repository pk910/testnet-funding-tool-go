package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type TxSender struct {
	wallet       *WalletState
	verbose      bool
	pendingLimit uint
	txChan       chan []byte
	state        *TxSenderState
}

type TxSenderState struct {
	running       bool
	runMutex      sync.Mutex
	stateMutex    sync.Mutex
	queuedCount   uint
	pendingCount  uint
	pendingHashes *list.List
}

func StartTxSender(wallet *WalletState, pendingLimit uint) (*TxSender, error) {
	txSender := TxSender{
		wallet:       wallet,
		pendingLimit: pendingLimit,
		txChan:       make(chan []byte, pendingLimit),
		state: &TxSenderState{
			running:       true,
			pendingHashes: list.New(),
		},
	}
	go processTxQueue(&txSender)

	return &txSender, nil
}

func (self *TxSender) Stop() {
	self.state.stateMutex.Lock()
	self.state.running = false
	self.state.stateMutex.Unlock()
	self.state.runMutex.Lock()
	self.state.runMutex.Unlock()
}

func (self *TxSender) Await() {
	hasPending := true
	for hasPending {
		self.state.stateMutex.Lock()
		hasPending = self.state.running && (self.state.queuedCount > 0 || self.state.pendingCount > 0)
		self.state.stateMutex.Unlock()

		time.Sleep(1000 * time.Millisecond)
	}
}

func (self *TxSender) SubmitTx(txBytes []byte) {
	self.txChan <- txBytes
	self.state.stateMutex.Lock()
	self.state.queuedCount++
	self.state.stateMutex.Unlock()
}

func (self *TxSender) ProcessTx(txBytes []byte) {
	self.SubmitTx(txBytes)
}

func processTxQueue(txSender *TxSender) {
	txSender.state.runMutex.Lock()
	defer txSender.state.runMutex.Unlock()
	for {
		txSender.state.stateMutex.Lock()
		isRunning := txSender.state.running
		txSender.state.stateMutex.Unlock()
		if !isRunning {
			return
		}

		txSender.state.stateMutex.Lock()

		if txSender.state.pendingCount > 0 {
			// check if tx has been confirmed
			checkPendingTxs(txSender)
		}

		txSender.state.stateMutex.Unlock()

		waitAndAcceptTx(txSender)
	}
}

func processReceivedTx(txSender *TxSender, txBytes []byte) {
	txSender.state.stateMutex.Lock()
	defer txSender.state.stateMutex.Unlock()

	txSender.state.queuedCount--
	txHash := txSender.wallet.SubmitTransaction(txBytes)
	if txHash != nil {
		txSender.state.pendingCount++
		txSender.state.pendingHashes.PushBack(txHash.Bytes())
	}
}

func checkPendingTxs(txSender *TxSender) {
	for {
		pendingTx := txSender.state.pendingHashes.Front()
		if pendingTx == nil {
			return
		}
		txReceipt := txSender.wallet.GetTransactionReceipt(pendingTx.Value.([]byte))
		if txReceipt == nil {
			return
		}
		if txSender.verbose {
			fmt.Printf("      tx confirmed: 0x%v\n", common.Bytes2Hex(pendingTx.Value.([]byte)))
		}
		txSender.state.pendingHashes.Remove(pendingTx)
		txSender.state.pendingCount--
	}
}

func waitAndAcceptTx(txSender *TxSender) {
	start := time.Now()
	var timeout time.Duration

	for {
		txSender.state.stateMutex.Lock()
		pendingCount := txSender.state.pendingCount
		if txSender.state.pendingCount > 0 || txSender.state.queuedCount > 0 {
			timeout = 6000 * time.Millisecond
		} else {
			timeout = 30000 * time.Millisecond
		}
		txSender.state.stateMutex.Unlock()

		now := time.Now()
		waitTime := start.Add(timeout).Sub(now)
		if waitTime < 0 {
			return
		}

		if pendingCount < txSender.pendingLimit {
			select {
			case txBytes := <-txSender.txChan:
				processReceivedTx(txSender, txBytes)
			case <-time.After(timeout):
				return
			}
		} else {
			time.Sleep(waitTime)
			return
		}

	}

}
