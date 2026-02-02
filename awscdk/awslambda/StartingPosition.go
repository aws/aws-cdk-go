package awslambda


// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction Function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // Enable basic event and error metrics
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String("basic-monitoring"),
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	// Provisioned mode is required for observability features
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(2),
//   		MaximumPollers: jsii.Number(10),
//   	},
//   	MetricsConfig: &MetricsConfig{
//   		Metrics: []MetricType{
//   			lambda.MetricType_EVENT_COUNT,
//   			lambda.MetricType_ERROR_COUNT,
//   		},
//   	},
//   }))
//
type StartingPosition string

const (
	// Start reading at the last untrimmed record in the shard in the system, which is the oldest data record in the shard.
	StartingPosition_TRIM_HORIZON StartingPosition = "TRIM_HORIZON"
	// Start reading just after the most recent record in the shard, so that you always read the most recent data in the shard.
	StartingPosition_LATEST StartingPosition = "LATEST"
	// Start reading from a position defined by a time stamp.
	//
	// Only supported for Amazon Kinesis streams, otherwise an error will occur.
	// If supplied, `startingPositionTimestamp` must also be set.
	StartingPosition_AT_TIMESTAMP StartingPosition = "AT_TIMESTAMP"
)

