package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
)

// ERC20Token implementation
type ERC20Token struct {
	address  Address
	name     string
	symbol   string
	decimals uint8
}

func NewERC20Token(address string, name string, symbol string, decimals uint8) (*ERC20Token, error) {
	if !common.IsHexAddress(address) {
		return nil, errors.New("Invalid address: " + address)
	}

	return &ERC20Token{
		address:  common.HexToAddress(address),
		name:     name,
		symbol:   symbol,
		decimals: decimals,
	}, nil
}

func (token ERC20Token) GetAddress() common.Address {
	return token.address
}

func (token ERC20Token) GetName() string {
	return token.name
}

func (token ERC20Token) GetSymbol() string {
	return token.symbol
}

func (token ERC20Token) GetDecimals() uint8 {
	return token.decimals
}
