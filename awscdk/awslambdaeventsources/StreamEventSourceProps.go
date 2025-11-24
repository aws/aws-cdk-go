package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// The set of properties for streaming event sources shared by Dynamo and Kinesis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventSourceDlq IEventSourceDlq
//   var filters interface{}
//   var key Key
//
//   streamEventSourceProps := &StreamEventSourceProps{
//   	StartingPosition: awscdk.Aws_lambda.StartingPosition_TRIM_HORIZON,
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnError: jsii.Boolean(false),
//   	Enabled: jsii.Boolean(false),
//   	FilterEncryption: key,
//   	Filters: []map[string]interface{}{
//   		map[string]interface{}{
//   			"filtersKey": filters,
//   		},
//   	},
//   	MaxBatchingWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	MaxRecordAge: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MetricsConfig: &MetricsConfig{
//   		Metrics: []eVENT_COUNT{
//   			awscdk.*Aws_lambda.MetricType_*eVENT_COUNT,
//   		},
//   	},
//   	OnFailure: eventSourceDlq,
//   	ParallelizationFactor: jsii.Number(123),
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MaximumPollers: jsii.Number(123),
//   		MinimumPollers: jsii.Number(123),
//
//   		// the properties below are optional
//   		PollerGroupName: jsii.String("pollerGroupName"),
//   	},
//   	ReportBatchItemFailures: jsii.Boolean(false),
//   	RetryAttempts: jsii.Number(123),
//   	TumblingWindow: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type StreamEventSourceProps struct {
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
	// If the function returns an error, split the batch in two and retry.
	// Default: false.
	//
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// Add Customer managed KMS key to encrypt Filter Criteria.
	// See: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
	//
	// Default: - none.
	//
	FilterEncryption awskms.IKey `field:"optional" json:"filterEncryption" yaml:"filterEncryption"`
	// Add filter criteria option.
	// Default: - None.
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	//
	// The default value is -1, which sets the maximum age to infinite.
	// When the value is set to infinite, Lambda never discards old records.
	// Record are valid until it expires in the event source.
	// Default: -1.
	//
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// Configuration for enhanced monitoring metrics collection When specified, enables collection of additional metrics for the stream event source.
	// Default: - Enhanced monitoring is disabled.
	//
	MetricsConfig *awslambda.MetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// An Amazon S3, Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Default: - discarded records are ignored.
	//
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	// Default: 1.
	//
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Default: false.
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// Maximum number of retry attempts.
	//
	// Set to -1 for infinite retries (until the record expires in the event source).
	//
	// Valid Range: -1 (infinite) or 0 to 10000.
	// Default: -1 (infinite retries).
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Default: - None.
	//
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
}

