package types

type IERC20Token interface {
	IContract
	GetName() string
	GetSymbol() string
	GetDecimals() uint8
}
