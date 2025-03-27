package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Policy Implementation of FilterOrPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterOrPolicy filterOrPolicy
//
//   policy := awscdk.Aws_sns.NewPolicy(map[string]filterOrPolicy{
//   	"policyDocKey": filterOrPolicy,
//   })
//
type Policy interface {
	FilterOrPolicy
	// policy argument to construct.
	PolicyDoc() *map[string]FilterOrPolicy
	// Type used in DFS buildFilterPolicyWithMessageBody to determine json value type.
	Type() FilterOrPolicyType
	// Check if instance is `Filter` type.
	IsFilter() *bool
	// Check if instance is `Policy` type.
	IsPolicy() *bool
}

// The jsii proxy struct for Policy
type jsiiProxy_Policy struct {
	jsiiProxy_FilterOrPolicy
}

func (j *jsiiProxy_Policy) PolicyDoc() *map[string]FilterOrPolicy {
	var returns *map[string]FilterOrPolicy
	_jsii_.Get(
		j,
		"policyDoc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) Type() FilterOrPolicyType {
	var returns FilterOrPolicyType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Policy constructor.
func NewPolicy(policyDoc *map[string]FilterOrPolicy) Policy {
	_init_.Initialize()

	if err := validateNewPolicyParameters(policyDoc); err != nil {
		panic(err)
	}
	j := jsiiProxy_Policy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Policy",
		[]interface{}{policyDoc},
		&j,
	)

	return &j
}

// Policy constructor.
func NewPolicy_Override(p Policy, policyDoc *map[string]FilterOrPolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Policy",
		[]interface{}{policyDoc},
		p,
	)
}

// Filter of MessageBody.
func Policy_Filter(filter SubscriptionFilter) Filter {
	_init_.Initialize()

	if err := validatePolicy_FilterParameters(filter); err != nil {
		panic(err)
	}
	var returns Filter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Policy",
		"filter",
		[]interface{}{filter},
		&returns,
	)

	return returns
}

// Policy of MessageBody.
func Policy_Policy(policy *map[string]FilterOrPolicy) Policy {
	_init_.Initialize()

	if err := validatePolicy_PolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns Policy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Policy",
		"policy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) IsFilter() *bool {
	var returns *bool

	_jsii_.Invoke(
		p,
		"isFilter",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) IsPolicy() *bool {
	var returns *bool

	_jsii_.Invoke(
		p,
		"isPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

