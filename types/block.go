package types

type BlockDetails struct {
	Author string        `json:"author"`
	Header BlockHeader   `json:"header"`
	Chunks []ChunkHeader `json:"chunks"`
}

type BlockChanges struct {
	BlockHash string `json:"block_hash"`
	Changes   []struct {
		Type      string `json:"type"`
		AccountId string `json:"account_id"`
	} `json:"changes"`
}

type BlockHeader struct {
	Height                uint64        `json:"height"`
	EpochId               string        `json:"epoch_id"`
	NextEpochId           string        `json:"next_epoch_id"`
	Hash                  string        `json:"hash"`
	PrevHash              string        `json:"prev_hash"`
	PrevStateRoot         string        `json:"prev_state_root"`
	ChunkReceiptsRoot     string        `json:"chunk_receipts_root"`
	ChunkHeadersRoot      string        `json:"chunk_headers_root"`
	ChunkTxRoot           string        `json:"chunk_tx_root"`
	OutcomeRoot           string        `json:"outcome_root"`
	ChunksIncluded        int           `json:"chunks_included"`
	ChallengesRoot        string        `json:"challenges_root"`
	Timestamp             int64         `json:"timestamp"`
	TimestampNanosec      string        `json:"timestamp_nanosec"`
	RandomValue           string        `json:"random_value"`
	ValidatorProposals    []interface{} `json:"validator_proposals"`
	ChunkMask             []bool        `json:"chunk_mask"`
	GasPrice              string        `json:"gas_price"`
	RentPaid              string        `json:"rent_paid"`
	ValidatorReward       string        `json:"validator_reward"`
	TotalSupply           string        `json:"total_supply"`
	ChallengesResult      []interface{} `json:"challenges_result"`
	LastFinalBlock        string        `json:"last_final_block"`
	LastDsFinalBlock      string        `json:"last_ds_final_block"`
	NextBpHash            string        `json:"next_bp_hash"`
	BlockMerkleRoot       string        `json:"block_merkle_root"`
	Approvals             []string      `json:"approvals"`
	Signature             string        `json:"signature"`
	LatestProtocolVersion int           `json:"latest_protocol_version"`
}

