package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC rate control mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacRateControlMode := medialive_alpha.AacRateControlMode_Of(jsii.String("value"))
//
// Experimental.
type AacRateControlMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacRateControlMode
type jsiiProxy_AacRateControlMode struct {
	_ byte // padding
}

func (j *jsiiProxy_AacRateControlMode) Value() *string {
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
func AacRateControlMode_Of(value *string) AacRateControlMode {
	_init_.Initialize()

	if err := validateAacRateControlMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacRateControlMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacRateControlMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacRateControlMode_CBR() AacRateControlMode {
	_init_.Initialize()
	var returns AacRateControlMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacRateControlMode",
		"CBR",
		&returns,
	)
	return returns
}

func AacRateControlMode_VBR() AacRateControlMode {
	_init_.Initialize()
	var returns AacRateControlMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacRateControlMode",
		"VBR",
		&returns,
	)
	return returns
}

