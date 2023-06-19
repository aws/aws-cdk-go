package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Accessor for pseudo parameters.
//
// Since pseudo parameters need to be anchored to a stack somewhere in the
// construct tree, this class takes an scope parameter; the pseudo parameter
// values can be obtained as properties from an scoped object.
// Experimental.
type Aws interface {
}

// The jsii proxy struct for Aws
type jsiiProxy_Aws struct {
	_ byte // padding
}

func Aws_ACCOUNT_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"ACCOUNT_ID",
		&returns,
	)
	return returns
}

func Aws_NO_VALUE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"NO_VALUE",
		&returns,
	)
	return returns
}

func Aws_NOTIFICATION_ARNS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"NOTIFICATION_ARNS",
		&returns,
	)
	return returns
}

func Aws_PARTITION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"PARTITION",
		&returns,
	)
	return returns
}

func Aws_REGION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"REGION",
		&returns,
	)
	return returns
}

func Aws_STACK_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"STACK_ID",
		&returns,
	)
	return returns
}

func Aws_STACK_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"STACK_NAME",
		&returns,
	)
	return returns
}

func Aws_URL_SUFFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.Aws",
		"URL_SUFFIX",
		&returns,
	)
	return returns
}

