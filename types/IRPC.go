package types

type RPCList = []IRPC

type IRPC interface {
	GetURL() string
	IsHTTP() bool
	IsWS() bool
}
