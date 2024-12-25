package types

type MulticallEntry struct {
	Target   Address
	ABI      *ABI
	Method   string
	CallData []byte
}
