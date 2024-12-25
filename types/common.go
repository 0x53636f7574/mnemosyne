package types

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type UInt256 = *big.Int

type Address = common.Address

type ABI = abi.ABI

type Message = ethereum.CallMsg
