package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
)

// Use a Lambda function as a subscription target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on message body:
//   // color: 'red' or 'orange'
//   myTopic.AddSubscription(subscriptions.NewLambdaSubscription(fn, &LambdaSubscriptionProps{
//   	FilterPolicyWithMessageBody: map[string]filterOrPolicy{
//   		"background": sns.*filterOrPolicy_policy(map[string]*filterOrPolicy{
//   			"color": sns.*filterOrPolicy_filter(sns.SubscriptionFilter_stringFilter(&StringConditions{
//   				"allowlist": []*string{
//   					jsii.String("red"),
//   					jsii.String("orange"),
//   				},
//   			})),
//   		}),
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

	if err := validateNewLambdaSubscriptionParameters(fn, props); err != nil {
		panic(err)
	}
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
	if err := l.validateBindParameters(topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

