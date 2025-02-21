package awslambda


// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
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
type MetricType string

const (
	// Event Count metrics provide insights into the processing behavior of your event source mapping, including the number of events successfully processed, filtered out, or dropped.
	//
	// These metrics help you monitor the flow and status of events through your event source mapping.
	MetricType_EVENT_COUNT MetricType = "EVENT_COUNT"
)

