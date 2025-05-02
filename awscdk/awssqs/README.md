# Amazon Simple Queue Service Construct Library

Amazon Simple Queue Service (SQS) is a fully managed message queuing service that
enables you to decouple and scale microservices, distributed systems, and serverless
applications. SQS eliminates the complexity and overhead associated with managing and
operating message oriented middleware, and empowers developers to focus on differentiating work.
Using SQS, you can send, store, and receive messages between software components at any volume,
without losing messages or requiring other services to be available.

## Installation

Import to your project:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"
```

## Basic usage

Here's how to add a basic queue to your application:

```go
sqs.NewQueue(this, jsii.String("Queue"))
```

## Encryption

By default queues are encrypted using SSE-SQS. If you want to change the encryption mode, set the `encryption` property.
The following encryption modes are supported:

* KMS key that SQS manages for you
* KMS key that you can managed yourself
* Server-side encryption managed by SQS (SSE-SQS)
* Unencrypted

To learn more about SSE-SQS on Amazon SQS, please visit the
[Amazon SQS documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html).

```go
// Use managed key
// Use managed key
sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	Encryption: sqs.QueueEncryption_KMS_MANAGED,
})

// Use custom key
myKey := kms.NewKey(this, jsii.String("Key"))

sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	Encryption: sqs.QueueEncryption_KMS,
	EncryptionMasterKey: myKey,
})

// Use SQS managed server side encryption (SSE-SQS)
// Use SQS managed server side encryption (SSE-SQS)
sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	Encryption: sqs.QueueEncryption_SQS_MANAGED,
})

// Unencrypted queue
// Unencrypted queue
sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	Encryption: sqs.QueueEncryption_UNENCRYPTED,
})
```

## Encryption in transit

If you want to enforce encryption of data in transit, set the `enforceSSL` property to `true`.
A resource policy statement that allows only encrypted connections over HTTPS (TLS)
will be added to the queue.

```go
sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	EnforceSSL: jsii.Boolean(true),
})
```

## First-In-First-Out (FIFO) queues

FIFO queues give guarantees on the order in which messages are dequeued, and have additional
features in order to help guarantee exactly-once processing. For more information, see
the SQS manual. Note that FIFO queues are not available in all AWS regions.

A queue can be made a FIFO queue by either setting `fifo: true`, giving it a name which ends
in `".fifo"`, or by enabling a FIFO specific feature such as: content-based deduplication,
deduplication scope or fifo throughput limit.

## Dead letter source queues permission

You can configure the permission settings for queues that can designate the created queue as their dead-letter queue using the `redriveAllowPolicy` attribute.

By default, all queues within the same account and region are permitted as source queues.

```go
var sourceQueue iQueue


// Only the sourceQueue can specify this queue as the dead-letter queue.
queue1 := sqs.NewQueue(this, jsii.String("Queue2"), &QueueProps{
	RedriveAllowPolicy: &RedriveAllowPolicy{
		SourceQueues: []*iQueue{
			sourceQueue,
		},
	},
})

// No source queues can specify this queue as the dead-letter queue.
queue2 := sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
	RedriveAllowPolicy: &RedriveAllowPolicy{
		RedrivePermission: sqs.RedrivePermission_DENY_ALL,
	},
})
```
