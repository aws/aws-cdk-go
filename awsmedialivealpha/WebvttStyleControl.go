package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether WebVTT passes through source style/position.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   webvttStyleControl := medialive_alpha.WebvttStyleControl_Of(jsii.String("value"))
//
// Experimental.
type WebvttStyleControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for WebvttStyleControl
type jsiiProxy_WebvttStyleControl struct {
	_ byte // padding
}

func (j *jsiiProxy_WebvttStyleControl) Value() *string {
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
func WebvttStyleControl_Of(value *string) WebvttStyleControl {
	_init_.Initialize()

	if err := validateWebvttStyleControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns WebvttStyleControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.WebvttStyleControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WebvttStyleControl_NO_STYLE_DATA() WebvttStyleControl {
	_init_.Initialize()
	var returns WebvttStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WebvttStyleControl",
		"NO_STYLE_DATA",
		&returns,
	)
	return returns
}

func WebvttStyleControl_PASSTHROUGH() WebvttStyleControl {
	_init_.Initialize()
	var returns WebvttStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WebvttStyleControl",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

