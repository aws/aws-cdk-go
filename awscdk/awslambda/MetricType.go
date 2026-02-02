package awslambda


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
//   		Metrics: []MetricType{
//   			lambda.MetricType_EVENT_COUNT,
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
	// Error Count metrics provide insights into invocation errors and failures in your event source mapping processing.
	MetricType_ERROR_COUNT MetricType = "ERROR_COUNT"
	// Kafka-specific metrics provide detailed insights into Kafka consumer behavior, including lag, throughput, and partition-specific metrics.
	MetricType_KAFKA_METRICS MetricType = "KAFKA_METRICS"
)

