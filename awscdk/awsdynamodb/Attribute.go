package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &AlarmProps{
//   	Metric: metric,
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(1),
//   })
//
type Attribute struct {
	// The name of an attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

