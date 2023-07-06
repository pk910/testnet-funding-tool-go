package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type OutputTxList struct {
	fname string
	fp    *os.File
}

func OpenOutputTxList(fname string) (*OutputTxList, error) {
	fp, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	return &OutputTxList{
		fname: fname,
		fp:    fp,
	}, nil
}

func (self *OutputTxList) Close() {
	if self.fp == nil {
		return
	}
	self.fp.Close()
	self.fp = nil
}

func (self *OutputTxList) AddTxEntry(txBytes []byte) {
	if self.fp == nil {
		return
	}

	self.fp.WriteString("0x")
	self.fp.WriteString(common.Bytes2Hex(txBytes))
	self.fp.WriteString("\n")
}

func (self *OutputTxList) ProcessTx(txBytes []byte) {
	self.AddTxEntry(txBytes)
}
