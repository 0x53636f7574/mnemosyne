package core

import (
	"errors"
	chainClient "github.com/0x53636f7574/mnemosyne/client"
	"github.com/0x53636f7574/mnemosyne/registry"
	"github.com/0x53636f7574/mnemosyne/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

var (
	ZERO_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func ContractCall[T any](chain types.IChain, address, method string, contractAbi *abi.ABI, args ...any) (T, error) {
	var result T

	client, err := chainClient.CreateChainClient(chain)

	if err != nil {
		return result, err
	}

	data, err := client.ContractCall(
		makeContractCallMessage(address, method, contractAbi, args),
		nil,
	)

	if err != nil {
		return result, err
	}

	err = contractAbi.UnpackIntoInterface(&result, method, data)

	return result, err
}

func Multicall[T any](chain types.IChain, calls ...types.MulticallEntry) (*T, types.UInt256, error) {

	resultData := new(T)
	reflectedData := reflect.ValueOf(resultData).Elem()

	if chain.GetMulticall3Address() == nil {
		return nil, nil, errors.New("multicall is not supported")
	}

	if reflectedData.NumField() != len(calls) {
		return nil, nil, errors.New("signature of target object doesn't correspond to number of calls")
	}

	var multicallResult struct {
		BlockNumber types.UInt256
		ReturnData  [][]byte
	}

	client, err := chainClient.CreateChainClient(chain)

	if err != nil {
		return nil, nil, err
	}

	multicall3ABI := registry.ExtractFromGlobalABIRegistry("multicall3")

	response, err := client.ContractCall(
		makeMultiCallContractMessage(chain.GetMulticall3Address(), calls),
		nil,
	)

	err = multicall3ABI.UnpackIntoInterface(&multicallResult, "aggregate", response)
	if err != nil {
		return nil, nil, err
	}

	for index, call := range calls {
		if len(multicallResult.ReturnData[index]) == 0 {
			continue
		}
		decodeMulticallEntry(&call, reflectedData.Field(index).Addr().Interface(), multicallResult.ReturnData[index])
	}

	return resultData, multicallResult.BlockNumber, nil
}

func MakeMulticallEntry(address string, method string, abi *abi.ABI, args ...any) types.MulticallEntry {
	contractAddress := common.HexToAddress(address)
	var packedData []byte
	var err error
	if len(args) > 0 {
		packedData, err = abi.Pack(method, args...)
	} else {
		packedData, err = abi.Pack(method)
	}

	if err != nil {
		return types.MulticallEntry{}
	}

	return types.MulticallEntry{
		Target:   contractAddress,
		ABI:      abi,
		Method:   method,
		CallData: packedData,
	}
}

func makeContractCallMessage(address string, method string, abi *abi.ABI, args []any) *types.Message {
	contractAddress := common.HexToAddress(address)
	var packedData []byte
	var err error
	if len(args) > 0 {
		packedData, err = abi.Pack(method, args...)
	} else {
		packedData, err = abi.Pack(method)
	}

	if err != nil {
		return nil
	}

	return &ethereum.CallMsg{
		To:   &contractAddress,
		Data: packedData,
	}
}

func makeMultiCallContractMessage(address *types.Address, calls []types.MulticallEntry) *types.Message {
	multicall3ABI := registry.ExtractFromGlobalABIRegistry("multicall3")
	multicallPayload, err := multicall3ABI.Pack("aggregate", calls)
	if err != nil {
		return nil
	}

	return &ethereum.CallMsg{
		To:   address,
		Data: multicallPayload,
	}
}

func decodeMulticallEntry(entry *types.MulticallEntry, object any, data []byte) {
	err := entry.ABI.UnpackIntoInterface(object, entry.Method, data)
	if err != nil {
		panic(err.Error())
	}
}
