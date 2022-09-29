package awsdynamodb


// Supported DynamoDB table operations.
// Experimental.
type Operation string

const (
	// GetItem.
	// Experimental.
	Operation_GET_ITEM Operation = "GET_ITEM"
	// BatchGetItem.
	// Experimental.
	Operation_BATCH_GET_ITEM Operation = "BATCH_GET_ITEM"
	// Scan.
	// Experimental.
	Operation_SCAN Operation = "SCAN"
	// Query.
	// Experimental.
	Operation_QUERY Operation = "QUERY"
	// GetRecords.
	// Experimental.
	Operation_GET_RECORDS Operation = "GET_RECORDS"
	// PutItem.
	// Experimental.
	Operation_PUT_ITEM Operation = "PUT_ITEM"
	// DeleteItem.
	// Experimental.
	Operation_DELETE_ITEM Operation = "DELETE_ITEM"
	// UpdateItem.
	// Experimental.
	Operation_UPDATE_ITEM Operation = "UPDATE_ITEM"
	// BatchWriteItem.
	// Experimental.
	Operation_BATCH_WRITE_ITEM Operation = "BATCH_WRITE_ITEM"
	// TransactWriteItems.
	// Experimental.
	Operation_TRANSACT_WRITE_ITEMS Operation = "TRANSACT_WRITE_ITEMS"
	// TransactGetItems.
	// Experimental.
	Operation_TRANSACT_GET_ITEMS Operation = "TRANSACT_GET_ITEMS"
	// ExecuteTransaction.
	// Experimental.
	Operation_EXECUTE_TRANSACTION Operation = "EXECUTE_TRANSACTION"
	// BatchExecuteStatement.
	// Experimental.
	Operation_BATCH_EXECUTE_STATEMENT Operation = "BATCH_EXECUTE_STATEMENT"
	// ExecuteStatement.
	// Experimental.
	Operation_EXECUTE_STATEMENT Operation = "EXECUTE_STATEMENT"
)

