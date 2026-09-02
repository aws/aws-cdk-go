package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC input type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacInputType := medialive_alpha.AacInputType_Of(jsii.String("value"))
//
// Experimental.
type AacInputType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacInputType
type jsiiProxy_AacInputType struct {
	_ byte // padding
}

func (j *jsiiProxy_AacInputType) Value() *string {
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
func AacInputType_Of(value *string) AacInputType {
	_init_.Initialize()

	if err := validateAacInputType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacInputType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacInputType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacInputType_BROADCASTER_MIXED_AD() AacInputType {
	_init_.Initialize()
	var returns AacInputType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacInputType",
		"BROADCASTER_MIXED_AD",
		&returns,
	)
	return returns
}

func AacInputType_NORMAL() AacInputType {
	_init_.Initialize()
	var returns AacInputType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacInputType",
		"NORMAL",
		&returns,
	)
	return returns
}

