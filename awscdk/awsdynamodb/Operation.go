package awsdynamodb


// Supported DynamoDB table operations.
type Operation string

const (
	// GetItem.
	Operation_GET_ITEM Operation = "GET_ITEM"
	// BatchGetItem.
	Operation_BATCH_GET_ITEM Operation = "BATCH_GET_ITEM"
	// Scan.
	Operation_SCAN Operation = "SCAN"
	// Query.
	Operation_QUERY Operation = "QUERY"
	// GetRecords.
	Operation_GET_RECORDS Operation = "GET_RECORDS"
	// PutItem.
	Operation_PUT_ITEM Operation = "PUT_ITEM"
	// DeleteItem.
	Operation_DELETE_ITEM Operation = "DELETE_ITEM"
	// UpdateItem.
	Operation_UPDATE_ITEM Operation = "UPDATE_ITEM"
	// BatchWriteItem.
	Operation_BATCH_WRITE_ITEM Operation = "BATCH_WRITE_ITEM"
	// TransactWriteItems.
	Operation_TRANSACT_WRITE_ITEMS Operation = "TRANSACT_WRITE_ITEMS"
	// TransactGetItems.
	Operation_TRANSACT_GET_ITEMS Operation = "TRANSACT_GET_ITEMS"
	// ExecuteTransaction.
	Operation_EXECUTE_TRANSACTION Operation = "EXECUTE_TRANSACTION"
	// BatchExecuteStatement.
	Operation_BATCH_EXECUTE_STATEMENT Operation = "BATCH_EXECUTE_STATEMENT"
	// ExecuteStatement.
	Operation_EXECUTE_STATEMENT Operation = "EXECUTE_STATEMENT"
)

