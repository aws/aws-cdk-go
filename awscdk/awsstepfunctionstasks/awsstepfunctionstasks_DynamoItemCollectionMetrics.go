package awsstepfunctionstasks


// Determines whether item collection metrics are returned.
// Experimental.
type DynamoItemCollectionMetrics string

const (
	// If set to SIZE, the response includes statistics about item collections, if any, that were modified during the operation.
	// Experimental.
	DynamoItemCollectionMetrics_SIZE DynamoItemCollectionMetrics = "SIZE"
	// If set to NONE, no statistics are returned.
	// Experimental.
	DynamoItemCollectionMetrics_NONE DynamoItemCollectionMetrics = "NONE"
)

