package awslambda


// Configuration for collecting metrics from the event source.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Stream: dynamodb.StreamViewType_NEW_IMAGE,
//   })
//
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	MetricsConfig: &MetricsConfig{
//   		Metrics: []eVENT_COUNT{
//   			lambda.MetricType_*eVENT_COUNT,
//   		},
//   	},
//   }))
//
type MetricsConfig struct {
	// List of metric types to enable for this event source.
	Metrics *[]MetricType `field:"required" json:"metrics" yaml:"metrics"`
}

