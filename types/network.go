package types

import "time"

type Network struct {
	Version struct {
		Version string `json:"version"`
		Build   string `json:"build"`
	} `json:"version"`
	ChainId               string `json:"chain_id"`
	ProtocolVersion       int    `json:"protocol_version"`
	LatestProtocolVersion int    `json:"latest_protocol_version"`
	RpcAddr               string `json:"rpc_addr"`
	Validators            []struct {
		AccountId string `json:"account_id"`
		IsSlashed bool   `json:"is_slashed"`
	} `json:"validators"`
	SyncInfo struct {
		LatestBlockHash   string    `json:"latest_block_hash"`
		LatestBlockHeight int       `json:"latest_block_height"`
		LatestStateRoot   string    `json:"latest_state_root"`
		LatestBlockTime   time.Time `json:"latest_block_time"`
		Syncing           bool      `json:"syncing"`
	} `json:"sync_info"`
	ValidatorAccountId string `json:"validator_account_id"`
}

type NetworkInfo struct {
	ActivePeers []struct {
		Id        string      `json:"id"`
		Addr      string      `json:"addr"`
		AccountId interface{} `json:"account_id"`
	} `json:"active_peers"`
	NumActivePeers      int `json:"num_active_peers"`
	PeerMaxCount        int `json:"peer_max_count"`
	SentBytesPerSec     int `json:"sent_bytes_per_sec"`
	ReceivedBytesPerSec int `json:"received_bytes_per_sec"`
	KnownProducers      []struct {
		AccountId string      `json:"account_id"`
		Addr      interface{} `json:"addr"`
		PeerId    string      `json:"peer_id"`
	} `json:"known_producers"`
}

type Validators struct {
	CurrentValidators []struct {
		AccountId         string `json:"account_id"`
		PublicKey         string `json:"public_key"`
		IsSlashed         bool   `json:"is_slashed"`
		Stake             string `json:"stake"`
		Shards            []int  `json:"shards"`
		NumProducedBlocks int    `json:"num_produced_blocks"`
		NumExpectedBlocks int    `json:"num_expected_blocks"`
	} `json:"current_validators"`
	NextValidators []struct {
		AccountId string `json:"account_id"`
		PublicKey string `json:"public_key"`
		Stake     string `json:"stake"`
		Shards    []int  `json:"shards"`
	} `json:"next_validators"`
	CurrentFishermen []struct {
		AccountId string `json:"account_id"`
		PublicKey string `json:"public_key"`
		Stake     string `json:"stake"`
	} `json:"current_fishermen"`
	NextFishermen []struct {
		AccountId string `json:"account_id"`
		PublicKey string `json:"public_key"`
		Stake     string `json:"stake"`
	} `json:"next_fishermen"`
	CurrentProposals []struct {
		AccountId string `json:"account_id"`
		PublicKey string `json:"public_key"`
		Stake     string `json:"stake"`
	} `json:"current_proposals"`
	PrevEpochKickout []interface{} `json:"prev_epoch_kickout"`
	EpochStartHeight int           `json:"epoch_start_height"`
	EpochHeight      int           `json:"epoch_height"`
}