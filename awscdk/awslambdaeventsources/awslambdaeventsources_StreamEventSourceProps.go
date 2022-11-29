package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// The set of properties for streaming event sources shared by Dynamo and Kinesis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var eventSourceDlq iEventSourceDlq
//
//   streamEventSourceProps := &streamEventSourceProps{
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnError: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	maxBatchingWindow: duration,
//   	maxRecordAge: duration,
//   	onFailure: eventSourceDlq,
//   	parallelizationFactor: jsii.Number(123),
//   	reportBatchItemFailures: jsii.Boolean(false),
//   	retryAttempts: jsii.Number(123),
//   	tumblingWindow: duration,
//   }
//
// Experimental.
type StreamEventSourceProps struct {
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	// Experimental.
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	// Experimental.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
}

