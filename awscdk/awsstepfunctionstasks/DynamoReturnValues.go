package awsstepfunctionstasks


// Use ReturnValues if you want to get the item attributes as they appear before or after they are changed.
type DynamoReturnValues string

const (
	// Nothing is returned.
	DynamoReturnValues_NONE DynamoReturnValues = "NONE"
	// Returns all of the attributes of the item.
	DynamoReturnValues_ALL_OLD DynamoReturnValues = "ALL_OLD"
	// Returns only the updated attributes.
	DynamoReturnValues_UPDATED_OLD DynamoReturnValues = "UPDATED_OLD"
	// Returns all of the attributes of the item.
	DynamoReturnValues_ALL_NEW DynamoReturnValues = "ALL_NEW"
	// Returns only the updated attributes.
	DynamoReturnValues_UPDATED_NEW DynamoReturnValues = "UPDATED_NEW"
)

