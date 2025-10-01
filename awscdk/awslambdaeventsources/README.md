# AWS Lambda Event Sources

An event source mapping is an AWS Lambda resource that reads from an event source and invokes a Lambda function.
You can use event source mappings to process items from a stream or queue in services that don't invoke Lambda
functions directly. Lambda provides event source mappings for the following services. Read more about lambda
event sources [here](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html).

This module includes classes that allow using various AWS services as event
sources for AWS Lambda via the high-level `lambda.addEventSource(source)` API.

NOTE: In most cases, it is also possible to use the resource APIs to invoke an
AWS Lambda function. This library provides a uniform API for all Lambda event
sources regardless of the underlying mechanism they use.

The following code sets up a lambda function with an SQS queue event source -

```go
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

queue := sqs.NewQueue(this, jsii.String("MyQueue"))
eventSource := awscdk.NewSqsEventSource(queue)
fn.AddEventSource(eventSource)

eventSourceId := eventSource.eventSourceMappingId
eventSourceMappingArn := eventSource.eventSourceMappingArn
```

The `eventSourceId` property contains the event source id. This will be a
[token](https://docs.aws.amazon.com/cdk/latest/guide/tokens.html) that will resolve to the final value at the time of
deployment.

The `eventSourceMappingArn` property contains the event source mapping ARN. This will be a
[token](https://docs.aws.amazon.com/cdk/latest/guide/tokens.html) that will resolve to the final value at the time of
deployment.

## SQS

Amazon Simple Queue Service (Amazon SQS) allows you to build asynchronous
workflows. For more information about Amazon SQS, see Amazon Simple Queue
Service. You can configure AWS Lambda to poll for these messages as they arrive
and then pass the event to a Lambda function invocation. To view a sample event,
see [Amazon SQS Event](https://docs.aws.amazon.com/lambda/latest/dg/eventsources.html#eventsources-sqs).

To set up Amazon Simple Queue Service as an event source for AWS Lambda, you
first create or update an Amazon SQS queue and select custom values for the
queue parameters. The following parameters will impact Amazon SQS's polling
behavior:

* **visibilityTimeout**: May impact the period between retries.
* **batchSize**: Determines how many records are buffered before invoking your lambda function.
* **maxBatchingWindow**: The maximum amount of time to gather records before invoking the lambda. This increases the likelihood of a full batch at the cost of delayed processing.
* **maxConcurrency**: The maximum concurrency setting limits the number of concurrent instances of the function that an Amazon SQS event source can invoke.
* **enabled**: If the SQS event source mapping should be enabled. The default is true.

```go
import "github.com/aws/aws-cdk-go/awscdk"
var fn function


queue := sqs.NewQueue(this, jsii.String("MyQueue"), &QueueProps{
	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(30)),
})

fn.AddEventSource(awscdk.NewSqsEventSource(queue, &SqsEventSourceProps{
	BatchSize: jsii.Number(10),
	 // default
	MaxBatchingWindow: awscdk.Duration_Minutes(jsii.Number(5)),
	ReportBatchItemFailures: jsii.Boolean(true),
}))
```

## S3

You can write Lambda functions to process S3 bucket events, such as the
object-created or object-deleted events. For example, when a user uploads a
photo to a bucket, you might want Amazon S3 to invoke your Lambda function so
that it reads the image and creates a thumbnail for the photo.

You can use the bucket notification configuration feature in Amazon S3 to
configure the event source mapping, identifying the bucket events that you want
Amazon S3 to publish and which Lambda function to invoke.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var fn function


bucket := s3.NewBucket(this, jsii.String("mybucket"))

fn.AddEventSource(awscdk.NewS3EventSource(bucket, &S3EventSourceProps{
	Events: []eventType{
		s3.*eventType_OBJECT_CREATED,
		s3.*eventType_OBJECT_REMOVED,
	},
	Filters: []notificationKeyFilter{
		&notificationKeyFilter{
			Prefix: jsii.String("subdir/"),
		},
	},
}))
```

In the example above, `S3EventSource` is accepting `Bucket` type as parameter.
However, Functions like `from_bucket_name` and `from_bucket_arn` will return `IBucket`
and is not compliant with `S3EventSource`. If this is the case, please consider using
`S3EventSourceV2` instead, this class accepts `IBucket`.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var fn function


bucket := s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket"))

fn.AddEventSource(awscdk.NewS3EventSourceV2(bucket, &S3EventSourceProps{
	Events: []eventType{
		s3.*eventType_OBJECT_CREATED,
		s3.*eventType_OBJECT_REMOVED,
	},
	Filters: []notificationKeyFilter{
		&notificationKeyFilter{
			Prefix: jsii.String("subdir/"),
		},
	},
}))
```

## SNS

You can write Lambda functions to process Amazon Simple Notification Service
notifications. When a message is published to an Amazon SNS topic, the service
can invoke your Lambda function by passing the message payload as a parameter.
Your Lambda function code can then process the event, for example publish the
message to other Amazon SNS topics, or send the message to other AWS services.

This also enables you to trigger a Lambda function in response to Amazon
CloudWatch alarms and other AWS services that use Amazon SNS.

For an example event, see [Appendix: Message and JSON
Formats](https://docs.aws.amazon.com/sns/latest/dg/json-formats.html) and
[Amazon SNS Sample
Event](https://docs.aws.amazon.com/lambda/latest/dg/eventsources.html#eventsources-sns).
For an example use case, see [Using AWS Lambda with Amazon SNS from Different
Accounts](https://docs.aws.amazon.com/lambda/latest/dg/with-sns.html).

```go
import sns "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var topic topic

var fn function

deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
fn.AddEventSource(awscdk.NewSnsEventSource(topic, &SnsEventSourceProps{
	FilterPolicy: map[string]interface{}{
	},
	DeadLetterQueue: deadLetterQueue,
}))
```

When a user calls the SNS Publish API on a topic that your Lambda function is
subscribed to, Amazon SNS will call Lambda to invoke your function
asynchronously. Lambda will then return a delivery status. If there was an error
calling Lambda, Amazon SNS will retry invoking the Lambda function up to three
times. After three tries, if Amazon SNS still could not successfully invoke the
Lambda function, then Amazon SNS will send a delivery status failure message to
CloudWatch.

## DynamoDB Streams

You can write Lambda functions to process change events from a DynamoDB Table. An event is emitted to a DynamoDB stream (if configured) whenever a write (Put, Delete, Update)
operation is performed against the table. See [Using AWS Lambda with Amazon DynamoDB](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html) for more information about configuring Lambda function event sources with DynamoDB.

To process events with a Lambda function, first create or update a DynamoDB table and enable a `stream` specification. Then, create a `DynamoEventSource`
and add it to your Lambda function. The following parameters will impact Amazon DynamoDB's polling behavior:

* **batchSize**: Determines how many records are buffered before invoking your lambda function - could impact your function's memory usage (if too high) and ability to keep up with incoming data velocity (if too low).
* **bisectBatchOnError**: If a batch encounters an error, this will cause the batch to be split in two and have each new smaller batch retried, allowing the records in error to be isolated.
* **reportBatchItemFailures**: Allow functions to return partially successful responses for a batch of records.
* **maxBatchingWindow**: The maximum amount of time to gather records before invoking the lambda. This increases the likelihood of a full batch at the cost of delayed processing.
* **maxRecordAge**: The maximum age of a record that will be sent to the function for processing. Records that exceed the max age will be treated as failures.
* **onFailure**: In the event a record fails after all retries or if the record age has exceeded the configured value, the record will be sent to S3 bucket, SQS queue or SNS topic that is specified here
* **parallelizationFactor**: The number of batches to concurrently process on each shard.
* **retryAttempts**: The maximum number of times a record should be retried in the event of failure.
* **startingPosition**: Will determine where to being consumption, either at the most recent ('LATEST') record or the oldest record ('TRIM_HORIZON'). 'TRIM_HORIZON' will ensure you process all available data, while 'LATEST' will ignore all records that arrived prior to attaching the event source.
* **tumblingWindow**: The duration in seconds of a processing window when using streams.
* **enabled**: If the DynamoDB Streams event source mapping should be enabled. The default is true.
* **filters**: Filters to apply before sending a change event from a DynamoDB table to a Lambda function. Events that are filtered out are not sent to the Lambda function.

```go
import dynamodb "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var table table

var fn function


deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
fn.AddEventSource(awscdk.NewDynamoEventSource(table, &DynamoEventSourceProps{
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	BatchSize: jsii.Number(5),
	BisectBatchOnError: jsii.Boolean(true),
	OnFailure: awscdk.NewSqsDlq(deadLetterQueue),
	RetryAttempts: jsii.Number(10),
}))
```

The following code sets up a Lambda function with a DynamoDB event source. A filter is applied to only send DynamoDB events to
the Lambda function when the `id` column is a boolean that equals `true`.

```go
import dynamodb "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var table table

var fn function

fn.AddEventSource(awscdk.NewDynamoEventSource(table, &DynamoEventSourceProps{
	StartingPosition: lambda.StartingPosition_LATEST,
	Filters: []map[string]interface{}{
		lambda.FilterCriteria_Filter(map[string]interface{}{
			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
			"dynamodb": map[string]map[string]map[string]interface{}{
				"NewImage": map[string]map[string]interface{}{
					"id": map[string]interface{}{
						"BOOL": lambda.FilterRule_isEqual(jsii.Boolean(true)),
					},
				},
			},
		}),
	},
}))
```

## Kinesis

You can write Lambda functions to process streaming data in Amazon Kinesis Streams. For more information about Amazon Kinesis, see [Amazon Kinesis
Service](https://aws.amazon.com/kinesis/data-streams/). To learn more about configuring Lambda function event sources with kinesis and view a sample event,
see [Amazon Kinesis Event](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html).

To set up Amazon Kinesis as an event source for AWS Lambda, you
first create or update an Amazon Kinesis stream and select custom values for the
event source parameters. The following parameters will impact Amazon Kinesis's polling
behavior:

* **batchSize**: Determines how many records are buffered before invoking your lambda function - could impact your function's memory usage (if too high) and ability to keep up with incoming data velocity (if too low).
* **bisectBatchOnError**: If a batch encounters an error, this will cause the batch to be split in two and have each new smaller batch retried, allowing the records in error to be isolated.
* **reportBatchItemFailures**: Allow functions to return partially successful responses for a batch of records.
* **maxBatchingWindow**: The maximum amount of time to gather records before invoking the lambda. This increases the likelihood of a full batch at the cost of possibly delaying processing.
* **maxRecordAge**: The maximum age of a record that will be sent to the function for processing. Records that exceed the max age will be treated as failures.
* **onFailure**: In the event a record fails and consumes all retries, the record will be sent to S3 bucket, SQS queue or SNS topic that is specified here
* **parallelizationFactor**: The number of batches to concurrently process on each shard.
* **retryAttempts**: The maximum number of times a record should be retried in the event of failure.
* **startingPosition**: Will determine where to begin consumption. 'LATEST' will start at the most recent record and ignore all records that arrived prior to attaching the event source, 'TRIM_HORIZON' will start at the oldest record and ensure you process all available data, while 'AT_TIMESTAMP' will start reading records from a specified time stamp.
* **startingPositionTimestamp**: The time stamp from which to start reading. Used in conjunction with **startingPosition** when set to 'AT_TIMESTAMP'.
* **tumblingWindow**: The duration in seconds of a processing window when using streams.
* **enabled**: If the event source mapping should be enabled. The default is true.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var myFunction function


stream := kinesis.NewStream(this, jsii.String("MyStream"))
myFunction.AddEventSource(awscdk.NewKinesisEventSource(stream, &KinesisEventSourceProps{
	BatchSize: jsii.Number(100),
	 // default
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
}))
```

To use a dedicated-throughput consumer with enhanced fan-out

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var myFunction function


stream := kinesis.NewStream(this, jsii.String("MyStream"))
streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
	Stream: Stream,
	StreamConsumerName: jsii.String("MyStreamConsumer"),
})
myFunction.AddEventSource(awscdk.NewKinesisConsumerEventSource(streamConsumer, &KinesisEventSourceProps{
	BatchSize: jsii.Number(100),
	 // default
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
}))
```

## Kafka

You can write Lambda functions to process data either from [Amazon MSK](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html) or a [self-managed Kafka](https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html) cluster. The following parameters will impact to the polling behavior:

* **startingPosition**: Will determine where to begin consumption. 'LATEST' will start at the most recent record and ignore all records that arrived prior to attaching the event source, 'TRIM_HORIZON' will start at the oldest record and ensure you process all available data, while 'AT_TIMESTAMP' will start reading records from a specified time stamp.
* **startingPositionTimestamp**: The time stamp from which to start reading. Used in conjunction with **startingPosition** when set to 'AT_TIMESTAMP'.
* **batchSize**: Determines how many records are buffered before invoking your lambda function - could impact your function's memory usage (if too high) and ability to keep up with incoming data velocity (if too low).
* **maxBatchingWindow**: The maximum amount of time to gather records before invoking the lambda. This increases the likelihood of a full batch at the cost of possibly delaying processing.
* **onFailure**: In the event a record fails and consumes all retries, the record will be sent to SQS queue or SNS topic that is specified here
* **enabled**: If the Kafka event source mapping should be enabled. The default is true.

The following code sets up Amazon MSK as an event source for a lambda function. Credentials will need to be configured to access the
MSK cluster, as described in [Username/Password authentication](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html).

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var myFunction function


// Your MSK cluster arn
clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"

// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

// The secret that allows access to your MSK cluster
// You still have to make sure that it is associated with your cluster as described in the documentation
secret := awscdk.NewSecret(this, jsii.String("Secret"), &SecretProps{
	SecretName: jsii.String("AmazonMSK_KafkaSecret"),
})
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: topic,
	Secret: secret,
	BatchSize: jsii.Number(100),
	 // default
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
}))
```

The following code sets up a self managed Kafka cluster as an event source. Username and password based authentication
will need to be set up as described in [Managing access and permissions](https://docs.aws.amazon.com/lambda/latest/dg/smaa-permissions.html#smaa-permissions-add-secret).

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// The secret that allows access to your self hosted Kafka cluster
var secret secret

var myFunction function


// The list of Kafka brokers
bootstrapServers := []*string{
	"kafka-broker:9092",
}

// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

// (Optional) The consumer group id to use when connecting to the Kafka broker. If omitted the UUID of the event source mapping will be used.
consumerGroupId := "my-consumer-group-id"
myFunction.AddEventSource(awscdk.NewSelfManagedKafkaEventSource(&SelfManagedKafkaEventSourceProps{
	BootstrapServers: bootstrapServers,
	Topic: topic,
	ConsumerGroupId: consumerGroupId,
	Secret: secret,
	BatchSize: jsii.Number(100),
	 // default
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
}))
```

If your self managed Kafka cluster is only reachable via VPC also configure `vpc` `vpcSubnets` and `securityGroup`.

You can specify [event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-msk-smak)
for managed and self managed Kafka clusters using the `filters` property:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var myFunction function


// Your MSK cluster arn
clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"

// The Kafka topic you want to subscribe to
topic := "some-cool-topic"
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	Filters: []map[string]interface{}{
		lambda.FilterCriteria_Filter(map[string]interface{}{
			"stringEquals": lambda.FilterRule_isEqual(jsii.String("test")),
		}),
	},
}))
```

By default, Lambda will encrypt Filter Criteria using AWS managed keys. But if you want to use a self managed KMS key to encrypt the filters, You can specify the self managed key using the `filterEncryption` property.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var myFunction function


// Your MSK cluster arn
clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"

// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

// Your self managed KMS key
myKey := awscdk.Key_FromKeyArn(this, jsii.String("SourceBucketEncryptionKey"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/<key-id>"))
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	Filters: []map[string]interface{}{
		lambda.FilterCriteria_Filter(map[string]interface{}{
			"stringEquals": lambda.FilterRule_isEqual(jsii.String("test")),
		}),
	},
	FilterEncryption: myKey,
}))
```

You can also specify an S3 bucket as an "on failure" destination:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var bucket iBucket
var myFunction function


// Your MSK cluster arn
clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"

// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

s3OnFailureDestination := awscdk.NewS3OnFailureDestination(bucket)

myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	OnFailure: s3OnFailureDestination,
}))
```

Set configuration for provisioned pollers that read from the event source.

```go
import "github.com/aws/aws-cdk-go/awscdk"

// Your MSK cluster arn
var clusterArn string

var myFunction function


// The Kafka topic you want to subscribe to
topic := "some-cool-topic"
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	ProvisionedPollerConfig: &ProvisionedPollerConfig{
		MinimumPollers: jsii.Number(1),
		MaximumPollers: jsii.Number(3),
	},
}))
```

Set a confluent or self-managed schema registry to de-serialize events from the event source. Note, this will similarly work for `SelfManagedKafkaEventSource` but the example only shows setup for `ManagedKafkaEventSource`.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// Your MSK cluster arn
var clusterArn string

var myFunction function


// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

secret := awscdk.NewSecret(this, jsii.String("Secret"), &SecretProps{
	SecretName: jsii.String("AmazonMSK_KafkaSecret"),
})
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	ProvisionedPollerConfig: &ProvisionedPollerConfig{
		MinimumPollers: jsii.Number(1),
		MaximumPollers: jsii.Number(3),
	},
	SchemaRegistryConfig: awscdk.NewConfluentSchemaRegistry(&ConfluentSchemaRegistryProps{
		SchemaRegistryUri: jsii.String("https://example.com"),
		EventRecordFormat: lambda.EventRecordFormat_JSON(),
		AuthenticationType: lambda.KafkaSchemaRegistryAccessConfigType_BASIC_AUTH(),
		Secret: secret,
		SchemaValidationConfigs: []kafkaSchemaValidationConfig{
			&kafkaSchemaValidationConfig{
				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
			},
		},
	}),
}))
```

Set Glue schema registry to de-serialize events from the event source. Note, this will similarly work for `SelfManagedKafkaEventSource` but the example only shows setup for `ManagedKafkaEventSource`.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// Your MSK cluster arn
var clusterArn string

var myFunction function


// The Kafka topic you want to subscribe to
topic := "some-cool-topic"

// Your Glue Schema Registry
glueRegistry := awscdk.NewCfnRegistry(this, jsii.String("Registry"), &CfnRegistryProps{
	Name: jsii.String("schema-registry"),
	Description: jsii.String("Schema registry for event source"),
})
myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
	ClusterArn: jsii.String(ClusterArn),
	Topic: jsii.String(Topic),
	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
	ProvisionedPollerConfig: &ProvisionedPollerConfig{
		MinimumPollers: jsii.Number(1),
		MaximumPollers: jsii.Number(3),
	},
	SchemaRegistryConfig: awscdk.NewGlueSchemaRegistry(&GlueSchemaRegistryProps{
		SchemaRegistry: glueRegistry,
		EventRecordFormat: lambda.EventRecordFormat_JSON(),
		SchemaValidationConfigs: []kafkaSchemaValidationConfig{
			&kafkaSchemaValidationConfig{
				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
			},
		},
	}),
}))
```

## Roadmap

Eventually, this module will support all the event sources described under
[Supported Event
Sources](https://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-function.html)
in the AWS Lambda Developer Guide.
