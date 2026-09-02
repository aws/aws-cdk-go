package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacProfile := medialive_alpha.AacProfile_HEV1()
//
// Experimental.
type AacProfile interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacProfile
type jsiiProxy_AacProfile struct {
	_ byte // padding
}

func (j *jsiiProxy_AacProfile) Value() *string {
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
func AacProfile_Of(value *string) AacProfile {
	_init_.Initialize()

	if err := validateAacProfile_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacProfile

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacProfile",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacProfile_HEV1() AacProfile {
	_init_.Initialize()
	var returns AacProfile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacProfile",
		"HEV1",
		&returns,
	)
	return returns
}

func AacProfile_HEV2() AacProfile {
	_init_.Initialize()
	var returns AacProfile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacProfile",
		"HEV2",
		&returns,
	)
	return returns
}

func AacProfile_LC() AacProfile {
	_init_.Initialize()
	var returns AacProfile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacProfile",
		"LC",
		&returns,
	)
	return returns
}

