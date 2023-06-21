package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class for building the FilterPolicy by avoiding union types.
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
type FilterOrPolicy interface {
	// Type switch for disambiguating between subclasses.
	Type() FilterOrPolicyType
	// Check if instance is `Filter` type.
	IsFilter() *bool
	// Check if instance is `Policy` type.
	IsPolicy() *bool
}

// The jsii proxy struct for FilterOrPolicy
type jsiiProxy_FilterOrPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_FilterOrPolicy) Type() FilterOrPolicyType {
	var returns FilterOrPolicyType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewFilterOrPolicy_Override(f FilterOrPolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.FilterOrPolicy",
		nil, // no parameters
		f,
	)
}

// Filter of MessageBody.
func FilterOrPolicy_Filter(filter SubscriptionFilter) Filter {
	_init_.Initialize()

	if err := validateFilterOrPolicy_FilterParameters(filter); err != nil {
		panic(err)
	}
	var returns Filter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.FilterOrPolicy",
		"filter",
		[]interface{}{filter},
		&returns,
	)

	return returns
}

// Policy of MessageBody.
func FilterOrPolicy_Policy(policy *map[string]FilterOrPolicy) Policy {
	_init_.Initialize()

	if err := validateFilterOrPolicy_PolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns Policy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.FilterOrPolicy",
		"policy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterOrPolicy) IsFilter() *bool {
	var returns *bool

	_jsii_.Invoke(
		f,
		"isFilter",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterOrPolicy) IsPolicy() *bool {
	var returns *bool

	_jsii_.Invoke(
		f,
		"isPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

