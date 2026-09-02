package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 LFE control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3LfeControl := medialive_alpha.Eac3LfeControl_Of(jsii.String("value"))
//
// Experimental.
type Eac3LfeControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3LfeControl
type jsiiProxy_Eac3LfeControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3LfeControl) Value() *string {
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
func Eac3LfeControl_Of(value *string) Eac3LfeControl {
	_init_.Initialize()

	if err := validateEac3LfeControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3LfeControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3LfeControl_LFE() Eac3LfeControl {
	_init_.Initialize()
	var returns Eac3LfeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeControl",
		"LFE",
		&returns,
	)
	return returns
}

func Eac3LfeControl_NO_LFE() Eac3LfeControl {
	_init_.Initialize()
	var returns Eac3LfeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeControl",
		"NO_LFE",
		&returns,
	)
	return returns
}

