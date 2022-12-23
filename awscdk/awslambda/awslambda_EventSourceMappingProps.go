package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for declaring a new event source mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventSourceDlq iEventSourceDlq
//   var filters interface{}
//   var function_ function
//   var sourceAccessConfigurationType sourceAccessConfigurationType
//
//   eventSourceMappingProps := &eventSourceMappingProps{
//   	target: function_,
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnError: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	eventSourceArn: jsii.String("eventSourceArn"),
//   	filters: []map[string]interface{}{
//   		map[string]interface{}{
//   			"filtersKey": filters,
//   		},
//   	},
//   	kafkaBootstrapServers: []*string{
//   		jsii.String("kafkaBootstrapServers"),
//   	},
//   	kafkaConsumerGroupId: jsii.String("kafkaConsumerGroupId"),
//   	kafkaTopic: jsii.String("kafkaTopic"),
//   	maxBatchingWindow: cdk.duration.minutes(jsii.Number(30)),
//   	maxRecordAge: cdk.*duration.minutes(jsii.Number(30)),
//   	onFailure: eventSourceDlq,
//   	parallelizationFactor: jsii.Number(123),
//   	reportBatchItemFailures: jsii.Boolean(false),
//   	retryAttempts: jsii.Number(123),
//   	sourceAccessConfigurations: []sourceAccessConfiguration{
//   		&sourceAccessConfiguration{
//   			type: sourceAccessConfigurationType,
//   			uri: jsii.String("uri"),
//   		},
//   	},
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//   	startingPositionTimestamp: jsii.Number(123),
//   	tumblingWindow: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type EventSourceMappingProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10000.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// Set to false to disable the event source upon creation.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// Any record added to
	// this stream can invoke the Lambda function.
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// Add filter criteria to Event Source.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// A list of host and port pairs that are the addresses of the Kafka brokers in a self managed "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.example.com:9096`.
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. The value must have a lenght between 1 and 200 and full the pattern '[a-zA-Z0-9-\/*:_+=.@-]*'. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html
	//
	KafkaConsumerGroupId *string `field:"optional" json:"kafkaConsumerGroupId" yaml:"kafkaConsumerGroupId"`
	// The name of the Kafka topic.
	KafkaTopic *string `field:"optional" json:"kafkaTopic" yaml:"kafkaTopic"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
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
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
	//
	SourceAccessConfigurations *[]*SourceAccessConfiguration `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType
	//
	StartingPosition StartingPosition `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// The time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-windows
	//
	// Valid Range: 0 - 15 minutes.
	//
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
	// The target AWS Lambda function.
	Target IFunction `field:"required" json:"target" yaml:"target"`
}

