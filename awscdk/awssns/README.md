# Amazon Simple Notification Service Construct Library

Add an SNS Topic to your stack:

```go
topic := sns.NewTopic(this, jsii.String("Topic"), &TopicProps{
	DisplayName: jsii.String("Customer subscription topic"),
})
```

Add a FIFO SNS topic with content-based de-duplication to your stack:

```go
topic := sns.NewTopic(this, jsii.String("Topic"), &TopicProps{
	ContentBasedDeduplication: jsii.Boolean(true),
	DisplayName: jsii.String("Customer subscription topic"),
	Fifo: jsii.Boolean(true),
})
```

Note that FIFO topics require a topic name to be provided. The required `.fifo` suffix will be automatically generated and added to the topic name if it is not explicitly provided.

## Subscriptions

Various subscriptions can be added to the topic by calling the
`.addSubscription(...)` method on the topic. It accepts a *subscription* object,
default implementations of which can be found in the
`aws-cdk-lib/aws-sns-subscriptions` package:

Add an HTTPS Subscription to your topic:

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.AddSubscription(subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/")))
```

Subscribe a queue to the topic:

```go
var queue queue

myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.AddSubscription(subscriptions.NewSqsSubscription(queue))
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
myTopic.AddSubscription(subscriptions.NewLambdaSubscription(fn, &LambdaSubscriptionProps{
	FilterPolicy: map[string]subscriptionFilter{
		"color": sns.*subscriptionFilter_stringFilter(&StringConditions{
			"allowlist": []*string{
				jsii.String("red"),
				jsii.String("orange"),
			},
			"matchPrefixes": []*string{
				jsii.String("bl"),
			},
			"matchSuffixes": []*string{
				jsii.String("ue"),
			},
		}),
		"size": sns.*subscriptionFilter_stringFilter(&StringConditions{
			"denylist": []*string{
				jsii.String("small"),
				jsii.String("medium"),
			},
		}),
		"price": sns.*subscriptionFilter_numericFilter(&NumericConditions{
			"between": &BetweenCondition{
				"start": jsii.Number(100),
				"stop": jsii.Number(200),
			},
			"greaterThan": jsii.Number(300),
		}),
		"store": sns.*subscriptionFilter_existsFilter(),
	},
}))
```

#### Payload-based filtering

To filter messages based on the payload or body of the message, use the `filterPolicyWithMessageBody` property. This type of filter policy supports creating filters on nested objects.

Example with a Lambda subscription:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var fn function


myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

// Lambda should receive only message matching the following conditions on message body:
// color: 'red' or 'orange'
myTopic.AddSubscription(subscriptions.NewLambdaSubscription(fn, &LambdaSubscriptionProps{
	FilterPolicyWithMessageBody: map[string]filterOrPolicy{
		"background": sns.*filterOrPolicy_policy(map[string]*filterOrPolicy{
			"color": sns.*filterOrPolicy_filter(sns.SubscriptionFilter_stringFilter(&StringConditions{
				"allowlist": []*string{
					jsii.String("red"),
					jsii.String("orange"),
				},
			})),
		}),
	},
}))
```

### Example of Firehose Subscription

```go
import "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
var stream deliveryStream


topic := sns.NewTopic(this, jsii.String("Topic"))

sns.NewSubscription(this, jsii.String("Subscription"), &SubscriptionProps{
	Topic: Topic,
	Endpoint: stream.DeliveryStreamArn,
	Protocol: sns.SubscriptionProtocol_FIREHOSE,
	SubscriptionRoleArn: jsii.String("SAMPLE_ARN"),
})
```

## DLQ setup for SNS Subscription

CDK can attach provided Queue as DLQ for your SNS subscription.
See the [SNS DLQ configuration docs](https://docs.aws.amazon.com/sns/latest/dg/sns-configure-dead-letter-queue.html) for more information about this feature.

Example of usage with user provided DLQ.

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
dlQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"), &QueueProps{
	QueueName: jsii.String("MySubscription_DLQ"),
	RetentionPeriod: awscdk.Duration_Days(jsii.Number(14)),
})

sns.NewSubscription(this, jsii.String("Subscription"), &SubscriptionProps{
	Endpoint: jsii.String("endpoint"),
	Protocol: sns.SubscriptionProtocol_LAMBDA,
	Topic: Topic,
	DeadLetterQueue: dlQueue,
})
```

## CloudWatch Event Rule Target

SNS topics can be used as targets for CloudWatch event rules.

Use the `aws-cdk-lib/aws-events-targets.SnsTopic`:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var repo repository

myTopic := sns.NewTopic(this, jsii.String("Topic"))

repo.onCommit(jsii.String("OnCommit"), &OnCommitOptions{
	Target: targets.NewSnsTopic(myTopic),
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
topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &TopicPolicyProps{
	Topics: []iTopic{
		topic,
	},
})

topicPolicy.Document.AddStatements(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("sns:Subscribe"),
	},
	Principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
	Resources: []*string{
		topic.TopicArn,
	},
}))
```

A policy document can also be passed on `TopicPolicy` construction

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
	AssignSids: jsii.Boolean(true),
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("sns:Subscribe"),
			},
			Principals: []iPrincipal{
				iam.NewAnyPrincipal(),
			},
			Resources: []*string{
				topic.TopicArn,
			},
		}),
	},
})

topicPolicy := sns.NewTopicPolicy(this, jsii.String("Policy"), &TopicPolicyProps{
	Topics: []iTopic{
		topic,
	},
	PolicyDocument: PolicyDocument,
})
```

### Enforce encryption of data in transit when publishing to a topic

You can enforce SSL when creating a topic policy by setting the `enforceSSL` flag:

```go
topic := sns.NewTopic(this, jsii.String("Topic"))
policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
	AssignSids: jsii.Boolean(true),
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("sns:Publish"),
			},
			Principals: []iPrincipal{
				iam.NewServicePrincipal(jsii.String("s3.amazonaws.com")),
			},
			Resources: []*string{
				topic.TopicArn,
			},
		}),
	},
})

topicPolicy := sns.NewTopicPolicy(this, jsii.String("Policy"), &TopicPolicyProps{
	Topics: []iTopic{
		topic,
	},
	PolicyDocument: PolicyDocument,
	EnforceSSL: jsii.Boolean(true),
})
```

Similiarly you can enforce SSL by setting the `enforceSSL` flag on the topic:

```go
topic := sns.NewTopic(this, jsii.String("TopicAddPolicy"), &TopicProps{
	EnforceSSL: jsii.Boolean(true),
})

topic.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Principals: []iPrincipal{
		iam.NewServicePrincipal(jsii.String("s3.amazonaws.com")),
	},
	Actions: []*string{
		jsii.String("sns:Publish"),
	},
	Resources: []*string{
		topic.TopicArn,
	},
}))
```

## Delivery status logging

Amazon SNS provides support to log the delivery status of notification messages sent to topics with the following Amazon SNS endpoints:

* HTTP
* Amazon Kinesis Data Firehose
* AWS Lambda
* Platform application endpoint
* Amazon Simple Queue Service

Example with a delivery status logging configuration for SQS:

```go
var role role

topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
	LoggingConfigs: []loggingConfig{
		&loggingConfig{
			Protocol: sns.LoggingProtocol_SQS,
			FailureFeedbackRole: role,
			SuccessFeedbackRole: role,
			SuccessFeedbackSampleRate: jsii.Number(50),
		},
	},
})
```

A delivery status logging configuration can also be added to your topic by `addLoggingConfig` method:

```go
var role role

topic := sns.NewTopic(this, jsii.String("MyTopic"))

topic.AddLoggingConfig(&LoggingConfig{
	Protocol: sns.LoggingProtocol_SQS,
	FailureFeedbackRole: role,
	SuccessFeedbackRole: role,
	SuccessFeedbackSampleRate: jsii.Number(50),
})
```

Note that valid values for `successFeedbackSampleRate` are integer between 0-100.

## Archive Policy

Message archiving provides the ability to archive a single copy of all messages published to your topic.
You can store published messages within your topic by enabling the message archive policy on the topic, which enables message archiving for all subscriptions linked to that topic.
Messages can be archived for a minimum of one day to a maximum of 365 days.

Example with a archive policy for SQS:

```go
var role role

topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
	Fifo: jsii.Boolean(true),
	MessageRetentionPeriodInDays: jsii.Number(7),
})
```

**Note**: The `messageRetentionPeriodInDays` property is only available for FIFO topics.
