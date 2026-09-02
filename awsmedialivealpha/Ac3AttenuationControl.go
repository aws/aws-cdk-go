package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AC3 attenuation control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ac3AttenuationControl := medialive_alpha.Ac3AttenuationControl_Of(jsii.String("value"))
//
// Experimental.
type Ac3AttenuationControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Ac3AttenuationControl
type jsiiProxy_Ac3AttenuationControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Ac3AttenuationControl) Value() *string {
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
func Ac3AttenuationControl_Of(value *string) Ac3AttenuationControl {
	_init_.Initialize()

	if err := validateAc3AttenuationControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Ac3AttenuationControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Ac3AttenuationControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Ac3AttenuationControl_ATTENUATE_3_DB() Ac3AttenuationControl {
	_init_.Initialize()
	var returns Ac3AttenuationControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3AttenuationControl",
		"ATTENUATE_3_DB",
		&returns,
	)
	return returns
}

func Ac3AttenuationControl_NONE() Ac3AttenuationControl {
	_init_.Initialize()
	var returns Ac3AttenuationControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3AttenuationControl",
		"NONE",
		&returns,
	)
	return returns
}

