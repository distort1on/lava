package types

const (
	// ModuleName defines the module name
	ModuleName = "servicer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_servicer"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	BlockDeadlineForCallbackKey = "BlockDeadlineForCallback-value-"
)

const (
	UnstakingServicersAllSpecsKey      = "UnstakingServicersAllSpecs-value-"
	UnstakingServicersAllSpecsCountKey = "UnstakingServicersAllSpecs-count-"
)

const (
	CurrentSessionStartKey = "CurrentSessionStart-value-"
)

const (
	PreviousSessionBlocksKey = "PreviousSessionBlocks-value-"
)

const (
	EarliestSessionStartKey = "EarliestSessionStart-value-"
)