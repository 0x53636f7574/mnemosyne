package registry

import (
	"github.com/0x53636f7574/mnemosyne/types"
)

var (
	ChainList = []types.IChain{
		MAINNET, BASE, BLAST, OPTIMISM, AVALANCHE, FANTOM,
	}

	MAINNET, _ = types.NewBlockchain(
		1,
		"Etherium",
		"0xca11bde05977b3631167028862be2a173976ca11",
		false,
		[]string{
			"https://eth.llamarpc.com",
		},
	)

	BASE, _ = types.NewBlockchain(
		8453,
		"Base",
		"0xca11bde05977b3631167028862be2a173976ca11",
		false,
		[]string{
			"https://mainnet.base.org",
		},
	)

	BLAST, _ = types.NewBlockchain(
		81457,
		"BLAST",
		"0xcA11bde05977b3631167028862bE2a173976CA11",
		false,
		[]string{
			"https://rpc.blast.io",
		},
	)

	OPTIMISM, _ = types.NewBlockchain(
		10,
		"Optimism",
		"0xca11bde05977b3631167028862be2a173976ca11",
		false,
		[]string{
			"https://rpc.ankr.com/bsc",
		},
	)

	AVALANCHE, _ = types.NewBlockchain(
		43114,
		"Avalanche",
		"0xca11bde05977b3631167028862be2a173976ca11",
		false,
		[]string{
			"https://api.avax.network/ext/bc/C/rpc",
		},
	)

	FANTOM, _ = types.NewBlockchain(
		250,
		"Fantom",
		"0xca11bde05977b3631167028862be2a173976ca11",
		false,
		[]string{
			"https://rpc.ankr.com/fantom",
		},
	)
)
