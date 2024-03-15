package awslambda


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
//   	Topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	TumblingWindowInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
//
type CfnEventSourceMappingProps struct {
	// The name or ARN of the Lambda function.
	//
	// **Name formats** - *Function name* – `MyFunction` .
	// - *Function ARN* – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Version or Alias ARN* – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` .
	// - *Partial ARN* – `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig
	//
	AmazonManagedKafkaEventSourceConfig interface{} `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//
	// - *Amazon Kinesis* – Default 100. Max 10,000.
	// - *Amazon DynamoDB Streams* – Default 100. Max 10,000.
	// - *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	// - *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.
	// - *Self-managed Apache Kafka* – Default 100. Max 10,000.
	// - *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.
	// - *DocumentDB* – Default 100. Max 10,000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-bisectbatchonfunctionerror
	//
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event sources only) A configuration object that specifies the destination of an event after Lambda processes it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// Specific configuration settings for a DocumentDB event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-documentdbeventsourceconfig
	//
	DocumentDbEventSourceConfig interface{} `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//
	// Default: True.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// - *Amazon Kinesis* – The ARN of the data stream or a stream consumer.
	// - *Amazon DynamoDB Streams* – The ARN of the stream.
	// - *Amazon Simple Queue Service* – The ARN of the queue.
	// - *Amazon Managed Streaming for Apache Kafka* – The ARN of the cluster or the ARN of the VPC connection (for [cross-account event source mappings](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#msk-multi-vpc) ).
	// - *Amazon MQ* – The ARN of the broker.
	// - *Amazon DocumentDB* – The ARN of the DocumentDB change stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
	//
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-filtercriteria
	//
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// (Streams and SQS) A list of current response type enums applied to the event source mapping.
	//
	// Valid Values: `ReportBatchItemFailures`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionresponsetypes
	//
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// *Default ( Kinesis , DynamoDB , Amazon SQS event sources)* : 0
	//
	// *Default ( Amazon MSK , Kafka, Amazon MQ , Amazon DocumentDB event sources)* : 500 ms
	//
	// *Related setting:* For Amazon SQS event sources, when you set `BatchSize` to a value greater than 10, you must set `MaximumBatchingWindowInSeconds` to at least 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Kinesis and DynamoDB Streams only) Discard records older than the specified age.
	//
	// The default value is -1,
	// which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	//
	// > The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumrecordageinseconds
	//
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// (Kinesis and DynamoDB Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1,
	// which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-parallelizationfactor
	//
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-queues
	//
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// (Amazon SQS only) The scaling configuration for the event source.
	//
	// For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-scalingconfig
	//
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The self-managed Apache Kafka cluster for your event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-selfmanagedeventsource
	//
	SelfManagedEventSource interface{} `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// Specific configuration settings for a self-managed Apache Kafka event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig
	//
	SelfManagedKafkaEventSourceConfig interface{} `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-sourceaccessconfigurations
	//
	SourceAccessConfigurations interface{} `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB.
	//
	// - *LATEST* - Read only new records.
	// - *TRIM_HORIZON* - Process all available records.
	// - *AT_TIMESTAMP* - Specify a time from which to start reading records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// With `StartingPosition` set to `AT_TIMESTAMP` , the time from which to start reading, in Unix time seconds.
	//
	// `StartingPositionTimestamp` cannot be in the future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingpositiontimestamp
	//
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// The name of the Kafka topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-topics
	//
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.
	//
	// A value of 0 seconds indicates no tumbling window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-tumblingwindowinseconds
	//
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}

