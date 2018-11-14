// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package primitives

// this package collects all reputation calculation related information,
// then calculates the reputations of candidates.

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/consensus/dpor/backend"
	contract2 "bitbucket.org/cpchain/chain/contracts/dpor/contracts/campaign"
	"bitbucket.org/cpchain/chain/contracts/pdash/sol"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --sol contracts/primitive_contracts_inst.sol --pkg contracts --out contracts/primitive_contracts_inst.go

var (
	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal
)

const (
	Created = iota
	SellerConfirmed
	ProxyFetched
	ProxyDelivered
	BuyerConfirmed
	Finished
	SellerRated
	BuyerRated
	AllRated
	Disputed
	Withdrawn
)

// CollectorConfig is the config of rpt info collector
type CollectorConfig struct {
	Client      bind.ContractBackend
	ChainConfig *configs.ChainConfig
	DporConfig  *configs.DporConfig
}

type RptPrimitiveBackend interface {
	// Rank returns the rank for given account address at the given block number.
	Rank(address common.Address, number uint64) (int64, error)

	// TxVolume returns the transaction volumn for given account address at the given block number.
	TxVolume(address common.Address, number uint64) (int64, error)

	// Maintenance returns the maintenance score for given account address at the given block number.
	Maintenance(address common.Address, number uint64) (int64, error)

	// UploadCount returns the upload score for given account address at the given block number.
	UploadCount(address common.Address, number uint64) (int64, error)

	// ProxyInfo returns a value indicating whether the given address is proxy and the count of transactions processed
	// by the proxy at the given block number.
	ProxyInfo(address common.Address, number uint64) (isProxy int64, proxyCount int64, err error)
}

type RptEvaluator struct {
	backend.ChainBackend
	Config CollectorConfig
}

func NewCollectorConfig(Client backend.ChainBackend, config *CollectorConfig) (*RptEvaluator, error) {
	bc := &RptEvaluator{
		ChainBackend: Client,
		Config:       *config,
	}
	return bc, nil
}

// GetCoinAge is the func to get rank to rpt
func (re *RptEvaluator) Rank(address common.Address, number uint64) (int64, error) {
	var balances []float64
	myBalance, err := re.BalanceAt(context.Background(), address, big.NewInt(int64(number)))
	if err != nil {
		log.Warn("error with getReputationnode", "error", err)
	}
	contractAddress := re.Config.DporConfig.Contracts["campagin"]
	intance, err := contract2.NewCampaign(contractAddress, re.Config.Client)
	if err != nil {
		log.Warn("NewCampaign error", address, err)
	}
	rNodeAddress, err := intance.CandidatesOf(nil, big.NewInt(int64(number)))
	if err != nil {
		log.Warn("CandidatesOf error", address, err)
	}
	for _, committee := range rNodeAddress {
		balance, err := re.BalanceAt(context.Background(), committee, big.NewInt(int64(number)))
		if err != nil {
			log.Warn("error with bc.BalanceAt", "error", err)
			return 0, err
		}
		balances = append(balances, float64(balance.Uint64()))
	}
	var rank int64
	sort.Sort(sort.Reverse(sort.Float64Slice(balances)))
	index := sort.SearchFloat64s(balances, float64(myBalance.Uint64()))
	rank = int64(index / 21)
	return rank, err
}

// GetCoinAge is the func to get txVolume to rpt
func (re *RptEvaluator) TxVolume(address common.Address, number uint64) (int64, error) {
	block, err := re.BlockByNumber(context.Background(), big.NewInt(int64(number)))
	if err != nil {
		log.Warn("error with bc.getTxVolume", "error", err)
		return 0, err
	}
	txvs := int64(0)
	signer := types.NewPrivTxSupportEIP155Signer(re.Config.ChainConfig.ChainID)
	txs := block.Transactions()
	for _, tx := range txs {
		sender, err := signer.Sender(tx)
		if err != nil {
			continue
		}
		if sender == address {
			//		txvs += float64(tx.Value().Uint64())
			txvs += 1
		}
	}
	return txvs, nil
}

// leader:0,committee:1,rNode:2,nil:3
func (re *RptEvaluator) Maintenance(address common.Address, number uint64) (int64, error) {
	ld := int64(2)
	ifRnode, err := re.RNode(address, number)
	if ifRnode != true {
		return 3, nil
	}
	if re.Config.ChainConfig.ChainID.Uint64() == uint64(4) {
		return 0, nil
	}
	header, err := re.HeaderByNumber(context.Background(), big.NewInt(int64(number)))
	if err != nil {
		log.Warn("error with bc.getIfLeader", "error", err)
		return 0, err
	}
	number = number%re.Config.DporConfig.Epoch - 1
	leaderBytes := header.Extra[uint64(extraVanity)+number*common.AddressLength : uint64(extraVanity)+(number+1)*common.AddressLength]
	leader := common.BytesToAddress(leaderBytes)

	fmt.Println("leader.Hex():", leader.Hex())

	if leader == address {
		ld = 0
	} else {
		for _, committe := range re.CommitteeMember(header) {
			if address == committe {
				ld = 1
			}
		}
	}
	return ld, nil
}

// GetCoinAge is the func to get uploadnamber to rpt
func (re *RptEvaluator) UploadCount(address common.Address, number uint64) (int64, error) {
	uploadNumber := int64(0)
	contractAddress := re.Config.DporConfig.Contracts["register"]
	upload, err := pdash.NewRegister(contractAddress, re.Config.Client)
	if err != nil {
		log.Warn("NewRegister error", address, err)
		return uploadNumber, err
	}
	fileNumber, err := upload.GetUploadCount(nil, address)
	if err != nil {
		log.Warn("GetUploadCount error", address, err)
		return uploadNumber, err
	}
	return fileNumber.Int64(), err
}

func (re *RptEvaluator) ProxyInfo(address common.Address, number uint64) (isProxy int64, proxyCount int64, err error) {
	proxyCount = int64(0)
	isProxy = int64(0)
	var proxyAddresses []common.Address
	contractAddress := re.Config.DporConfig.Contracts["pdash"]
	pdash, err := pdash.NewPdash(contractAddress, re.Config.Client)

	if err != nil {
		log.Warn("NewPdash error", address, err)
		return proxyCount, 0, err
	}

	len, err := pdash.BlockOrdersLength(nil, big.NewInt(int64(number)))
	if err != nil {
		log.Warn("BlockOrdersLength err", address, err)
		return proxyCount, 0, err
	}

	for i := 0; i < int(len.Int64()); i++ {
		id, err := pdash.BlockOrders(nil, big.NewInt(int64(number)), big.NewInt(int64(i)))
		if err != nil {
			log.Warn("BlockOrders error", address, err)
			break
		}
		OrderRecord, err := pdash.OrderRecords(nil, id)
		proxyAddresses = append(proxyAddresses, OrderRecord.ProxyAddress)
	}

	for _, proxyAddress := range proxyAddresses {
		if proxyAddress == address {
			isProxy = 1
			break
		}
	}

	for i := 0; i < int(len.Int64()); i++ {
		id, err := pdash.BlockOrders(nil, big.NewInt(int64(number)), big.NewInt(int64(i)))
		if err != nil {
			log.Warn("BlockOrders error", address, err)
			break
		}
		OrderRecord, err := pdash.OrderRecords(nil, id)
		if OrderRecord.ProxyAddress == address && OrderRecord.State == Finished {
			proxyCount += 1
			if proxyCount == 100 {
				break
			}
		}
	}

	return isProxy, proxyCount, err
}

func (re *RptEvaluator) CommitteeMember(header *types.Header) []common.Address {
	committee := make([]common.Address, (len(header.Extra)-extraVanity-extraSeal)/common.AddressLength)
	for i := 0; i < len(committee); i++ {
		copy(committee[i][:], header.Extra[extraVanity+i*common.AddressLength:extraVanity+(i+1)*common.AddressLength])
	}
	return committee
}

func (re *RptEvaluator) RNode(address common.Address, number uint64) (bool, error) {
	contractAddress := re.Config.DporConfig.Contracts["campagin"]
	instance, err := contract2.NewCampaign(contractAddress, re.Config.Client)
	if err != nil {
		log.Warn("NewCampaign error", address, err)
		return false, err
	}
	rNdoeAddress, err := instance.CandidatesOf(nil, big.NewInt(int64(number)))
	if err != nil {
		log.Warn("CandidatesOf error", address, err)
		return false, err
	}
	for _, rNode := range rNdoeAddress {
		if rNode == address {
			return true, nil
		}
	}
	return false, nil
}
