package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for declaring a new event source mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventSourceDlq iEventSourceDlq
//   var filters interface{}
//   var function_ function
//   var sourceAccessConfigurationType sourceAccessConfigurationType
//
//   eventSourceMappingProps := &EventSourceMappingProps{
//   	Target: function_,
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnError: jsii.Boolean(false),
//   	Enabled: jsii.Boolean(false),
//   	EventSourceArn: jsii.String("eventSourceArn"),
//   	Filters: []map[string]interface{}{
//   		map[string]interface{}{
//   			"filtersKey": filters,
//   		},
//   	},
//   	KafkaBootstrapServers: []*string{
//   		jsii.String("kafkaBootstrapServers"),
//   	},
//   	KafkaConsumerGroupId: jsii.String("kafkaConsumerGroupId"),
//   	KafkaTopic: jsii.String("kafkaTopic"),
//   	MaxBatchingWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	MaxConcurrency: jsii.Number(123),
//   	MaxRecordAge: cdk.Duration_*Minutes(jsii.Number(30)),
//   	OnFailure: eventSourceDlq,
//   	ParallelizationFactor: jsii.Number(123),
//   	ReportBatchItemFailures: jsii.Boolean(false),
//   	RetryAttempts: jsii.Number(123),
//   	SourceAccessConfigurations: []sourceAccessConfiguration{
//   		&sourceAccessConfiguration{
//   			Type: sourceAccessConfigurationType,
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	StartingPosition: awscdk.Aws_lambda.StartingPosition_TRIM_HORIZON,
//   	StartingPositionTimestamp: jsii.Number(123),
//   	SupportS3OnFailureDestination: jsii.Boolean(false),
//   	TumblingWindow: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type EventSourceMappingProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10000.
	// Default: - Amazon Kinesis, Amazon DynamoDB, and Amazon MSK is 100 records.
	// The default for Amazon SQS is 10 messages. For standard SQS queues, the maximum is 10,000. For FIFO SQS queues, the maximum is 10.
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Default: false.
	//
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// Set to false to disable the event source upon creation.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// Any record added to
	// this stream can invoke the Lambda function.
	// Default: - not set if using a self managed Kafka cluster, throws an error otherwise.
	//
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// Add filter criteria to Event Source.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	//
	// Default: - none.
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// A list of host and port pairs that are the addresses of the Kafka brokers in a self managed "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.example.com:9096`.
	// Default: - none.
	//
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. The value must have a lenght between 1 and 200 and full the pattern '[a-zA-Z0-9-\/*:_+=.@-]*'. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html
	//
	// Default: - none.
	//
	KafkaConsumerGroupId *string `field:"optional" json:"kafkaConsumerGroupId" yaml:"kafkaConsumerGroupId"`
	// The name of the Kafka topic.
	// Default: - no topic.
	//
	KafkaTopic *string `field:"optional" json:"kafkaTopic" yaml:"kafkaTopic"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Default: Duration.seconds(0)
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The maximum concurrency setting limits the number of concurrent instances of the function that an Amazon SQS event source can invoke.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency
	//
	// Valid Range: Minimum value of 2. Maximum value of 1000.
	//
	// Default: - No specific limit.
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	// Default: - infinite or until the record expires.
	//
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Default: discarded records are ignored.
	//
	OnFailure IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
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
	// The maximum number of times to retry when the function returns an error.
	//
	// Set to `undefined` if you want lambda to keep retrying infinitely or until
	// the record expires.
	//
	// Valid Range:
	// * Minimum value of 0
	// * Maximum value of 10000.
	// Default: - infinite or until the record expires.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
	//
	// Default: - none.
	//
	SourceAccessConfigurations *[]*SourceAccessConfiguration `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType
	//
	// Default: - no starting position.
	//
	StartingPosition StartingPosition `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// The time from which to start reading, in Unix time seconds.
	// Default: - no timestamp.
	//
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// Check if support S3 onfailure destination(ODF).
	//
	// Currently only MSK and self managed kafka event support S3 ODF.
	// Default: false.
	//
	SupportS3OnFailureDestination *bool `field:"optional" json:"supportS3OnFailureDestination" yaml:"supportS3OnFailureDestination"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-windows
	//
	// Valid Range: 0 - 15 minutes.
	//
	// Default: - None.
	//
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
	// The target AWS Lambda function.
	Target IFunction `field:"required" json:"target" yaml:"target"`
}

