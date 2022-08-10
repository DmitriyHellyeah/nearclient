package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) ViewAccount(ctx context.Context, accountId string) (resp types.Account, err error) {

	params := map[string]interface{}{
		"request_type": "view_account",
		"finality": "final",
		"account_id": []string{accountId},
	}
	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewAccountChanges(ctx context.Context, accountIds []string, blockId uint64) (resp types.AccountChanges, err error) {

	params := map[string]interface{}{
		"changes_type": "account_changes",
		"account_ids": accountIds,
		"block_id": blockId,
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

