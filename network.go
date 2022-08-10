package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) Network(ctx context.Context) (resp types.Network, err error) {

	err = c.RPCClient.CallWithDecode("status", nil, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) NetworkInfo(ctx context.Context) (resp types.NetworkInfo, err error) {

	err = c.RPCClient.CallWithDecode("network_info", nil, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) Validators(ctx context.Context, params interface{}) (resp types.Validators, err error) {

	err = c.RPCClient.CallWithDecode("validators", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ValidatorsByBlockId(ctx context.Context, id uint64) (resp types.Validators, err error) {

	resp, err = c.Validators(context.Background(), []interface{}{id})

	if err != nil {
		return
	}
	return
}

func (c *Client) ValidatorsByBlockHash(ctx context.Context, hash string) (resp types.Validators, err error) {

	resp, err = c.Validators(context.Background(), []interface{}{hash})

	if err != nil {
		return
	}
	return
}
func (c *Client) ValidatorsNull(ctx context.Context) (resp types.Validators, err error) {

	resp, err = c.Validators(context.Background(), []interface{}{nil})

	if err != nil {
		return
	}
	return
}
