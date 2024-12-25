package types

type IChain interface {
	GetID() uint
	GetName() string
	GetMulticall3Address() *Address
	GetRPCList() RPCList
	IsTestnet() bool
}
