package nearclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DmitriyHellyeah/nearclient/types"
)

func (c *Client) ViewContractCode(ctx context.Context, accountId string) (resp types.ContractCode, err error) {
	params := map[string]interface{}{
		"request_type": "view_code",
		"finality": "final",
		"account_id": accountId,
	}
	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewContractState(ctx context.Context, accountId, prefixBase string) (resp types.ContractState, err error) {
	params := map[string]interface{}{
		"request_type": "view_state",
		"finality": "final",
		"account_id": accountId,
		"prefix_base64": prefixBase,
	}
	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewContractStateChanges(ctx context.Context, accountIds []string, prefixBase string, blockId uint64) (resp types.ContractStateChanges, err error) {
	params := map[string]interface{}{
		"changes_type": "data_changes",
		"account_ids": accountIds,
		"prefix_base64": prefixBase,
		"block_id": blockId,
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) ViewContractCodeChanges(ctx context.Context, accountIds []string, blockId uint64) (resp types.ContractCodeChanges, err error) {
	params := map[string]interface{}{
		"changes_type": "contract_code_changes",
		"account_ids": accountIds,
		"block_id": blockId,
	}
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_changes", params, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) CallContractFunction(ctx context.Context, accountId, contractMethod string, args interface{}) (resp types.CallContractFunction, err error) {
	if len(accountId) == 0 {
		err = errors.New("accountId can't be empty")
		return
	}
	if len(contractMethod) == 0 {
		err = errors.New("contractMethod can't be empty")
		return
	}
	b, err := json.Marshal(args)

	if err != nil {
		err = fmt.Errorf("can't encode args. %v", err)
		return
	}

	encodedArgs := base64.StdEncoding.EncodeToString(b)

	params := map[string]interface{}{
		"request_type": "call_function",
		"finality": "final",
		"account_id": accountId,
		"method_name": contractMethod,
		"args_base64": encodedArgs,
	}
	resp, err = c.CallContractFunctionRaw(context.Background(), params)

	if err != nil {
		return
	}

	return
}

func (c *Client) CallContractFunctionRaw(ctx context.Context, params interface{}) (resp types.CallContractFunction, err error) {
	err = c.RPCClient.CallWithDecode("query", params, &resp)

	if err != nil {
		return
	}

	return
}
