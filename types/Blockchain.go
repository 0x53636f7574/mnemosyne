package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// Blockchain implementation
type Blockchain struct {
	id         uint
	name       string
	multicall3 *Address
	rpcList    []IRPC
	isTestnet  bool
}

func NewBlockchain(id uint, name string, multicall3Contract string, isTestnet bool, rpcList []string) (*Blockchain, error) {
	multicall3Address := common.HexToAddress(multicall3Contract)
	chain := &Blockchain{
		id:         id,
		name:       name,
		multicall3: &multicall3Address,
		isTestnet:  isTestnet,
	}

	for _, rpc := range rpcList {
		err := chain.AddRPC(rpc)
		if err != nil {
			return nil, err
		}
	}
	return chain, nil
}

func (chain *Blockchain) GetID() uint {
	return chain.id
}

func (chain *Blockchain) GetName() string {
	return chain.name
}

func (chain *Blockchain) GetRPCList() RPCList {
	return chain.rpcList
}

func (chain *Blockchain) IsTestnet() bool {
	return chain.isTestnet
}

func (chain *Blockchain) AddRPC(url string) error {
	isHttp := strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
	isWS := strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")
	if !isHttp && !isWS {
		return errors.New("Invalid rpc url: " + url)
	}
	chain.rpcList = append(chain.rpcList, &RPC{isHTTP: isHttp, url: url})
	return nil
}

func (chain *Blockchain) GetMulticall3Address() *Address {
	return chain.multicall3
}
