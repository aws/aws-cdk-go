package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls whether a fixed grid is used to generate the subtitle bitmap (Teletext input).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionTeletextGridControl := medialive_alpha.CaptionTeletextGridControl_Of(jsii.String("value"))
//
// Experimental.
type CaptionTeletextGridControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionTeletextGridControl
type jsiiProxy_CaptionTeletextGridControl struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionTeletextGridControl) Value() *string {
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
func CaptionTeletextGridControl_Of(value *string) CaptionTeletextGridControl {
	_init_.Initialize()

	if err := validateCaptionTeletextGridControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionTeletextGridControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionTeletextGridControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionTeletextGridControl_FIXED() CaptionTeletextGridControl {
	_init_.Initialize()
	var returns CaptionTeletextGridControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionTeletextGridControl",
		"FIXED",
		&returns,
	)
	return returns
}

func CaptionTeletextGridControl_SCALED() CaptionTeletextGridControl {
	_init_.Initialize()
	var returns CaptionTeletextGridControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionTeletextGridControl",
		"SCALED",
		&returns,
	)
	return returns
}

