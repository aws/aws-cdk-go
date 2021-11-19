package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssnssubscriptions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Use an email address as a subscription target.
//
// Email subscriptions require confirmation.
//
// TODO: EXAMPLE
//
// Experimental.
type EmailSubscription interface {
	awssns.ITopicSubscription
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for EmailSubscription
type jsiiProxy_EmailSubscription struct {
	internal.Type__awssnsITopicSubscription
}

// Experimental.
func NewEmailSubscription(emailAddress *string, props *EmailSubscriptionProps) EmailSubscription {
	_init_.Initialize()

	j := jsiiProxy_EmailSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmailSubscription_Override(e EmailSubscription, emailAddress *string, props *EmailSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		e,
	)
}

// Returns a configuration for an email address to subscribe to an SNS topic.
// Experimental.
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
// Experimental.
type EmailSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
	// Indicates if the full notification JSON should be sent to the email address or just the message text.
	// Experimental.
	Json *bool `json:"json"`
}

// Use a Lambda function as a subscription target.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaSubscription interface {
	awssns.ITopicSubscription
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for LambdaSubscription
type jsiiProxy_LambdaSubscription struct {
	internal.Type__awssnsITopicSubscription
}

// Experimental.
func NewLambdaSubscription(fn awslambda.IFunction, props *LambdaSubscriptionProps) LambdaSubscription {
	_init_.Initialize()

	j := jsiiProxy_LambdaSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.LambdaSubscription",
		[]interface{}{fn, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaSubscription_Override(l LambdaSubscription, fn awslambda.IFunction, props *LambdaSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.LambdaSubscription",
		[]interface{}{fn, props},
		l,
	)
}

// Returns a configuration for a Lambda function to subscribe to an SNS topic.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type LambdaSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
}

// Use an sms address as a subscription target.
//
// TODO: EXAMPLE
//
// Experimental.
type SmsSubscription interface {
	awssns.ITopicSubscription
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SmsSubscription
type jsiiProxy_SmsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

// Experimental.
func NewSmsSubscription(phoneNumber *string, props *SmsSubscriptionProps) SmsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SmsSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSmsSubscription_Override(s SmsSubscription, phoneNumber *string, props *SmsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		s,
	)
}

// Returns a configuration used to subscribe to an SNS topic.
// Experimental.
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
// Experimental.
type SmsSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
}

// Use an SQS queue as a subscription target.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsSubscription interface {
	awssns.ITopicSubscription
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SqsSubscription
type jsiiProxy_SqsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

// Experimental.
func NewSqsSubscription(queue awssqs.IQueue, props *SqsSubscriptionProps) SqsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SqsSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsSubscription_Override(s SqsSubscription, queue awssqs.IQueue, props *SqsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		s,
	)
}

// Returns a configuration for an SQS queue to subscribe to an SNS topic.
// Experimental.
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
// Experimental.
type SqsSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	// Experimental.
	RawMessageDelivery *bool `json:"rawMessageDelivery"`
}

// Options to subscribing to an SNS topic.
// Experimental.
type SubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
}

// Use a URL as a subscription target.
//
// The message will be POSTed to the given URL.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-http-https-endpoint-as-subscriber.html
//
// Experimental.
type UrlSubscription interface {
	awssns.ITopicSubscription
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for UrlSubscription
type jsiiProxy_UrlSubscription struct {
	internal.Type__awssnsITopicSubscription
}

// Experimental.
func NewUrlSubscription(url *string, props *UrlSubscriptionProps) UrlSubscription {
	_init_.Initialize()

	j := jsiiProxy_UrlSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUrlSubscription_Override(u UrlSubscription, url *string, props *UrlSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		u,
	)
}

// Returns a configuration for a URL to subscribe to an SNS topic.
// Experimental.
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
// Experimental.
type UrlSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
	// The subscription's protocol.
	// Experimental.
	Protocol awssns.SubscriptionProtocol `json:"protocol"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	// Experimental.
	RawMessageDelivery *bool `json:"rawMessageDelivery"`
}

