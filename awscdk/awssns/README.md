# Amazon Simple Notification Service Construct Library

Add an SNS Topic to your stack:

```go
topic := sns.NewTopic(this, jsii.String("Topic"), &topicProps{
	displayName: jsii.String("Customer subscription topic"),
})
```

Add a FIFO SNS topic with content-based de-duplication to your stack:

```go
topic := sns.NewTopic(this, jsii.String("Topic"), &topicProps{
	contentBasedDeduplication: jsii.Boolean(true),
	displayName: jsii.String("Customer subscription topic"),
	fifo: jsii.Boolean(true),
})
```

Note that FIFO topics require a topic name to be provided. The required `.fifo` suffix will be automatically generated and added to the topic name if it is not explicitly provided.

## Subscriptions

Various subscriptions can be added to the topic by calling the
`.addSubscription(...)` method on the topic. It accepts a *subscription* object,
default implementations of which can be found in the
`@aws-cdk/aws-sns-subscriptions` package:

Add an HTTPS Subscription to your topic:

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.addSubscription(subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/")))
```

Subscribe a queue to the topic:

```go
var queue queue

myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.addSubscription(subscriptions.NewSqsSubscription(queue))
```

Note that subscriptions of queues in different accounts need to be manually confirmed by
reading the initial message from the queue and visiting the link found in it.

### Filter policy

A filter policy can be specified when subscribing an endpoint to a topic.

Example with a Lambda subscription:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var fn function


myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

// Lambda should receive only message matching the following conditions on attributes:
// color: 'red' or 'orange' or begins with 'bl'
// size: anything but 'small' or 'medium'
// price: between 100 and 200 or greater than 300
// store: attribute must be present
myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
	filterPolicy: map[string]subscriptionFilter{
		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
			"allowlist": []*string{
				jsii.String("red"),
				jsii.String("orange"),
			},
			"matchPrefixes": []*string{
				jsii.String("bl"),
			},
		}),
		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
			"denylist": []*string{
				jsii.String("small"),
				jsii.String("medium"),
			},
		}),
		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
			"between": &BetweenCondition{
				"start": jsii.Number(100),
				"stop": jsii.Number(200),
			},
			"greaterThan": jsii.Number(300),
		}),
		"store": sns.*subscriptionFilter.existsFilter(),
	},
}))
```

### Example of Firehose Subscription

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws-samples/dummy/awscdklibawskinesisfirehose"
var stream DeliveryStream


topic := sns.NewTopic(this, jsii.String("Topic"))

sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
	topic: topic,
	endpoint: stream.deliveryStreamArn,
	protocol: sns.subscriptionProtocol_FIREHOSE,
	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
})
```

## DLQ setup for SNS Subscription

CDK can attach provided Queue as DLQ for your SNS subscription.
See the [SNS DLQ configuration docs](https://docs.aws.amazon.com/sns/latest/dg/sns-configure-dead-letter-queue.html) for more information about this feature.

Example of usage with user provided DLQ.

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
dlQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"), &queueProps{
	queueName: jsii.String("MySubscription_DLQ"),
	retentionPeriod: awscdk.Duration.days(jsii.Number(14)),
})

sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
	endpoint: jsii.String("endpoint"),
	protocol: sns.subscriptionProtocol_LAMBDA,
	topic: topic,
	deadLetterQueue: dlQueue,
})
```

## CloudWatch Event Rule Target

SNS topics can be used as targets for CloudWatch event rules.

Use the `@aws-cdk/aws-events-targets.SnsTopic`:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var repo repository

myTopic := sns.NewTopic(this, jsii.String("Topic"))

repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
	target: targets.NewSnsTopic(myTopic),
})
```

This will result in adding a target to the event rule and will also modify the
topic resource policy to allow CloudWatch events to publish to the topic.

## Topic Policy

A topic policy is automatically created when `addToResourcePolicy` is called, if
one doesn't already exist. Using `addToResourcePolicy` is the simplest way to
add policies, but a `TopicPolicy` can also be created manually.

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &topicPolicyProps{
	topics: []iTopic{
		topic,
	},
})

topicPolicy.document.addStatements(iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("sns:Subscribe"),
	},
	principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
	resources: []*string{
		topic.topicArn,
	},
}))
```

A policy document can also be passed on `TopicPolicy` construction

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
policyDocument := iam.NewPolicyDocument(&policyDocumentProps{
	assignSids: jsii.Boolean(true),
	statements: []policyStatement{
		iam.NewPolicyStatement(&policyStatementProps{
			actions: []*string{
				jsii.String("sns:Subscribe"),
			},
			principals: []iPrincipal{
				iam.NewAnyPrincipal(),
			},
			resources: []*string{
				topic.topicArn,
			},
		}),
	},
})

topicPolicy := sns.NewTopicPolicy(this, jsii.String("Policy"), &topicPolicyProps{
	topics: []iTopic{
		topic,
	},
	policyDocument: policyDocument,
})
```
