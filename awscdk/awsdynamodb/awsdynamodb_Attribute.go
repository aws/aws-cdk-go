package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   })
//
//   metric := table.metricThrottledRequestsForOperations(&operationsMetricOptions{
//   	operations: []operation{
//   		dynamodb.*operation_PUT_ITEM,
//   	},
//   	period: awscdk.Duration.minutes(jsii.Number(1)),
//   })
//
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &alarmProps{
//   	metric: metric,
//   	evaluationPeriods: jsii.Number(1),
//   	threshold: jsii.Number(1),
//   })
//
type Attribute struct {
	// The name of an attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of an attribute.
	Type AttributeType `field:"required" json:"type" yaml:"type"`
}

