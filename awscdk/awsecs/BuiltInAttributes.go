package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The built-in container instance attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   builtInAttributes := awscdk.Aws_ecs.NewBuiltInAttributes()
//
type BuiltInAttributes interface {
}

// The jsii proxy struct for BuiltInAttributes
type jsiiProxy_BuiltInAttributes struct {
	_ byte // padding
}

func NewBuiltInAttributes() BuiltInAttributes {
	_init_.Initialize()

	j := jsiiProxy_BuiltInAttributes{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewBuiltInAttributes_Override(b BuiltInAttributes) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		nil, // no parameters
		b,
	)
}

func BuiltInAttributes_AMI_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"AMI_ID",
		&returns,
	)
	return returns
}

func BuiltInAttributes_AVAILABILITY_ZONE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"AVAILABILITY_ZONE",
		&returns,
	)
	return returns
}

func BuiltInAttributes_INSTANCE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"INSTANCE_ID",
		&returns,
	)
	return returns
}

func BuiltInAttributes_INSTANCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"INSTANCE_TYPE",
		&returns,
	)
	return returns
}

func BuiltInAttributes_OS_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"OS_TYPE",
		&returns,
	)
	return returns
}

