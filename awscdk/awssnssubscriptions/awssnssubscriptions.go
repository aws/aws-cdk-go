package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an email address as a subscription target.
//
// Email subscriptions require confirmation.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//   emailAddress := awscdk.NewCfnParameter(this, jsii.String("email-param"))
//
//   myTopic.addSubscription(subscriptions.NewEmailSubscription(emailAddress.valueAsString))
//
type EmailSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for an email address to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for EmailSubscription
type jsiiProxy_EmailSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewEmailSubscription(emailAddress *string, props *EmailSubscriptionProps) EmailSubscription {
	_init_.Initialize()

	j := jsiiProxy_EmailSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		&j,
	)

	return &j
}

func NewEmailSubscription_Override(e EmailSubscription, emailAddress *string, props *EmailSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		e,
	)
}

func (e *jsiiProxy_EmailSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

// Options for email subscriptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   emailSubscriptionProps := &emailSubscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	json: jsii.Boolean(false),
//   }
//
type EmailSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// Indicates if the full notification JSON should be sent to the email address or just the message text.
	Json *bool `field:"optional" json:"json" yaml:"json"`
}

// Use a Lambda function as a subscription target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
type LambdaSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for a Lambda function to subscribe to an SNS topic.
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for LambdaSubscription
type jsiiProxy_LambdaSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewLambdaSubscription(fn awslambda.IFunction, props *LambdaSubscriptionProps) LambdaSubscription {
	_init_.Initialize()

	j := jsiiProxy_LambdaSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.LambdaSubscription",
		[]interface{}{fn, props},
		&j,
	)

	return &j
}

func NewLambdaSubscription_Override(l LambdaSubscription, fn awslambda.IFunction, props *LambdaSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.LambdaSubscription",
		[]interface{}{fn, props},
		l,
	)
}

func (l *jsiiProxy_LambdaSubscription) Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

// Properties for a Lambda subscription.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
type LambdaSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

// Use an sms address as a subscription target.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   myTopic.addSubscription(subscriptions.NewSmsSubscription(jsii.String("+15551231234")))
//
type SmsSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration used to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SmsSubscription
type jsiiProxy_SmsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewSmsSubscription(phoneNumber *string, props *SmsSubscriptionProps) SmsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SmsSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		&j,
	)

	return &j
}

func NewSmsSubscription_Override(s SmsSubscription, phoneNumber *string, props *SmsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		s,
	)
}

func (s *jsiiProxy_SmsSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

// Options for SMS subscriptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   smsSubscriptionProps := &smsSubscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   }
//
type SmsSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

// Use an SQS queue as a subscription target.
//
// Example:
//   var queue queue
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.addSubscription(subscriptions.NewSqsSubscription(queue))
//
type SqsSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for an SQS queue to subscribe to an SNS topic.
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SqsSubscription
type jsiiProxy_SqsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewSqsSubscription(queue awssqs.IQueue, props *SqsSubscriptionProps) SqsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SqsSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsSubscription_Override(s SqsSubscription, queue awssqs.IQueue, props *SqsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsSubscription) Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

// Properties for an SQS subscription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   sqsSubscriptionProps := &sqsSubscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	rawMessageDelivery: jsii.Boolean(false),
//   }
//
type SqsSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
}

// Options to subscribing to an SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   subscriptionProps := &subscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   }
//
type SubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

// Use a URL as a subscription target.
//
// The message will be POSTed to the given URL.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.addSubscription(subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/")))
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-http-https-endpoint-as-subscriber.html
//
type UrlSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for a URL to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for UrlSubscription
type jsiiProxy_UrlSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewUrlSubscription(url *string, props *UrlSubscriptionProps) UrlSubscription {
	_init_.Initialize()

	j := jsiiProxy_UrlSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

func NewUrlSubscription_Override(u UrlSubscription, url *string, props *UrlSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		u,
	)
}

func (u *jsiiProxy_UrlSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

// Options for URL subscriptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   urlSubscriptionProps := &urlSubscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	protocol: awscdk.Aws_sns.subscriptionProtocol_HTTP,
//   	rawMessageDelivery: jsii.Boolean(false),
//   }
//
type UrlSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The subscription's protocol.
	Protocol awssns.SubscriptionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
}

