package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How the ARIB Captions PID is selected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsAribCaptionsPidControl := medialive_alpha.M2tsAribCaptionsPidControl_Of(jsii.String("value"))
//
// Experimental.
type M2tsAribCaptionsPidControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsAribCaptionsPidControl
type jsiiProxy_M2tsAribCaptionsPidControl struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsAribCaptionsPidControl) Value() *string {
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
func M2tsAribCaptionsPidControl_Of(value *string) M2tsAribCaptionsPidControl {
	_init_.Initialize()

	if err := validateM2tsAribCaptionsPidControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsAribCaptionsPidControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsAribCaptionsPidControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsAribCaptionsPidControl_AUTO() M2tsAribCaptionsPidControl {
	_init_.Initialize()
	var returns M2tsAribCaptionsPidControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAribCaptionsPidControl",
		"AUTO",
		&returns,
	)
	return returns
}

func M2tsAribCaptionsPidControl_USE_CONFIGURED() M2tsAribCaptionsPidControl {
	_init_.Initialize()
	var returns M2tsAribCaptionsPidControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAribCaptionsPidControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

