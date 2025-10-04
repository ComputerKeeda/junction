package types

// Event types for the evmbridge module
const (
	// EventTypeTokensLocked is emitted when tokens are locked in the bridge
	EventTypeTokensLocked = "tokens_locked"
	// EventTypeTokensUnlocked is emitted when tokens are unlocked from the bridge
	EventTypeTokensUnlocked = "tokens_unlocked"
)

// Event attribute keys
const (
	// AttributeKeyCreator is the address that initiated the lock
	AttributeKeyCreator = "creator"
	// AttributeKeyAmount is the amount of tokens locked/unlocked
	AttributeKeyAmount = "amount"
	// AttributeKeyToAddress is the EVM address to mint tokens to
	AttributeKeyToAddress = "to_address"
	// AttributeKeyBlockHeight is the block height when operation occurred
	AttributeKeyBlockHeight = "block_height"
	// AttributeKeyTimestamp is the timestamp when operation occurred
	AttributeKeyTimestamp = "timestamp"
	// AttributeKeyRecipient is the recipient address for unlock operations
	AttributeKeyRecipient = "recipient"
	// AttributeKeyEVMTxHash is the EVM transaction hash that triggered the unlock
	AttributeKeyEVMTxHash = "evm_tx_hash"
	// AttributeKeyFromEvmAddress is the EVM address that triggered the unlock
	AttributeKeyFromEvmAddress = "from_evm_address"
)
