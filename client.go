package nearclient

import "github.com/DmitriyHellyeah/nearclient/jsonrpc"

type Client struct {
	RPCClient jsonrpc.RPCClient
	Config    NearConfig
}

func NewClient(config NearConfig) (client Client, err error) {
	client.RPCClient = jsonrpc.NewRPCClient(config.Host)
	client.Config = config
	if err != nil {
		return
	}

	return
}

func (c *Client) NetworkAddr() string {
	return c.Config.Host
}
