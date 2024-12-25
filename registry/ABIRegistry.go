package registry

import (
	"github.com/0x53636f7574/mnemosyne/types"
	"github.com/0x53636f7574/mnemosyne/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"strings"
)

type ABIRegistry struct {
	AbstractRegistry[string, *types.ABI]
}

var globalABIRegistry = (&ABIRegistry{
	AbstractRegistry[string, *types.ABI]{
		data: make(map[string]*types.ABI),
	},
}).Bootstrap()

func (registry *ABIRegistry) Bootstrap() *ABIRegistry {
	multicall3, _ := LoadABIFromFile(utils.BuildRelativePath("/abis/multicall3.json"))
	registry.Save("multicall3", multicall3)

	erc20, _ := LoadABIFromFile(utils.BuildRelativePath("/abis/ERC20.json"))
	registry.Save("ERC20", erc20)

	return registry
}

func SaveToGlobalABIRegistry(key string, abi *types.ABI) {
	globalABIRegistry.Save(key, abi)
}

func ExtractFromGlobalABIRegistry(key string) *types.ABI {
	return globalABIRegistry.Extract(key)
}

func SetGlobalABIRegistry(registry *ABIRegistry) {
	globalABIRegistry = registry
}

func LoadABIFromFile(filepath string) (*abi.ABI, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return LoadABIFromJSON(string(data))
}

func LoadABIFromJSON(json string) (*abi.ABI, error) {
	loadedAbi, err := abi.JSON(strings.NewReader(string(json)))
	if err != nil {
		return nil, err
	}

	return &loadedAbi, err
}
