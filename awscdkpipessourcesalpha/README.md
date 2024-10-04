# Amazon EventBridge Pipes Sources Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

EventBridge Pipes Sources let you create a source for a EventBridge Pipe.

For more details see the service documentation:

[Documentation](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-source.html)

## Pipe sources

Pipe sources are the starting point of a EventBridge Pipe. They are the source of the events that are sent to the pipe.

### Amazon SQS

A SQS message queue can be used as a source for a pipe. The queue will be polled for new messages and the messages will be sent to the pipe.

```go
var sourceQueue queue
var targetQueue queue


pipeSource := sources.NewSqsSource(sourceQueue)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: pipeSource,
	Target: NewSomeTarget(targetQueue),
})
```

The polling configuration can be customized:

```go
var sourceQueue queue
var targetQueue queue


pipeSource := sources.NewSqsSource(sourceQueue, &SqsSourceParameters{
	BatchSize: jsii.Number(10),
	MaximumBatchingWindow: cdk.Duration_Seconds(jsii.Number(10)),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: pipeSource,
	Target: NewSomeTarget(targetQueue),
})
```

### Amazon Kinesis

A Kinesis stream can be used as a source for a pipe. The stream will be polled for new messages and the messages will be sent to the pipe.

```go
var sourceStream stream
var targetQueue queue


pipeSource := sources.NewKinesisSource(sourceStream, &KinesisSourceParameters{
	StartingPosition: sources.KinesisStartingPosition_LATEST,
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: pipeSource,
	Target: NewSomeTarget(targetQueue),
})
```

### Amazon DynamoDB

A DynamoDB stream can be used as a source for a pipe. The stream will be polled for new messages and the messages will be sent to the pipe.

```go
var targetQueue queue
table := ddb.NewTableV2(this, jsii.String("MyTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: ddb.AttributeType_STRING,
	},
	DynamoStream: ddb.StreamViewType_NEW_IMAGE,
})

pipeSource := sources.NewDynamoDBSource(table, &DynamoDBSourceParameters{
	StartingPosition: sources.DynamoDBStartingPosition_LATEST,
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: pipeSource,
	Target: NewSomeTarget(targetQueue),
})
```
