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
var sourceQueue IQueue


// Only the sourceQueue can specify this queue as the dead-letter queue.
queue1 := sqs.NewQueue(this, jsii.String("Queue2"), &QueueProps{
	RedriveAllowPolicy: &RedriveAllowPolicy{
		SourceQueues: []IQueue{
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

## Monitoring

SQS metrics are available as `metric*` methods on a queue; `metric()` returns any metric by name:

```go
queue := sqs.NewQueue(this, jsii.String("Queue"))

queue.metricApproximateAgeOfOldestMessage().CreateAlarm(this, jsii.String("MessagesTooOld"), &CreateAlarmOptions{
	Threshold: awscdk.Duration_Minutes(jsii.Number(15)).ToSeconds(),
	EvaluationPeriods: jsii.Number(3),
})
```

`metricApproximateNumberOfMessagesOutstanding()` returns
`ApproximateNumberOfMessagesVisible + ApproximateNumberOfMessagesNotVisible` as a metric math
expression: messages waiting to be picked up, plus messages received but not yet deleted.

### Autoscaling consumers on queue depth

Scaling a worker fleet on queue depth needs a different metric in each direction:

* **Scale out on `ApproximateNumberOfMessagesVisible`** — work nobody has started yet. An in-flight
  message is already owned by a consumer, so adding capacity for it produces an idle consumer.
* **Scale in on `metricApproximateNumberOfMessagesOutstanding()`** — everything still owed.
  Receiving a message moves it from `Visible` to `NotVisible`, so a policy watching `Visible` alone
  cannot tell a consumer that just picked up work from one that finished it, and can terminate a
  consumer mid-message. The message reappears only after its visibility timeout, so the longer
  consumers hold messages, the longer that work stalls.

That means two one-sided policies; the `change: 0` step keeps each from acting in the other
direction:

```go
var service FargateService

queue := sqs.NewQueue(this, jsii.String("Queue"))

taskCount := service.AutoScaleTaskCount(&EnableScalingProps{
	MinCapacity: jsii.Number(1),
	MaxCapacity: jsii.Number(10),
})

taskCount.ScaleOnMetric(jsii.String("ScaleOutOnWaitingWork"), &BasicStepScalingPolicyProps{
	Metric: queue.metricApproximateNumberOfMessagesVisible(&MetricOptions{
		Period: awscdk.Duration_Minutes(jsii.Number(1)),
	}),
	ScalingSteps: []ScalingInterval{
		&ScalingInterval{
			Upper: jsii.Number(30),
			Change: jsii.Number(0),
		},
		&ScalingInterval{
			Lower: jsii.Number(30),
			Change: +jsii.Number(1),
		},
	},
	AdjustmentType: appscaling.AdjustmentType_CHANGE_IN_CAPACITY,
})

// Remove a task only when nothing is outstanding, so it cannot be holding a message.
taskCount.ScaleOnMetric(jsii.String("ScaleInOnOutstandingWork"), &BasicStepScalingPolicyProps{
	Metric: queue.MetricApproximateNumberOfMessagesOutstanding(&MetricOptions{
		Period: awscdk.Duration_*Minutes(jsii.Number(1)),
	}),
	ScalingSteps: []ScalingInterval{
		&ScalingInterval{
			Upper: jsii.Number(0),
			Change: -jsii.Number(1),
		},
		&ScalingInterval{
			Lower: jsii.Number(0),
			Change: jsii.Number(0),
		},
	},
	AdjustmentType: appscaling.AdjustmentType_CHANGE_IN_CAPACITY,
})
```

Caveats:

* **Target tracking rejects this metric.** It accepts only direct metrics, so
  `scaleToTrackCustomMetric()` throws `Only direct metrics are supported for Target Tracking`
  ([aws-cdk#20659](https://github.com/aws/aws-cdk/issues/20659)).
* `ApproximateNumberOfMessagesNotVisible` can briefly report non-zero on an empty queue if an SQS
  storage server is unavailable. Evaluate several consecutive datapoints, especially when scaling in
  to zero.
* Neither term counts delayed messages. Add `metricApproximateNumberOfMessagesDelayed()` if you use
  delay queues or `DelaySeconds`.
