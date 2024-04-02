package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

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
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
//   		}),
//   	},
//   }))
//
type DynamoEventSourceProps struct {
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
	// If the function returns an error, split the batch in two and retry.
	// Default: false.
	//
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
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
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
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
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	//
	// The default value is -1, which sets the maximum number of retries to infinite.
	// When MaximumRetryAttempts is infinite, Lambda retries failed records until
	// the record expires in the event source.
	// Default: -1.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Default: - None.
	//
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
}

