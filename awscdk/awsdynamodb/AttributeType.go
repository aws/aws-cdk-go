package awsdynamodb


// Data types for attributes within a table.
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
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes
//
type AttributeType string

const (
	// Up to 400KiB of binary data (which must be encoded as base64 before sending to DynamoDB).
	AttributeType_BINARY AttributeType = "BINARY"
	// Numeric values made of up to 38 digits (positive, negative or zero).
	AttributeType_NUMBER AttributeType = "NUMBER"
	// Up to 400KiB of UTF-8 encoded text.
	AttributeType_STRING AttributeType = "STRING"
)

