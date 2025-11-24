package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// The set of properties for streaming event sources shared by Dynamo, Kinesis and Kafka.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseStreamEventSourceProps := &BaseStreamEventSourceProps{
//   	StartingPosition: awscdk.Aws_lambda.StartingPosition_TRIM_HORIZON,
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	MaxBatchingWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MaximumPollers: jsii.Number(123),
//   		MinimumPollers: jsii.Number(123),
//
//   		// the properties below are optional
//   		PollerGroupName: jsii.String("pollerGroupName"),
//   	},
//   }
//
type BaseStreamEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//   * 1000 for `DynamoEventSource`
	// * 10000 for `KinesisEventSource`, `ManagedKafkaEventSource` and `SelfManagedKafkaEventSource`.
	// Default: 100.
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5).
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-batching
	//
	// Default: - Duration.seconds(0) for Kinesis, DynamoDB, and SQS event sources, Duration.millis(500) for MSK, self-managed Kafka, and Amazon MQ.
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// Configuration for provisioned pollers that read from the event source.
	//
	// When specified, allows control over the minimum and maximum number of pollers
	// that can be provisioned to process events from the source.
	// Default: - no provisioned pollers.
	//
	ProvisionedPollerConfig *ProvisionedPollerConfig `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
}

