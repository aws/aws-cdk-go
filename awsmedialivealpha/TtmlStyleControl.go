package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether TTML passes through source style/position.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ttmlStyleControl := medialive_alpha.TtmlStyleControl_Of(jsii.String("value"))
//
// Experimental.
type TtmlStyleControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TtmlStyleControl
type jsiiProxy_TtmlStyleControl struct {
	_ byte // padding
}

func (j *jsiiProxy_TtmlStyleControl) Value() *string {
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
func TtmlStyleControl_Of(value *string) TtmlStyleControl {
	_init_.Initialize()

	if err := validateTtmlStyleControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TtmlStyleControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TtmlStyleControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TtmlStyleControl_PASSTHROUGH() TtmlStyleControl {
	_init_.Initialize()
	var returns TtmlStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TtmlStyleControl",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

func TtmlStyleControl_USE_CONFIGURED() TtmlStyleControl {
	_init_.Initialize()
	var returns TtmlStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TtmlStyleControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

