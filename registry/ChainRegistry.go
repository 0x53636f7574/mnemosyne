package registry

import "github.com/0x53636f7574/mnemosyne/types"

type ChainRegistry struct {
	AbstractRegistry[string, types.IChain]
}

var globalChainRegistry = (&ChainRegistry{
	AbstractRegistry[string, types.IChain]{
		data: make(map[string]types.IChain),
	},
}).Bootstrap()

func (registry *ChainRegistry) Bootstrap() *ChainRegistry {
	for _, chain := range ChainList {
		registry.Save(chain.GetName(), chain)
	}
	return registry
}

func SaveToGlobalChainRegistry(key string, abi types.IChain) {
	globalChainRegistry.Save(key, abi)
}

func ExtractFromGlobalChainRegistry(key string) types.IChain {
	return globalChainRegistry.Extract(key)
}

func SetGlobalChainRegistry(registry *ChainRegistry) {
	globalChainRegistry = registry
}
