package types

// RPC implementation
type RPC struct {
	isHTTP bool
	url    string
}

func (rpc *RPC) GetURL() string {
	return rpc.url
}

func (rpc *RPC) IsHTTP() bool {
	return rpc.isHTTP
}

func (rpc *RPC) IsWS() bool {
	return !rpc.isHTTP
}
