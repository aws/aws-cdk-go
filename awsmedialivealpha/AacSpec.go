package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacSpec := medialive_alpha.AacSpec_Of(jsii.String("value"))
//
// Experimental.
type AacSpec interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacSpec
type jsiiProxy_AacSpec struct {
	_ byte // padding
}

func (j *jsiiProxy_AacSpec) Value() *string {
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
func AacSpec_Of(value *string) AacSpec {
	_init_.Initialize()

	if err := validateAacSpec_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacSpec

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacSpec",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacSpec_MPEG2() AacSpec {
	_init_.Initialize()
	var returns AacSpec
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacSpec",
		"MPEG2",
		&returns,
	)
	return returns
}

func AacSpec_MPEG4() AacSpec {
	_init_.Initialize()
	var returns AacSpec
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacSpec",
		"MPEG4",
		&returns,
	)
	return returns
}

