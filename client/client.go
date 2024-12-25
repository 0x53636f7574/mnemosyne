package client

import (
	"context"
	"errors"
	"github.com/0x53636f7574/mnemosyne/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
	"time"
)

type internalOculusStorage struct {
	chainClientMap          map[uint]*Client
	chainClientCallbackList []func()
}

var storage *internalOculusStorage = &internalOculusStorage{
	chainClientMap:          map[uint]*Client{},
	chainClientCallbackList: []func(){},
}

type Client struct {
	rpcList            types.RPCList
	rpcClientMap       map[string]*ethclient.Client
	rpcClientCallbacks []func()
	fallbackDelay      uint
	maxRetries         uint
	rpcPointer         int
}

func NewClient(rpcList types.RPCList) *Client {
	return &Client{rpcList, make(map[string]*ethclient.Client), []func(){}, 0, 0, 0}
}

func CreateChainClient(chain types.IChain) (*Client, error) {
	if chain == nil {
		return nil, errors.New("blockchain is nil")
	}

	chainClient := storage.chainClientMap[chain.GetID()]
	if chainClient == nil {
		storage.chainClientMap[chain.GetID()] = NewClient(chain.GetRPCList())
		chainClient = storage.chainClientMap[chain.GetID()]
	}
	return chainClient, nil
}

func (client *Client) Close() error {
	for _, callback := range client.rpcClientCallbacks {
		callback()
	}

	return nil
}

func (client *Client) GetBlockNumber() (uint64, error) {
	remoteClient, err := client.getRPCClient()
	lastBlockNumber, err := executeWithRetries(
		func() (uint64, error) {
			return remoteClient.BlockNumber(context.TODO())
		},
		client.maxRetries,
		client.fallbackDelay,
	)

	if err != nil {
		return 0, err
	}
	return lastBlockNumber, nil
}

func (client *Client) GetFallbackDelay() uint {
	return client.fallbackDelay
}

func (client *Client) SetFallbackDelay(delay uint) {
	client.fallbackDelay = delay
}

func (client *Client) GetMaxRetries() uint {
	return client.maxRetries
}

func (client *Client) SetMaxRetries(maxRetries uint) {
	client.maxRetries = maxRetries
}

func (client *Client) getRPCClient() (*ethclient.Client, error) {
	if len(client.rpcList) == 0 {
		return nil, errors.New("no available RPC")
	}

	currentRPC := client.rpcList[client.rpcPointer]
	var rpcClient *ethclient.Client = nil
	var nextRPC types.IRPC = nil
	var err error = nil

	if client.rpcClientMap[currentRPC.GetURL()] == nil {
		rpcClient, err = ethclient.Dial(currentRPC.GetURL())
		if err == nil {
			return rpcClient, nil
		}
	}

	for {
		client.rpcPointer++
		if client.rpcPointer == len(client.rpcList) {
			client.rpcPointer = 0
		}

		nextRPC = client.rpcList[client.rpcPointer]
		if strings.EqualFold(nextRPC.GetURL(), currentRPC.GetURL()) {
			err = errors.New("couldn't establish connection with RPC")
		} else {
			rpcClient, err = ethclient.Dial(nextRPC.GetURL())
		}

		if rpcClient != nil {
			return rpcClient, err
		}
	}
}

func (client *Client) ContractCall(message *types.Message, blockNumber types.UInt256) ([]byte, error) {

	rpcClient, err := executeWithRetries(client.getRPCClient, client.maxRetries, client.fallbackDelay)

	if err != nil {
		return []byte{}, err
	}

	return executeWithRetries(
		func() ([]byte, error) {
			return rpcClient.CallContract(context.TODO(), *message, blockNumber)
		},
		client.maxRetries,
		client.fallbackDelay,
	)
}

func executeWithRetries[Target any](callable func() (Target, error), retries uint, delay uint) (Target, error) {
	result, err := callable()
	for try := 1; (try < int(retries)) && (err != nil); try++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		result, err = callable()
	}

	if err != nil {
		return result, err
	}
	return Target(result), err
}
