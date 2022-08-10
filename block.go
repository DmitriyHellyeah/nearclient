package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) BlockDetails(ctx context.Context, params interface{}) (resp types.BlockDetails, err error) {
	err = c.RPCClient.CallWithDecode("block", params, &resp)

	if err != nil {
		return
	}

	return
}

func (c *Client) BlockDetailsFinality(ctx context.Context) (resp types.BlockDetails, err error) {
	resp, err = c.BlockDetails(ctx, map[string]interface{}{"finality": "final"})

	if err != nil {
		return
	}

	return
}

func (c *Client) BlockDetailsById(ctx context.Context, id uint64) (resp types.BlockDetails, err error) {
	return c.BlockDetails(ctx, map[string]interface{}{"block_id": id})
}


func (c *Client) BlockChanges(ctx context.Context, params interface{}) (resp types.BlockChanges, err error) {
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes_in_block", params, &resp)

	if err != nil {
		return
	}

	return
}

func (c *Client) BlockChangesById(ctx context.Context, id uint64) (resp types.BlockChanges, err error) {
	resp, err = c.BlockChanges(ctx, map[string]interface{}{"block_id": id})

	if err != nil {
		return
	}

	return
}

func (c *Client) BlockChangesFinality(ctx context.Context) (resp types.BlockChanges, err error) {
	resp, err = c.BlockChanges(ctx, map[string]interface{}{"finality": "final"})

	if err != nil {
		return
	}

	return
}

