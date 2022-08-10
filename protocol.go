package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) GenesisConfig(ctx context.Context) (resp types.Genesis, err error) {

	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_genesis_config", nil, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ProtocolConfig(ctx context.Context) (resp types.Genesis, err error) {

	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_protocol_config", map[string]interface{}{"finality":"final"}, &resp)

	if err != nil {
		return
	}
	return
}
