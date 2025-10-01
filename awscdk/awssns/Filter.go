package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter implementation of FilterOrPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subscriptionFilter subscriptionFilter
//
//   filter := awscdk.Aws_sns.NewFilter(subscriptionFilter)
//
type Filter interface {
	FilterOrPolicy
	// filter argument to construct.
	FilterDoc() SubscriptionFilter
	// Type used in DFS buildFilterPolicyWithMessageBody to determine json value type.
	Type() FilterOrPolicyType
	// Check if instance is `Filter` type.
	IsFilter() *bool
	// Check if instance is `Policy` type.
	IsPolicy() *bool
}

// The jsii proxy struct for Filter
type jsiiProxy_Filter struct {
	jsiiProxy_FilterOrPolicy
}

func (j *jsiiProxy_Filter) FilterDoc() SubscriptionFilter {
	var returns SubscriptionFilter
	_jsii_.Get(
		j,
		"filterDoc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Filter) Type() FilterOrPolicyType {
	var returns FilterOrPolicyType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Policy constructor.
func NewFilter(filterDoc SubscriptionFilter) Filter {
	_init_.Initialize()

	if err := validateNewFilterParameters(filterDoc); err != nil {
		panic(err)
	}
	j := jsiiProxy_Filter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Filter",
		[]interface{}{filterDoc},
		&j,
	)

	return &j
}

// Policy constructor.
func NewFilter_Override(f Filter, filterDoc SubscriptionFilter) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Filter",
		[]interface{}{filterDoc},
		f,
	)
}

// Filter of MessageBody.
func Filter_Filter(filter SubscriptionFilter) Filter {
	_init_.Initialize()

	if err := validateFilter_FilterParameters(filter); err != nil {
		panic(err)
	}
	var returns Filter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Filter",
		"filter",
		[]interface{}{filter},
		&returns,
	)

	return returns
}

// Policy of MessageBody.
func Filter_Policy(policy *map[string]FilterOrPolicy) Policy {
	_init_.Initialize()

	if err := validateFilter_PolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns Policy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Filter",
		"policy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Filter) IsFilter() *bool {
	var returns *bool

	_jsii_.Invoke(
		f,
		"isFilter",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Filter) IsPolicy() *bool {
	var returns *bool

	_jsii_.Invoke(
		f,
		"isPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

