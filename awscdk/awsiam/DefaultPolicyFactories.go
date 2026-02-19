package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Default factories for resources with policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultPolicyFactories := awscdk.Aws_iam.NewDefaultPolicyFactories()
//
type DefaultPolicyFactories interface {
}

// The jsii proxy struct for DefaultPolicyFactories
type jsiiProxy_DefaultPolicyFactories struct {
	_ byte // padding
}

func NewDefaultPolicyFactories() DefaultPolicyFactories {
	_init_.Initialize()

	j := jsiiProxy_DefaultPolicyFactories{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.DefaultPolicyFactories",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDefaultPolicyFactories_Override(d DefaultPolicyFactories) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.DefaultPolicyFactories",
		nil, // no parameters
		d,
	)
}

// Get the default factory for a given CloudFormation resource type.
func DefaultPolicyFactories_Get(type_ *string) IResourcePolicyFactory {
	_init_.Initialize()

	if err := validateDefaultPolicyFactories_GetParameters(type_); err != nil {
		panic(err)
	}
	var returns IResourcePolicyFactory

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.DefaultPolicyFactories",
		"get",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Check if a default factory is registered for a given CloudFormation resource type.
func DefaultPolicyFactories_Has(type_ *string) *bool {
	_init_.Initialize()

	if err := validateDefaultPolicyFactories_HasParameters(type_); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.DefaultPolicyFactories",
		"has",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Register a default factory for a given CloudFormation resource type.
func DefaultPolicyFactories_Set(type_ *string, factory IResourcePolicyFactory) {
	_init_.Initialize()

	if err := validateDefaultPolicyFactories_SetParameters(type_, factory); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_iam.DefaultPolicyFactories",
		"set",
		[]interface{}{type_, factory},
	)
}

