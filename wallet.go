package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletState struct {
	mutex   sync.Mutex
	privkey *ecdsa.PrivateKey
	address common.Address
	offline bool
	chainid uint64
	nonce   uint64
	client  *ethclient.Client
	balance *big.Int
}

func (state *WalletState) LoadPrivateKey(privkey string) error {
	var privateKey *ecdsa.PrivateKey
	if privkey == "" {
		var err error
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			return err
		}
	} else {
		var err error
		privateKey, err = crypto.HexToECDSA(privkey)
		if err != nil {
			return err
		}
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	state.privkey = privateKey
	state.address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return nil
}

func (state *WalletState) InitOfflineMode(chainid uint64, nonce uint64) {
	state.offline = true
	state.chainid = chainid
	state.nonce = nonce
}

func (state *WalletState) InitOnlineMode(rpchost string) error {
	client, err := ethclient.Dial(rpchost)
	if err != nil {
		return err
	}
	state.client = client

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	state.chainid = chainid.Uint64()

	nonce, err := client.PendingNonceAt(context.Background(), state.address)
	if err != nil {
		return err
	}
	state.nonce = nonce

	balance, err := client.PendingBalanceAt(context.Background(), state.address)
	if err != nil {
		return err
	}
	state.balance = balance

	return nil
}

func (state *WalletState) BuildEthTransaction(to *common.Address, value *big.Int, feeCap uint64, tipCap uint64, gasLimit uint64, data []byte) (*types.Transaction, error) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	chainId := new(big.Int).SetUint64(state.chainid)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     state.nonce,
		GasFeeCap: new(big.Int).SetUint64(feeCap),
		GasTipCap: new(big.Int).SetUint64(tipCap),
		Gas:       gasLimit,
		To:        to,
		Value:     value,
		Data:      data,
	})
	state.nonce++

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainId), state.privkey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (state *WalletState) BuildDeployTransaction(code []byte, feeCap uint64, tipCap uint64, gasLimit uint64) (*types.Transaction, error) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	chainId := new(big.Int).SetUint64(state.chainid)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     state.nonce,
		GasFeeCap: new(big.Int).SetUint64(feeCap),
		GasTipCap: new(big.Int).SetUint64(tipCap),
		Gas:       gasLimit,
		To:        nil,
		Value:     new(big.Int).SetUint64(0),
		Data:      code,
	})
	state.nonce++

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainId), state.privkey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (state *WalletState) SubmitTransaction(txBytes []byte) *common.Hash {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(txBytes)
	if err != nil {
		fmt.Printf("Error while decoding transaction: %v (%v)\n", err, len(txBytes))
		return nil
	}

	err = state.client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Printf("Error while sending transaction: %v\n", err)
		return nil
	}

	txHash := tx.Hash()
	fmt.Printf("    submitted transaction %v\n", txHash.String())

	return &txHash
}

func (state *WalletState) GetTransactionReceipt(txHash []byte) *types.Receipt {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	hash := common.Hash{}
	hash.SetBytes(txHash)
	receipt, err := state.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil
	}
	return receipt
}
