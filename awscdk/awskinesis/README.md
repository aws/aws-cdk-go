# Amazon Kinesis Construct Library

[Amazon Kinesis](https://docs.aws.amazon.com/streams/latest/dev/introduction.html) provides collection and processing of large
[streams](https://aws.amazon.com/streaming-data/) of data records in real time. Kinesis data streams can be used for rapid and continuous data
intake and aggregation.

## Table Of Contents

* [Streams](#streams)

  * [Encryption](#encryption)
  * [Import](#import)
  * [Permission Grants](#permission-grants)

    * [Read Permissions](#read-permissions)
    * [Write Permissions](#write-permissions)
    * [Custom Permissions](#custom-permissions)
  * [Metrics](#metrics)

    * [Shard-level Metrics](#shard-level-metrics)
* [Stream Consumers](#stream-consumers)

  * [Read Permissions](#read-permissions-1)
* [Resource Policy](#resource-policy)

## Streams

Amazon Kinesis Data Streams ingests a large amount of data in real time, durably stores the data, and makes the data available for consumption.

Using the CDK, a new Kinesis stream can be created as part of the stack using the construct's constructor. You may specify the `streamName` to give
your own identifier to the stream. If not, CloudFormation will generate a name.

```go
kinesis.NewStream(this, jsii.String("MyFirstStream"), &StreamProps{
	StreamName: jsii.String("my-awesome-stream"),
})
```

You can also specify properties such as `shardCount` to indicate how many shards the stream should choose and a `retentionPeriod`
to specify how long the data in the shards should remain accessible.
Read more at [Creating and Managing Streams](https://docs.aws.amazon.com/streams/latest/dev/working-with-streams.html)

```go
kinesis.NewStream(this, jsii.String("MyFirstStream"), &StreamProps{
	StreamName: jsii.String("my-awesome-stream"),
	ShardCount: jsii.Number(3),
	RetentionPeriod: awscdk.Duration_Hours(jsii.Number(48)),
})
```

### Encryption

[Stream encryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html) enables
server-side encryption using an AWS KMS key for a specified stream.

Encryption is enabled by default on your stream with the master key owned by Kinesis Data Streams in regions where it is supported.

```go
kinesis.NewStream(this, jsii.String("MyEncryptedStream"))
```

You can enable encryption on your stream with a user-managed key by specifying the `encryption` property.
A KMS key will be created for you and associated with the stream.

```go
kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
	Encryption: kinesis.StreamEncryption_KMS,
})
```

You can also supply your own external KMS key to use for stream encryption by specifying the `encryptionKey` property.

```go
key := kms.NewKey(this, jsii.String("MyKey"))

kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
	Encryption: kinesis.StreamEncryption_KMS,
	EncryptionKey: key,
})
```

### Import

Any Kinesis stream that has been created outside the stack can be imported into your CDK app.

Streams can be imported by their ARN via the `Stream.fromStreamArn()` API

```go
importedStream := kinesis.Stream_FromStreamArn(this, jsii.String("ImportedStream"), jsii.String("arn:aws:kinesis:us-east-2:123456789012:stream/f3j09j2230j"))
```

Encrypted Streams can also be imported by their attributes via the `Stream.fromStreamAttributes()` API

```go
importedStream := kinesis.Stream_FromStreamAttributes(this, jsii.String("ImportedEncryptedStream"), &StreamAttributes{
	StreamArn: jsii.String("arn:aws:kinesis:us-east-2:123456789012:stream/f3j09j2230j"),
	EncryptionKey: kms.Key_FromKeyArn(this, jsii.String("key"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")),
})
```

### Permission Grants

IAM roles, users or groups which need to be able to work with Amazon Kinesis streams at runtime should be granted IAM permissions.

Any object that implements the `IGrantable` interface (has an associated principal) can be granted permissions by calling:

* `grantRead(principal)` - grants the principal read access
* `grantWrite(principal)` - grants the principal write permissions to a Stream
* `grantReadWrite(principal)` - grants principal read and write permissions

#### Read Permissions

Grant `read` access to a stream by calling the `grantRead()` API.
If the stream has an encryption key, read permissions will also be granted to the key.

```go
lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
	Description: jsii.String("Example role..."),
})

stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
	Encryption: kinesis.StreamEncryption_KMS,
})

// give lambda permissions to read stream
stream.grantRead(lambdaRole)
```

The following read permissions are provided to a service principal by the `grantRead()` API:

* `kinesis:DescribeStreamSummary`
* `kinesis:GetRecords`
* `kinesis:GetShardIterator`
* `kinesis:ListShards`
* `kinesis:SubscribeToShard`

#### Write Permissions

Grant `write` permissions to a stream is provided by calling the `grantWrite()` API.
If the stream has an encryption key, write permissions will also be granted to the key.

```go
lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
	Description: jsii.String("Example role..."),
})

stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
	Encryption: kinesis.StreamEncryption_KMS,
})

// give lambda permissions to write to stream
stream.grantWrite(lambdaRole)
```

The following write permissions are provided to a service principal by the `grantWrite()` API:

* `kinesis:ListShards`
* `kinesis:PutRecord`
* `kinesis:PutRecords`

#### Custom Permissions

You can add any set of permissions to a stream by calling the `grant()` API.

```go
user := iam.NewUser(this, jsii.String("MyUser"))

stream := kinesis.NewStream(this, jsii.String("MyStream"))

// give my user permissions to list shards
stream.grant(user, jsii.String("kinesis:ListShards"))
```

### Metrics

You can use common metrics from your stream to create alarms and/or dashboards. The `stream.metric('MetricName')` method creates a metric with the stream namespace and dimension. You can also use pre-define methods like `stream.metricGetRecordsSuccess()`. To find out more about Kinesis metrics check [Monitoring the Amazon Kinesis Data Streams Service with Amazon CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html).

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"))

// Using base metric method passing the metric name
stream.metric(jsii.String("GetRecords.Success"))

// using pre-defined metric method
stream.metricGetRecordsSuccess()

// using pre-defined and overriding the statistic
stream.metricGetRecordsSuccess(&MetricOptions{
	Statistic: jsii.String("Maximum"),
})
```

#### Shard-level Metrics

You can enable enhanced shard-level metrics for your Kinesis stream to get detailed monitoring of individual shards. Shard-level metrics provide more granular insights into the performance and health of your stream.

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"), &StreamProps{
	ShardLevelMetrics: []shardLevelMetrics{
		kinesis.*shardLevelMetrics_ALL,
	},
})
```

You can also specify individual metrics that you want to monitor:

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"), &StreamProps{
	ShardLevelMetrics: []shardLevelMetrics{
		kinesis.*shardLevelMetrics_INCOMING_BYTES,
		kinesis.*shardLevelMetrics_INCOMING_RECORDS,
		kinesis.*shardLevelMetrics_ITERATOR_AGE_MILLISECONDS,
		kinesis.*shardLevelMetrics_OUTGOING_BYTES,
		kinesis.*shardLevelMetrics_OUTGOING_RECORDS,
		kinesis.*shardLevelMetrics_READ_PROVISIONED_THROUGHPUT_EXCEEDED,
		kinesis.*shardLevelMetrics_WRITE_PROVISIONED_THROUGHPUT_EXCEEDED,
	},
})
```

Available shard-level metrics include:

* `INCOMING_BYTES` - The number of bytes successfully put to the shard
* `INCOMING_RECORDS` - The number of records successfully put to the shard
* `ITERATOR_AGE_MILLISECONDS` - The age of the last record in all GetRecords calls made against a shard
* `OUTGOING_BYTES` - The number of bytes retrieved from the shard
* `OUTGOING_RECORDS` - The number of records retrieved from the shard
* `READ_PROVISIONED_THROUGHPUT_EXCEEDED` - The number of GetRecords calls throttled for the shard
* `WRITE_PROVISIONED_THROUGHPUT_EXCEEDED` - The number of records rejected due to throttling for the shard
* `ALL` - All available metrics

Note: You cannot specify `ALL` together with other individual metrics. If you want all metrics, use `ALL` alone.

For more information about shard-level metrics, see [Monitoring the Amazon Kinesis Data Streams Service with Amazon CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html#kinesis-metrics-shard).

## Stream Consumers

Creating stream consumers allow consumers to receive data from the stream using enhanced fan-out at a rate of up to 2 MiB per second for every shard.
This rate is unaffected by the total number of consumers that read from the same stream.

For more information, see [Develop enhanced fan-out consumers with dedicated throughput](https://docs.aws.amazon.com/streams/latest/dev/enhanced-consumers.html).

To create and associate a stream consumer with a stream

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"))

streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
	StreamConsumerName: jsii.String("MyStreamConsumer"),
	Stream: Stream,
})
```

#### Read Permissions

Grant `read` access to a stream consumer, and the stream it is registered with, by calling the `grantRead()` API.

```go
lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
	Description: jsii.String("Example role..."),
})

stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
	Encryption: kinesis.StreamEncryption_KMS,
})
streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
	StreamConsumerName: jsii.String("MyStreamConsumer"),
	Stream: Stream,
})

// give lambda permissions to read stream via the stream consumer
streamConsumer.grantRead(lambdaRole)
```

In addition to stream's permissions, the following permissions are provided to a service principal by the `grantRead()` API:

* `kinesis:DescribeStreamConsumer`,
* `kinesis:SubscribeToShard`,

## Resource Policy

You can create a resource policy for a data stream or a stream consumer.
For more information, see [Controlling access to Amazon Kinesis Data Streams resources using IAM](https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html).

A resource policy is automatically created when `addToResourcePolicy` is called, if one doesn't already exist.

Using `addToResourcePolicy` is the simplest way to add a resource policy:

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"))
streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
	StreamConsumerName: jsii.String("MyStreamConsumer"),
	Stream: Stream,
})

// create a stream resource policy via addToResourcePolicy method
stream.addToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Resources: []*string{
		stream.StreamArn,
	},
	Actions: []*string{
		jsii.String("kinesis:GetRecords"),
	},
	Principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
}))

// create a stream consumer resource policy via addToResourcePolicy method
streamConsumer.addToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Resources: []*string{
		stream.*StreamArn,
	},
	Actions: []*string{
		jsii.String("kinesis:DescribeStreamConsumer"),
	},
	Principals: []*iPrincipal{
		iam.NewAnyPrincipal(),
	},
}))
```

You can create a resource manually by using `ResourcePolicy`.
Also, you can set a custom policy document to `ResourcePolicy`.
If not, a blank policy document will be set.

```go
stream := kinesis.NewStream(this, jsii.String("MyStream"))
streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
	StreamConsumerName: jsii.String("MyStreamConsumer"),
	Stream: Stream,
})

// create a custom policy document
policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
	AssignSids: jsii.Boolean(true),
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("kinesis:GetRecords"),
			},
			Resources: []*string{
				stream.StreamArn,
			},
			Principals: []iPrincipal{
				iam.NewAnyPrincipal(),
			},
		}),
	},
})

// create a stream resource policy manually
// create a stream resource policy manually
kinesis.NewResourcePolicy(this, jsii.String("ResourcePolicy"), &ResourcePolicyProps{
	Stream: Stream,
	PolicyDocument: PolicyDocument,
})

// create a stream consumer resource policy manually
// create a stream consumer resource policy manually
kinesis.NewResourcePolicy(this, jsii.String("ResourcePolicy"), &ResourcePolicyProps{
	StreamConsumer: StreamConsumer,
	PolicyDocument: PolicyDocument,
})
```
