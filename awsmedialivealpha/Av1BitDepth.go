package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 bit depth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1BitDepth := medialive_alpha.Av1BitDepth_Of(jsii.String("value"))
//
// Experimental.
type Av1BitDepth interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1BitDepth
type jsiiProxy_Av1BitDepth struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1BitDepth) Value() *string {
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
func Av1BitDepth_Of(value *string) Av1BitDepth {
	_init_.Initialize()

	if err := validateAv1BitDepth_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1BitDepth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1BitDepth",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1BitDepth_BIT_DEPTH_10() Av1BitDepth {
	_init_.Initialize()
	var returns Av1BitDepth
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1BitDepth",
		"BIT_DEPTH_10",
		&returns,
	)
	return returns
}

func Av1BitDepth_BIT_DEPTH_8() Av1BitDepth {
	_init_.Initialize()
	var returns Av1BitDepth
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1BitDepth",
		"BIT_DEPTH_8",
		&returns,
	)
	return returns
}

