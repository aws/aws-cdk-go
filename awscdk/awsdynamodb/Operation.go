package awsdynamodb


// Supported DynamoDB table operations.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
//   metric := table.metricThrottledRequestsForOperations(&OperationsMetricOptions{
//   	Operations: []operation{
//   		dynamodb.*operation_PUT_ITEM,
//   	},
//   	Period: awscdk.Duration_Minutes(jsii.Number(1)),
//   })
//
//   cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
//   	Metric: metric,
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(1),
//   })
//
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

