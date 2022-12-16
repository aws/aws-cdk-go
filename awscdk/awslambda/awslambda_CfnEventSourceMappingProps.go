package awslambda


// Properties for defining a `CfnEventSourceMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSourceMappingProps := &cfnEventSourceMappingProps{
//   	functionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	amazonManagedKafkaEventSourceConfig: &amazonManagedKafkaEventSourceConfigProperty{
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnFunctionError: jsii.Boolean(false),
//   	destinationConfig: &destinationConfigProperty{
//   		onFailure: &onFailureProperty{
//   			destination: jsii.String("destination"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	eventSourceArn: jsii.String("eventSourceArn"),
//   	filterCriteria: &filterCriteriaProperty{
//   		filters: []interface{}{
//   			&filterProperty{
//   				pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	functionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	parallelizationFactor: jsii.Number(123),
//   	queues: []*string{
//   		jsii.String("queues"),
//   	},
//   	scalingConfig: &scalingConfigProperty{
//   		maximumConcurrency: jsii.Number(123),
//   	},
//   	selfManagedEventSource: &selfManagedEventSourceProperty{
//   		endpoints: &endpointsProperty{
//   			kafkaBootstrapServers: []*string{
//   				jsii.String("kafkaBootstrapServers"),
//   			},
//   		},
//   	},
//   	selfManagedKafkaEventSourceConfig: &selfManagedKafkaEventSourceConfigProperty{
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	sourceAccessConfigurations: []interface{}{
//   		&sourceAccessConfigurationProperty{
//   			type: jsii.String("type"),
//   			uri: jsii.String("uri"),
//   		},
//   	},
//   	startingPosition: jsii.String("startingPosition"),
//   	startingPositionTimestamp: jsii.Number(123),
//   	topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	tumblingWindowInSeconds: jsii.Number(123),
//   }
//
type CfnEventSourceMappingProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Version or Alias ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// `AWS::Lambda::EventSourceMapping.AmazonManagedKafkaEventSourceConfig`.
	AmazonManagedKafkaEventSourceConfig interface{} `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//
	// - *Amazon Kinesis* - Default 100. Max 10,000.
	// - *Amazon DynamoDB Streams* - Default 100. Max 10,000.
	// - *Amazon Simple Queue Service* - Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	// - *Amazon Managed Streaming for Apache Kafka* - Default 100. Max 10,000.
	// - *Self-Managed Apache Kafka* - Default 100. Max 10,000.
	// - *Amazon MQ (ActiveMQ and RabbitMQ)* - Default 100. Max 10,000.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// (Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// (Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//
	// Default: True.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// - *Amazon Kinesis* - The ARN of the data stream or a stream consumer.
	// - *Amazon DynamoDB Streams* - The ARN of the stream.
	// - *Amazon Simple Queue Service* - The ARN of the queue.
	// - *Amazon Managed Streaming for Apache Kafka* - The ARN of the cluster.
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// (Streams and Amazon SQS) An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) .
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// (Streams and SQS) A list of current response type enums applied to the event source mapping.
	//
	// Valid Values: `ReportBatchItemFailures`.
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// *Default ( Kinesis , DynamoDB , Amazon SQS event sources)* : 0
	//
	// *Default ( Amazon MSK , Kafka, Amazon MQ event sources)* : 500 ms.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Streams only) Discard records older than the specified age.
	//
	// The default value is -1,
	// which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// (Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1,
	// which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// (Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// `AWS::Lambda::EventSourceMapping.ScalingConfig`.
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource interface{} `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// `AWS::Lambda::EventSourceMapping.SelfManagedKafkaEventSourceConfig`.
	SelfManagedKafkaEventSourceConfig interface{} `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	SourceAccessConfigurations interface{} `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB.
	//
	// - *LATEST* - Read only new records.
	// - *TRIM_HORIZON* - Process all available records.
	// - *AT_TIMESTAMP* - Specify a time from which to start reading records.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// With `StartingPosition` set to `AT_TIMESTAMP` , the time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// The name of the Kafka topic.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// (Streams only) The duration in seconds of a processing window.
	//
	// The range is between 1 second up to 900 seconds.
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}

