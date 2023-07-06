package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type InputTxList struct {
	fname   string
	fp      *os.File
	scanner *bufio.Scanner
	lineidx uint
}

func OpenInputTxList(fname string) (*InputTxList, error) {
	fp, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	return &InputTxList{
		fname:   fname,
		fp:      fp,
		scanner: bufio.NewScanner(fp),
	}, nil
}

func (fl *InputTxList) Close() {
	if fl.scanner == nil {
		return
	}

	fl.fp.Close()
	fl.scanner = nil
	fl.fp = nil
}

func (fl *InputTxList) GetNextEntry() []byte {
	if fl.scanner == nil {
		return nil
	}

	for fl.scanner.Scan() {
		line := strings.Trim(fl.scanner.Text(), " \t")
		fl.lineidx++

		if line[0:2] == "0x" {
			line = line[2:]
		}

		entry := common.Hex2Bytes(line)
		if entry != nil {
			return entry
		}
	}

	fl.Close()
	return nil
}
