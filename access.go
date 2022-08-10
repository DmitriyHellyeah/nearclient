package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) ViewAccessKey(ctx context.Context, params interface{}) (resp types.ViewAccessKey, err error) {

	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyByBlockId(ctx context.Context, accountId, publicKey string, blockId uint64) (resp types.ViewAccessKey, err error) {
	params := map[string]interface{}{
		"request_type": "view_access_key",
		"block_id": blockId,
		"account_id": accountId,
		"public_key": publicKey,
	}
	resp, err = c.ViewAccessKey(context.Background(), params)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyFinality(ctx context.Context, accountId, publicKey string) (resp types.ViewAccessKey, err error) {
	params := map[string]interface{}{
		"request_type": "view_access_key",
		"finality": "final",
		"account_id": accountId,
		"public_key": publicKey,
	}
	resp, err = c.ViewAccessKey(context.Background(), params)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyList(ctx context.Context, accountId string) (resp types.ViewAccessKey, err error) {
	params := map[string]interface{}{
		"request_type": "view_access_key_list",
		"finality": "final",
		"account_id": accountId,
	}
	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyChangesSingle(ctx context.Context, accountId, publicKey string) (resp types.ViewAccessKey, err error) {

	params := map[string]interface{}{
		"changes_type": "single_access_key_changes",
		"keys": []interface{}{
			map[string]interface{}{
				"account_id": accountId,
				"public_key": publicKey,
			},
		},
		"finality": "final",
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyChangesAll(ctx context.Context, params interface{}) (resp types.ViewAccessKey, err error) {
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyChangesAllByBlockId(ctx context.Context, accountIds []string, blockId uint64) (resp types.ViewAccessKey, err error) {

	params := map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids": accountIds,
		"block_id": blockId,
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccessKeyChangesAllFinality(ctx context.Context, accountIds []string) (resp types.ViewAccessKey, err error) {

	params := map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids": accountIds,
		"finality": "final",
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}


