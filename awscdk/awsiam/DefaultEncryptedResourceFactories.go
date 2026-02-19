package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Default factories for encrypted resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultEncryptedResourceFactories := awscdk.Aws_iam.NewDefaultEncryptedResourceFactories()
//
type DefaultEncryptedResourceFactories interface {
}

// The jsii proxy struct for DefaultEncryptedResourceFactories
type jsiiProxy_DefaultEncryptedResourceFactories struct {
	_ byte // padding
}

func NewDefaultEncryptedResourceFactories() DefaultEncryptedResourceFactories {
	_init_.Initialize()

	j := jsiiProxy_DefaultEncryptedResourceFactories{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.DefaultEncryptedResourceFactories",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDefaultEncryptedResourceFactories_Override(d DefaultEncryptedResourceFactories) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.DefaultEncryptedResourceFactories",
		nil, // no parameters
		d,
	)
}

// Get the default factory for a given CloudFormation resource type.
func DefaultEncryptedResourceFactories_Get(type_ *string) IEncryptedResourceFactory {
	_init_.Initialize()

	if err := validateDefaultEncryptedResourceFactories_GetParameters(type_); err != nil {
		panic(err)
	}
	var returns IEncryptedResourceFactory

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.DefaultEncryptedResourceFactories",
		"get",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Check if a default factory is registered for a given CloudFormation resource type.
func DefaultEncryptedResourceFactories_Has(type_ *string) *bool {
	_init_.Initialize()

	if err := validateDefaultEncryptedResourceFactories_HasParameters(type_); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.DefaultEncryptedResourceFactories",
		"has",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Register a default factory for a given CloudFormation resource type.
func DefaultEncryptedResourceFactories_Set(type_ *string, factory IEncryptedResourceFactory) {
	_init_.Initialize()

	if err := validateDefaultEncryptedResourceFactories_SetParameters(type_, factory); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_iam.DefaultEncryptedResourceFactories",
		"set",
		[]interface{}{type_, factory},
	)
}

