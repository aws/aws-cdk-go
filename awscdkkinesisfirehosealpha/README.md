# Amazon Data Firehose Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Developer Preview](https://img.shields.io/badge/cdk--constructs-developer--preview-informational.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are in **developer preview** before they
> become stable. We will only make breaking changes to address unforeseen API issues. Therefore,
> these APIs are not subject to [Semantic Versioning](https://semver.org/), and breaking changes
> will be announced in release notes. This means that while you may use them, you may need to
> update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[Amazon Data Firehose](https://docs.aws.amazon.com/firehose/latest/dev/what-is-this-service.html), [formerly known as Amazon Kinesis Data Firehose](https://aws.amazon.com/about-aws/whats-new/2024/02/amazon-data-firehose-formerly-kinesis-data-firehose/),
is a service for fully-managed delivery of real-time streaming data to storage services
such as Amazon S3, Amazon Redshift, Amazon Elasticsearch, Splunk, or any custom HTTP
endpoint or third-party services such as Datadog, Dynatrace, LogicMonitor, MongoDB, New
Relic, and Sumo Logic.

Amazon Data Firehose delivery streams are distinguished from Kinesis data streams in
their models of consumption. Whereas consumers read from a data stream by actively pulling
data from the stream, a delivery stream pushes data to its destination on a regular
cadence. This means that data streams are intended to have consumers that do on-demand
processing, like AWS Lambda or Amazon EC2. On the other hand, delivery streams are
intended to have destinations that are sources for offline processing and analytics, such
as Amazon S3 and Amazon Redshift.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk)
project. It allows you to define Amazon Data Firehose delivery streams.

## Defining a Delivery Stream

In order to define a Delivery Stream, you must specify a destination. An S3 bucket can be
used as a destination. Currently the CDK supports only S3 as a destination which is covered [below](#destinations).

```go
bucket := s3.NewBucket(this, jsii.String("Bucket"))
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destinations.NewS3Bucket(bucket),
})
```

The above example defines the following resources:

* An S3 bucket
* An Amazon Data Firehose delivery stream with Direct PUT as the source and CloudWatch
  error logging turned on.
* An IAM role which gives the delivery stream permission to write to the S3 bucket.

## Sources

An Amazon Data Firehose delivery stream can accept data from three main sources: Kinesis Data Streams, Managed Streaming for Apache Kafka (MSK), or via a "direct put" (API calls). Currently only Kinesis Data Streams and direct put are supported in the CDK.

See: [Sending Data to a Delivery Stream](https://docs.aws.amazon.com/firehose/latest/dev/basic-write.html)
in the *Amazon Data Firehose Developer Guide*.

### Kinesis Data Stream

A delivery stream can read directly from a Kinesis data stream as a consumer of the data
stream. Configure this behaviour by passing in a data stream in the `source`
property via the `KinesisStreamSource` class when constructing a delivery stream:

```go
var destination iDestination

sourceStream := kinesis.NewStream(this, jsii.String("Source Stream"))

firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Source: firehose.NewKinesisStreamSource(sourceStream),
	Destination: destination,
})
```

### Direct Put

Data must be provided via "direct put", ie., by using a `PutRecord` or
`PutRecordBatch` API call. There are a number of ways of doing so, such as:

* Kinesis Agent: a standalone Java application that monitors and delivers files while
  handling file rotation, checkpointing, and retries. See: [Writing to Amazon Data Firehose Using Kinesis Agent](https://docs.aws.amazon.com/firehose/latest/dev/writing-with-agents.html)
  in the *Amazon Data Firehose Developer Guide*.
* AWS SDK: a general purpose solution that allows you to deliver data to a delivery stream
  from anywhere using Java, .NET, Node.js, Python, or Ruby. See: [Writing to Amazon Data Firehose Using the AWS SDK](https://docs.aws.amazon.com/firehose/latest/dev/writing-with-sdk.html)
  in the *Amazon Data Firehose Developer Guide*.
* CloudWatch Logs: subscribe to a log group and receive filtered log events directly into
  a delivery stream. See: [logs-destinations](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-logs-destinations-readme.html).
* Eventbridge: add an event rule target to send events to a delivery stream based on the
  rule filtering. See: [events-targets](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-events-targets-readme.html).
* SNS: add a subscription to send all notifications from the topic to a delivery
  stream. See: [sns-subscriptions](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-sns-subscriptions-readme.html).
* IoT: add an action to an IoT rule to send various IoT information to a delivery stream

## Destinations

Amazon Data Firehose supports multiple AWS and third-party services as destinations, including Amazon S3, Amazon Redshift, and more. You can find the full list of supported destination [here](https://docs.aws.amazon.com/firehose/latest/dev/create-destination.html).

Currently in the AWS CDK, only S3 is implemented as an L2 construct destination. Other destinations can still be configured using L1 constructs. See [kinesisfirehose-destinations](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-kinesisfirehose-destinations-readme.html)
for the implementations of these destinations.

### S3

Defining a delivery stream with an S3 bucket destination:

```go
var bucket bucket

s3Destination := destinations.NewS3Bucket(bucket)

firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: s3Destination,
})
```

The S3 destination also supports custom dynamic prefixes. `dataOutputPrefix`
will be used for files successfully delivered to S3. `errorOutputPrefix` will be added to
failed records before writing them to S3.

```go
var bucket bucket

s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	DataOutputPrefix: jsii.String("myFirehose/DeliveredYear=!{timestamp:yyyy}/anyMonth/rand=!{firehose:random-string}"),
	ErrorOutputPrefix: jsii.String("myFirehoseFailures/!{firehose:error-output-type}/!{timestamp:yyyy}/anyMonth/!{timestamp:dd}"),
})
```

See: [Custom S3 Prefixes](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html)
in the *Amazon Data Firehose Developer Guide*.

## Server-side Encryption

Enabling server-side encryption (SSE) requires Amazon Data Firehose to encrypt all data
sent to delivery stream when it is stored at rest. This means that data is encrypted
before being written to the service's internal storage layer and decrypted after it is
received from the internal storage layer. The service manages keys and cryptographic
operations so that sources and destinations do not need to, as the data is encrypted and
decrypted at the boundaries of the service (i.e., before the data is delivered to a
destination). By default, delivery streams do not have SSE enabled.

The Key Management Service keys (KMS keys) used for SSE can either be AWS-owned or
customer-managed. AWS-owned KMS keys are created, owned and managed by AWS for use in
multiple AWS accounts. As a customer, you cannot view, use, track, or manage these keys,
and you are not charged for their use. On the other hand, customer-managed KMS keys are
created and owned within your account and managed entirely by you. As a customer, you are
responsible for managing access, rotation, aliases, and deletion for these keys, and you
are changed for their use.

See: [AWS KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys)
in the *KMS Developer Guide*.

```go
var destination iDestination
// SSE with an customer-managed key that is explicitly specified
var key key


// SSE with an AWS-owned key
// SSE with an AWS-owned key
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with AWS Owned Key"), &DeliveryStreamProps{
	Encryption: firehose.StreamEncryption_AwsOwnedKey(),
	Destination: destination,
})
// SSE with an customer-managed key that is created automatically by the CDK
// SSE with an customer-managed key that is created automatically by the CDK
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with Customer Managed Key"), &DeliveryStreamProps{
	Encryption: firehose.StreamEncryption_CustomerManagedKey(),
	Destination: destination,
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with Customer Managed and Provided Key"), &DeliveryStreamProps{
	Encryption: firehose.StreamEncryption_*CustomerManagedKey(key),
	Destination: destination,
})
```

See: [Data Protection](https://docs.aws.amazon.com/firehose/latest/dev/encryption.html)
in the *Amazon Data Firehose Developer Guide*.

## Monitoring

Amazon Data Firehose is integrated with CloudWatch, so you can monitor the performance of
your delivery streams via logs and metrics.

### Logs

Amazon Data Firehose will send logs to CloudWatch when data transformation or data
delivery fails. The CDK will enable logging by default and create a CloudWatch LogGroup
and LogStream with default settings for your Delivery Stream.

When creating a destination, you can provide an `ILoggingConfig`, which can either be an `EnableLogging` or `DisableLogging` instance.
If you use `EnableLogging`, the CDK will create a CloudWatch LogGroup and LogStream with all CloudFormation default settings for you, or you can optionally
specify your own log group to be used for capturing and storing log events. For example:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"
var bucket bucket


logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	LoggingConfig: destinations.NewEnableLogging(logGroup),
})

firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destination,
})
```

Logging can also be disabled:

```go
var bucket bucket

destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	LoggingConfig: destinations.NewDisableLogging(),
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destination,
})
```

See: [Monitoring using CloudWatch Logs](https://docs.aws.amazon.com/firehose/latest/dev/monitoring-with-cloudwatch-logs.html)
in the *Amazon Data Firehose Developer Guide*.

### Metrics

Amazon Data Firehose sends metrics to CloudWatch so that you can collect and analyze the
performance of the delivery stream, including data delivery, data ingestion, data
transformation, format conversion, API usage, encryption, and resource usage. You can then
use CloudWatch alarms to alert you, for example, when data freshness (the age of the
oldest record in the delivery stream) exceeds the buffering limit (indicating that data is
not being delivered to your destination), or when the rate of incoming records exceeds the
limit of records per second (indicating data is flowing into your delivery stream faster
than it is configured to process).

CDK provides methods for accessing delivery stream metrics with default configuration,
such as `metricIncomingBytes`, and `metricIncomingRecords` (see [`IDeliveryStream`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-kinesisfirehose.IDeliveryStream.html)
for a full list). CDK also provides a generic `metric` method that can be used to produce
metric configurations for any metric provided by Amazon Data Firehose; the configurations
are pre-populated with the correct dimensions for the delivery stream.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var deliveryStream deliveryStream


// Alarm that triggers when the per-second average of incoming bytes exceeds 90% of the current service limit
incomingBytesPercentOfLimit := cloudwatch.NewMathExpression(&MathExpressionProps{
	Expression: jsii.String("incomingBytes / 300 / bytePerSecLimit"),
	UsingMetrics: map[string]iMetric{
		"incomingBytes": deliveryStream.metricIncomingBytes(&MetricOptions{
			"statistic": cloudwatch.Statistic_SUM,
		}),
		"bytePerSecLimit": deliveryStream.metric(jsii.String("BytesPerSecondLimit")),
	},
})

cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
	Metric: incomingBytesPercentOfLimit,
	Threshold: jsii.Number(0.9),
	EvaluationPeriods: jsii.Number(3),
})
```

See: [Monitoring Using CloudWatch Metrics](https://docs.aws.amazon.com/firehose/latest/dev/monitoring-with-cloudwatch-metrics.html)
in the *Amazon Data Firehose Developer Guide*.

## Compression

Your data can automatically be compressed when it is delivered to S3 as either a final or
an intermediary/backup destination. Supported compression formats are: gzip, Snappy,
Hadoop-compatible Snappy, and ZIP, except for Redshift destinations, where Snappy
(regardless of Hadoop-compatibility) and ZIP are not supported. By default, data is
delivered to S3 without compression.

```go
// Compress data delivered to S3 using Snappy
var bucket bucket

s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	Compression: destinations.Compression_SNAPPY(),
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: s3Destination,
})
```

## Buffering

Incoming data is buffered before it is delivered to the specified destination. The
delivery stream will wait until the amount of incoming data has exceeded some threshold
(the "buffer size") or until the time since the last data delivery occurred exceeds some
threshold (the "buffer interval"), whichever happens first. You can configure these
thresholds based on the capabilities of the destination and your use-case. By default, the
buffer size is 5 MiB and the buffer interval is 5 minutes.

```go
// Increase the buffer interval and size to 10 minutes and 8 MiB, respectively
var bucket bucket

destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	BufferingInterval: awscdk.Duration_Minutes(jsii.Number(10)),
	BufferingSize: awscdk.Size_Mebibytes(jsii.Number(8)),
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destination,
})
```

See: [Data Delivery Frequency](https://docs.aws.amazon.com/firehose/latest/dev/basic-deliver.html#frequency)
in the *Amazon Data Firehose Developer Guide*.

Zero buffering, where Amazon Data Firehose stream can be configured to not buffer data before delivery, is supported by
setting the "buffer interval" to 0.

```go
// Setup zero buffering
var bucket bucket

destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	BufferingInterval: awscdk.Duration_Seconds(jsii.Number(0)),
})
firehose.NewDeliveryStream(this, jsii.String("ZeroBufferDeliveryStream"), &DeliveryStreamProps{
	Destination: destination,
})
```

See: [Buffering Hints](https://docs.aws.amazon.com/firehose/latest/dev/buffering-hints.html).

## Destination Encryption

Your data can be automatically encrypted when it is delivered to S3 as a final or an
intermediary/backup destination. Amazon Data Firehose supports Amazon S3 server-side
encryption with AWS Key Management Service (AWS KMS) for encrypting delivered data in
Amazon S3. You can choose to not encrypt the data or to encrypt with a key from the list
of AWS KMS keys that you own. For more information,
see [Protecting Data Using Server-Side Encryption with AWS KMS–Managed Keys (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
By default, encryption isn’t directly enabled on the delivery stream; instead, it uses the default encryption settings of the destination S3 bucket.

```go
var bucket bucket
var key key

destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	EncryptionKey: key,
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destination,
})
```

## Backup

A delivery stream can be configured to back up data to S3 that it attempted to deliver to
the configured destination. Backed up data can be all the data that the delivery stream
attempted to deliver or just data that it failed to deliver (Redshift and S3 destinations
can only back up all data). CDK can create a new S3 bucket where it will back up data, or
you can provide a bucket where data will be backed up. You can also provide a prefix under
which your backed-up data will be placed within the bucket. By default, source data is not
backed up to S3.

```go
// Enable backup of all source records (to an S3 bucket created by CDK).
var bucket bucket
// Explicitly provide an S3 bucket to which all source records will be backed up.
var backupBucket bucket

firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All"), &DeliveryStreamProps{
	Destination:
	destinations.NewS3Bucket(bucket, &S3BucketProps{
		S3Backup: &DestinationS3BackupProps{
			Mode: destinations.BackupMode_ALL,
		},
	}),
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Bucket"), &DeliveryStreamProps{
	Destination:
	destinations.NewS3Bucket(bucket, &S3BucketProps{
		S3Backup: &DestinationS3BackupProps{
			Bucket: backupBucket,
		},
	}),
})
// Explicitly provide an S3 prefix under which all source records will be backed up.
// Explicitly provide an S3 prefix under which all source records will be backed up.
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Prefix"), &DeliveryStreamProps{
	Destination:
	destinations.NewS3Bucket(bucket, &S3BucketProps{
		S3Backup: &DestinationS3BackupProps{
			Mode: destinations.BackupMode_ALL,
			DataOutputPrefix: jsii.String("mybackup"),
		},
	}),
})
```

If any Data Processing or Transformation is configured on your Delivery Stream, the source
records will be backed up in their original format.

## Data Processing/Transformation

Data can be transformed before being delivered to destinations. There are two types of
data processing for delivery streams: record transformation with AWS Lambda, and record
format conversion using a schema stored in an AWS Glue table. If both types of data
processing are configured, then the Lambda transformation is performed first. By default,
no data processing occurs. This construct library currently only supports data
transformation with AWS Lambda. See [#15501](https://github.com/aws/aws-cdk/issues/15501)
to track the status of adding support for record format conversion.

### Data transformation with AWS Lambda

To transform the data, Amazon Data Firehose will call a Lambda function that you provide
and deliver the data returned in place of the source record. The function must return a
result that contains records in a specific format, including the following fields:

* `recordId` -- the ID of the input record that corresponds the results.
* `result` -- the status of the transformation of the record: "Ok" (success), "Dropped"
  (not processed intentionally), or "ProcessingFailed" (not processed due to an error).
* `data` -- the transformed data, Base64-encoded.

The data is buffered up to 1 minute and up to 3 MiB by default before being sent to the
function, but can be configured using `bufferInterval` and `bufferSize`
in the processor configuration (see: [Buffering](#buffering)). If the function invocation
fails due to a network timeout or because of hitting an invocation limit, the invocation
is retried 3 times by default, but can be configured using `retries` in the processor
configuration.

```go
var bucket bucket
// Provide a Lambda function that will transform records before delivery, with custom
// buffering and retry configuration
lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("process-records"))),
})
lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &DataProcessorProps{
	BufferInterval: awscdk.Duration_Minutes(jsii.Number(5)),
	BufferSize: awscdk.Size_Mebibytes(jsii.Number(5)),
	Retries: jsii.Number(5),
})
s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	Processor: lambdaProcessor,
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: s3Destination,
})
```

```go
import path "github.com/aws-samples/dummy/path"
import "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
import "github.com/aws/aws-cdk-go/awscdk"
import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
import logs "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
import "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"

app := cdk.NewApp()

stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))

bucket := s3.NewBucket(stack, jsii.String("FirehoseDeliveryStreamS3AllPropertiesBucket"), &BucketProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	AutoDeleteObjects: jsii.Boolean(true),
})

backupBucket := s3.NewBucket(stack, jsii.String("FirehoseDeliveryStreamS3AllPropertiesBackupBucket"), &BucketProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	AutoDeleteObjects: jsii.Boolean(true),
})
logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
})

dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &NodejsFunctionProps{
	Entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
	Timeout: cdk.Duration_Minutes(jsii.Number(1)),
})

processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &DataProcessorProps{
	BufferInterval: cdk.Duration_Seconds(jsii.Number(60)),
	BufferSize: cdk.Size_Mebibytes(jsii.Number(1)),
	Retries: jsii.Number(1),
})

key := kms.NewKey(stack, jsii.String("Key"), &KeyProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
})

backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &KeyProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
})

deliveryStream := firehose.NewDeliveryStream(stack, jsii.String("DeliveryStream"), &DeliveryStreamProps{
	Destination: destinations.NewS3Bucket(bucket, &S3BucketProps{
		LoggingConfig: destinations.NewEnableLogging(logGroup),
		Processor: processor,
		Compression: destinations.Compression_GZIP(),
		DataOutputPrefix: jsii.String("regularPrefix"),
		ErrorOutputPrefix: jsii.String("errorPrefix"),
		BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
		BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
		EncryptionKey: key,
		S3Backup: &DestinationS3BackupProps{
			Mode: destinations.BackupMode_ALL,
			Bucket: backupBucket,
			Compression: destinations.Compression_ZIP(),
			DataOutputPrefix: jsii.String("backupPrefix"),
			ErrorOutputPrefix: jsii.String("backupErrorPrefix"),
			BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
			BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
			EncryptionKey: backupKey,
		},
	}),
})

firehose.NewDeliveryStream(stack, jsii.String("ZeroBufferingDeliveryStream"), &DeliveryStreamProps{
	Destination: destinations.NewS3Bucket(bucket, &S3BucketProps{
		Compression: destinations.Compression_GZIP(),
		DataOutputPrefix: jsii.String("regularPrefix"),
		ErrorOutputPrefix: jsii.String("errorPrefix"),
		BufferingInterval: cdk.Duration_*Seconds(jsii.Number(0)),
	}),
})

testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("integ-tests"), &IntegTestProps{
	TestCases: []stack{
		stack,
	},
	Regions: []*string{
		jsii.String("us-east-1"),
	},
})

testCase.Assertions.AwsApiCall(jsii.String("Firehose"), jsii.String("putRecord"), map[string]interface{}{
	"DeliveryStreamName": deliveryStream.deliveryStreamName,
	"Record": map[string]*string{
		"Data": jsii.String("testData123"),
	},
})

s3ApiCall := testCase.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("listObjectsV2"), map[string]interface{}{
	"Bucket": bucket.bucketName,
	"MaxKeys": jsii.Number(1),
}).Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"KeyCount": jsii.Number(1),
})).WaitForAssertions(&WaiterStateMachineOptions{
	Interval: cdk.Duration_*Seconds(jsii.Number(30)),
	TotalTimeout: cdk.Duration_*Minutes(jsii.Number(10)),
})

if s3ApiCall instanceof awscdkintegtestsalpha.AwsApiCall && s3ApiCall.WaiterProvider {
	s3ApiCall.WaiterProvider.AddToRolePolicy(map[string]interface{}{
		"Effect": jsii.String("Allow"),
		"Action": []*string{
			jsii.String("s3:GetObject"),
			jsii.String("s3:ListBucket"),
		},
		"Resource": []*string{
			jsii.String("*"),
		},
	})
}
```

See: [Data Transformation](https://docs.aws.amazon.com/firehose/latest/dev/data-transformation.html)
in the *Amazon Data Firehose Developer Guide*.

## Specifying an IAM role

The DeliveryStream class automatically creates IAM service roles with all the minimum
necessary permissions for Amazon Data Firehose to access the resources referenced by your
delivery stream. One service role is created for the delivery stream that allows Amazon
Data Firehose to read from a Kinesis data stream (if one is configured as the delivery
stream source) and for server-side encryption. Note that if the DeliveryStream is created
without specifying a `source` or `encryptionKey`, this role is not created as it is not needed.

Another service role is created for each destination, which gives Amazon Data Firehose write
access to the destination resource, as well as the ability to invoke data transformers and
read schemas for record format conversion. If you wish, you may specify your own IAM role for
either the delivery stream or the destination service role, or both. It must have the correct
trust policy (it must allow Amazon Data Firehose to assume it) or delivery stream creation or
data delivery will fail. Other required permissions to destination resources, encryption keys, etc.,
will be provided automatically.

```go
// Specify the roles created above when defining the destination and delivery stream.
var bucket bucket
// Create service roles for the delivery stream and destination.
// These can be used for other purposes and granted access to different resources.
// They must include the Amazon Data Firehose service principal in their trust policies.
// Two separate roles are shown below, but the same role can be used for both purposes.
deliveryStreamRole := iam.NewRole(this, jsii.String("Delivery Stream Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
})
destinationRole := iam.NewRole(this, jsii.String("Destination Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
})
destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
	Role: destinationRole,
})
firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destination: destination,
	Role: deliveryStreamRole,
})
```

See [Controlling Access](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html)
in the *Amazon Data Firehose Developer Guide*.

## Granting application access to a delivery stream

IAM roles, users or groups which need to be able to work with delivery streams should be
granted IAM permissions.

Any object that implements the `IGrantable` interface (i.e., has an associated principal)
can be granted permissions to a delivery stream by calling:

* `grantPutRecords(principal)` - grants the principal the ability to put records onto the
  delivery stream
* `grant(principal, ...actions)` - grants the principal permission to a custom set of
  actions

```go
// Give the role permissions to write data to the delivery stream
var deliveryStream deliveryStream
lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})
deliveryStream.grantPutRecords(lambdaRole)
```

The following write permissions are provided to a service principal by the
`grantPutRecords()` method:

* `firehose:PutRecord`
* `firehose:PutRecordBatch`

## Granting a delivery stream access to a resource

Conversely to the above, Amazon Data Firehose requires permissions in order for delivery
streams to interact with resources that you own. For example, if an S3 bucket is specified
as a destination of a delivery stream, the delivery stream must be granted permissions to
put and get objects from the bucket. When using the built-in AWS service destinations
found in the `@aws-cdk/aws-kinesisfirehose-destinations-alpha` module, the CDK grants the
permissions automatically. However, custom or third-party destinations may require custom
permissions. In this case, use the delivery stream as an `IGrantable`, as follows:

```go
var deliveryStream deliveryStream
fn := lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromInline(jsii.String("exports.handler = (event) => {}")),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
})
fn.GrantInvoke(deliveryStream)
```
