package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventSourceMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSourceMappingProps := &CfnEventSourceMappingProps{
//   	FunctionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	AmazonManagedKafkaEventSourceConfig: &AmazonManagedKafkaEventSourceConfigProperty{
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnFunctionError: jsii.Boolean(false),
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &OnFailureProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   	},
//   	DocumentDbEventSourceConfig: &DocumentDBEventSourceConfigProperty{
//   		CollectionName: jsii.String("collectionName"),
//   		DatabaseName: jsii.String("databaseName"),
//   		FullDocument: jsii.String("fullDocument"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	EventSourceArn: jsii.String("eventSourceArn"),
//   	FilterCriteria: &FilterCriteriaProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	FunctionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	ParallelizationFactor: jsii.Number(123),
//   	Queues: []*string{
//   		jsii.String("queues"),
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		MaximumConcurrency: jsii.Number(123),
//   	},
//   	SelfManagedEventSource: &SelfManagedEventSourceProperty{
//   		Endpoints: &EndpointsProperty{
//   			KafkaBootstrapServers: []*string{
//   				jsii.String("kafkaBootstrapServers"),
//   			},
//   		},
//   	},
//   	SelfManagedKafkaEventSourceConfig: &SelfManagedKafkaEventSourceConfigProperty{
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	SourceAccessConfigurations: []interface{}{
//   		&SourceAccessConfigurationProperty{
//   			Type: jsii.String("type"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	StartingPositionTimestamp: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	TumblingWindowInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
//
type CfnEventSourceMappingProps struct {
	// The name of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Specific configuration settings for an MSK event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig
	//
	AmazonManagedKafkaEventSourceConfig interface{} `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// The maximum number of items to retrieve in a single batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// (Streams) If the function returns an error, split the batch in two and retry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-bisectbatchonfunctionerror
	//
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// A configuration object that specifies the destination of an event after Lambda processes it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// Document db event source config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-documentdbeventsourceconfig
	//
	DocumentDbEventSourceConfig interface{} `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// Disables the event source mapping to pause polling and invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
	//
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// The filter criteria to control event filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-filtercriteria
	//
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// (Streams) A list of response types supported by the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionresponsetypes
	//
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// The Amazon Resource Name (ARN) of the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Streams) The maximum age of a record that Lambda sends to a function for processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumrecordageinseconds
	//
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// (Streams) The maximum number of times to retry when the function returns an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// (Streams) The number of batches to process from each shard concurrently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-parallelizationfactor
	//
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// (ActiveMQ) A list of ActiveMQ queues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-queues
	//
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// The scaling configuration for the event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-scalingconfig
	//
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The configuration used by AWS Lambda to access a self-managed event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-selfmanagedeventsource
	//
	SelfManagedEventSource interface{} `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// Specific configuration settings for a Self-Managed Apache Kafka event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig
	//
	SelfManagedKafkaEventSourceConfig interface{} `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// A list of SourceAccessConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-sourceaccessconfigurations
	//
	SourceAccessConfigurations interface{} `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in a stream from which to start reading.
	//
	// Required for Amazon Kinesis and Amazon DynamoDB Streams sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingpositiontimestamp
	//
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// (Kafka) A list of Kafka topics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-topics
	//
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-tumblingwindowinseconds
	//
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}

