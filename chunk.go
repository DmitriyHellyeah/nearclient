package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) ChunkDetails(ctx context.Context, hash string) (resp types.ChunkDetails, err error) {
	err = c.RPCClient.CallWithDecode("chunk", []string{hash}, &resp)

	if err != nil {
		return
	}

	return
}
