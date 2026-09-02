package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 attenuation control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3AttenuationControl := medialive_alpha.Eac3AttenuationControl_Of(jsii.String("value"))
//
// Experimental.
type Eac3AttenuationControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3AttenuationControl
type jsiiProxy_Eac3AttenuationControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3AttenuationControl) Value() *string {
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
func Eac3AttenuationControl_Of(value *string) Eac3AttenuationControl {
	_init_.Initialize()

	if err := validateEac3AttenuationControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3AttenuationControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3AttenuationControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3AttenuationControl_ATTENUATE_3_DB() Eac3AttenuationControl {
	_init_.Initialize()
	var returns Eac3AttenuationControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AttenuationControl",
		"ATTENUATE_3_DB",
		&returns,
	)
	return returns
}

func Eac3AttenuationControl_NONE() Eac3AttenuationControl {
	_init_.Initialize()
	var returns Eac3AttenuationControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AttenuationControl",
		"NONE",
		&returns,
	)
	return returns
}

