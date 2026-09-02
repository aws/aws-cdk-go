package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 phase control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3PhaseControl := medialive_alpha.Eac3PhaseControl_Of(jsii.String("value"))
//
// Experimental.
type Eac3PhaseControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3PhaseControl
type jsiiProxy_Eac3PhaseControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3PhaseControl) Value() *string {
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
func Eac3PhaseControl_Of(value *string) Eac3PhaseControl {
	_init_.Initialize()

	if err := validateEac3PhaseControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3PhaseControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3PhaseControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3PhaseControl_NO_SHIFT() Eac3PhaseControl {
	_init_.Initialize()
	var returns Eac3PhaseControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3PhaseControl",
		"NO_SHIFT",
		&returns,
	)
	return returns
}

func Eac3PhaseControl_SHIFT_90_DEGREES() Eac3PhaseControl {
	_init_.Initialize()
	var returns Eac3PhaseControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3PhaseControl",
		"SHIFT_90_DEGREES",
		&returns,
	)
	return returns
}

