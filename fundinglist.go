package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type FundingList struct {
	fname   string
	fp      *os.File
	scanner *bufio.Scanner
	lineidx uint
}

type FundingEntry struct {
	address *common.Address
	amount  *big.Int
}

func OpenFundingList(fname string) (*FundingList, error) {
	fp, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	return &FundingList{
		fname:   fname,
		fp:      fp,
		scanner: bufio.NewScanner(fp),
	}, nil
}

func (fl *FundingList) Close() {
	if fl.scanner == nil {
		return
	}

	fl.fp.Close()
	fl.scanner = nil
	fl.fp = nil
}

func (fl *FundingList) GetNextEntry() *FundingEntry {
	if fl.scanner == nil {
		return nil
	}

	for fl.scanner.Scan() {
		line := strings.Trim(fl.scanner.Text(), " \t")
		fl.lineidx++

		entry := parseEntry(line, fl.lineidx)
		if entry != nil {
			return entry
		}
	}

	fl.Close()
	return nil
}

func parseEntry(line string, lineidx uint) *FundingEntry {
	cmtpos := strings.IndexByte(line, '#')
	if cmtpos != -1 {
		line = line[:cmtpos]
	}
	if line == "" {
		return nil
	}

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		fmt.Errorf("Invalid format in funding list on line %v: %v\n", lineidx, line)
		return nil
	}

	addr := common.HexToAddress(parts[0])

	amountStr := parts[1]
	amountStr = strings.ReplaceAll(amountStr, " ", "")

	re := regexp.MustCompile(`(?i)ETH`)
	amountStr = re.ReplaceAllString(amountStr, "000000000000000000")

	re = regexp.MustCompile(`(?i)gwei`)
	amountStr = re.ReplaceAllString(amountStr, "000000000")

	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		fmt.Errorf("Invalid amount in funding list on line %v: %v\n", lineidx, amountStr)
		return nil
	}

	return &FundingEntry{
		address: &addr,
		amount:  amount,
	}
}
