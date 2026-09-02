package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of SDI input.
//
// Example:
//   var stack Stack
//
//   sdi := medialive.NewSdiSource(stack, jsii.String("Sdi"), &SdiSourceProps{
//   	SdiSourceName: jsii.String("camera-1"),
//   	Type: medialive.SdiType_SINGLE(),
//   })
//
// Experimental.
type SdiType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SdiType
type jsiiProxy_SdiType struct {
	_ byte // padding
}

func (j *jsiiProxy_SdiType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func SdiType_Of(value *string) SdiType {
	_init_.Initialize()

	if err := validateSdiType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SdiType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SdiType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SdiType_QUAD() SdiType {
	_init_.Initialize()
	var returns SdiType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SdiType",
		"QUAD",
		&returns,
	)
	return returns
}

func SdiType_SINGLE() SdiType {
	_init_.Initialize()
	var returns SdiType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SdiType",
		"SINGLE",
		&returns,
	)
	return returns
}

