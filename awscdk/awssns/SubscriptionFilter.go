package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A subscription filter for an attribute.
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
//   myTopic.AddSubscription(subscriptions.NewLambdaSubscription(fn, &LambdaSubscriptionProps{
//   	FilterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter_stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   			"matchSuffixes": []*string{
//   				jsii.String("ue"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter_stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter_numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter_existsFilter(),
//   	},
//   }))
//
type SubscriptionFilter interface {
	// conditions that specify the message attributes that should be included, excluded, matched, etc.
	Conditions() *[]interface{}
}

// The jsii proxy struct for SubscriptionFilter
type jsiiProxy_SubscriptionFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_SubscriptionFilter) Conditions() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}


func NewSubscriptionFilter(conditions *[]interface{}) SubscriptionFilter {
	_init_.Initialize()

	j := jsiiProxy_SubscriptionFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		[]interface{}{conditions},
		&j,
	)

	return &j
}

func NewSubscriptionFilter_Override(s SubscriptionFilter, conditions *[]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		[]interface{}{conditions},
		s,
	)
}

// Returns a subscription filter for attribute key matching.
func SubscriptionFilter_ExistsFilter() SubscriptionFilter {
	_init_.Initialize()

	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		"existsFilter",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a subscription filter for a numeric attribute.
func SubscriptionFilter_NumericFilter(numericConditions *NumericConditions) SubscriptionFilter {
	_init_.Initialize()

	if err := validateSubscriptionFilter_NumericFilterParameters(numericConditions); err != nil {
		panic(err)
	}
	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		"numericFilter",
		[]interface{}{numericConditions},
		&returns,
	)

	return returns
}

// Returns a subscription filter for a string attribute.
func SubscriptionFilter_StringFilter(stringConditions *StringConditions) SubscriptionFilter {
	_init_.Initialize()

	if err := validateSubscriptionFilter_StringFilterParameters(stringConditions); err != nil {
		panic(err)
	}
	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		"stringFilter",
		[]interface{}{stringConditions},
		&returns,
	)

	return returns
}

