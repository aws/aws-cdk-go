package awsstepfunctionstasks


// Determines the level of detail about provisioned throughput consumption that is returned.
type DynamoConsumedCapacity string

const (
	// The response includes the aggregate ConsumedCapacity for the operation, together with ConsumedCapacity for each table and secondary index that was accessed.
	DynamoConsumedCapacity_INDEXES DynamoConsumedCapacity = "INDEXES"
	// The response includes only the aggregate ConsumedCapacity for the operation.
	DynamoConsumedCapacity_TOTAL DynamoConsumedCapacity = "TOTAL"
	// No ConsumedCapacity details are included in the response.
	DynamoConsumedCapacity_NONE DynamoConsumedCapacity = "NONE"
)

