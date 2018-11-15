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

import (
	"math/big"

	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/configs"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --sol ../contracts/primitive_contracts.sol --pkg contracts --out ../contracts/primitive_contracts.go

// Definitions of primitive contracts

type GetRank struct {
	Backend RptPrimitiveBackend
}

func (c *GetRank) RequiredGas(input []byte) uint64 {
	return configs.GetRank
}

func (c *GetRank) Run(input []byte) ([]byte, error) {
	const getBalanceInputLength = 52

	if len(input) != getBalanceInputLength {
		log.Warnf("The expected input is %d bytes, but got %d", getBalanceInputLength, len(input))
		return common.LeftPadBytes(new(big.Int).Bytes(), 32), nil
	}

	addr := common.Bytes2Hex(input[0:20])
	number := new(big.Int).SetBytes(input[20:52])
	log.Infof("primitive_coinage, address %s, block number %d", addr, number)

	// TODO: @AC get cpchain Backend and read balance.
	coinAge, err := c.Backend.Rank(common.HexToAddress(addr), number.Uint64())
	if err != nil {
		log.Fatal("NewBasicCollector,error", err)
	}
	ret := new(big.Int).SetInt64(int64(coinAge))
	return common.LeftPadBytes(ret.Bytes(), 32), nil
}
