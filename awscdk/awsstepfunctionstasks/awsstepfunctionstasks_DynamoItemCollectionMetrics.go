package awsstepfunctionstasks


// Determines whether item collection metrics are returned.
type DynamoItemCollectionMetrics string

const (
	// If set to SIZE, the response includes statistics about item collections, if any, that were modified during the operation.
	DynamoItemCollectionMetrics_SIZE DynamoItemCollectionMetrics = "SIZE"
	// If set to NONE, no statistics are returned.
	DynamoItemCollectionMetrics_NONE DynamoItemCollectionMetrics = "NONE"
)

