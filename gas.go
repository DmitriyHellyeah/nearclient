package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) GasPrice(ctx context.Context, params interface{}) (resp types.Gas, err error) {

	err = c.RPCClient.CallWithDecode("gas_price", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) GasPriceByBlockId(ctx context.Context, id uint64) (resp types.Gas, err error) {

	resp, err = c.GasPrice(context.Background(), []interface{}{id})

	if err != nil {
		return
	}
	return
}

func (c *Client) GasPriceByBlockHash(ctx context.Context, hash string) (resp types.Gas, err error) {

	resp, err = c.GasPrice(context.Background(), []interface{}{hash})

	if err != nil {
		return
	}
	return
}

func (c *Client) GasPriceNull(ctx context.Context) (resp types.Gas, err error) {

	resp, err = c.GasPrice(context.Background(), []interface{}{nil})

	if err != nil {
		return
	}
	return
}
