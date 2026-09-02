package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Background color for burn-in and DVB-Sub captions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionBackgroundColor := medialive_alpha.CaptionBackgroundColor_BLACK()
//
// Experimental.
type CaptionBackgroundColor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionBackgroundColor
type jsiiProxy_CaptionBackgroundColor struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionBackgroundColor) Value() *string {
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
func CaptionBackgroundColor_Of(value *string) CaptionBackgroundColor {
	_init_.Initialize()

	if err := validateCaptionBackgroundColor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionBackgroundColor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionBackgroundColor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionBackgroundColor_BLACK() CaptionBackgroundColor {
	_init_.Initialize()
	var returns CaptionBackgroundColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionBackgroundColor",
		"BLACK",
		&returns,
	)
	return returns
}

func CaptionBackgroundColor_NONE() CaptionBackgroundColor {
	_init_.Initialize()
	var returns CaptionBackgroundColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionBackgroundColor",
		"NONE",
		&returns,
	)
	return returns
}

func CaptionBackgroundColor_WHITE() CaptionBackgroundColor {
	_init_.Initialize()
	var returns CaptionBackgroundColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionBackgroundColor",
		"WHITE",
		&returns,
	)
	return returns
}

