package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Shadow color for burn-in and DVB-Sub captions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionShadowColor := medialive_alpha.CaptionShadowColor_BLACK()
//
// Experimental.
type CaptionShadowColor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionShadowColor
type jsiiProxy_CaptionShadowColor struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionShadowColor) Value() *string {
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
func CaptionShadowColor_Of(value *string) CaptionShadowColor {
	_init_.Initialize()

	if err := validateCaptionShadowColor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionShadowColor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionShadowColor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionShadowColor_BLACK() CaptionShadowColor {
	_init_.Initialize()
	var returns CaptionShadowColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionShadowColor",
		"BLACK",
		&returns,
	)
	return returns
}

func CaptionShadowColor_NONE() CaptionShadowColor {
	_init_.Initialize()
	var returns CaptionShadowColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionShadowColor",
		"NONE",
		&returns,
	)
	return returns
}

func CaptionShadowColor_WHITE() CaptionShadowColor {
	_init_.Initialize()
	var returns CaptionShadowColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionShadowColor",
		"WHITE",
		&returns,
	)
	return returns
}

