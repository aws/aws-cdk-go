package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 passthrough control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3PassthroughControl := medialive_alpha.Eac3PassthroughControl_Of(jsii.String("value"))
//
// Experimental.
type Eac3PassthroughControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3PassthroughControl
type jsiiProxy_Eac3PassthroughControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3PassthroughControl) Value() *string {
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
func Eac3PassthroughControl_Of(value *string) Eac3PassthroughControl {
	_init_.Initialize()

	if err := validateEac3PassthroughControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3PassthroughControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3PassthroughControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3PassthroughControl_NO_PASSTHROUGH() Eac3PassthroughControl {
	_init_.Initialize()
	var returns Eac3PassthroughControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3PassthroughControl",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func Eac3PassthroughControl_WHEN_POSSIBLE() Eac3PassthroughControl {
	_init_.Initialize()
	var returns Eac3PassthroughControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3PassthroughControl",
		"WHEN_POSSIBLE",
		&returns,
	)
	return returns
}

