package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Action to take when the current input completes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputEndAction := medialive_alpha.InputEndAction_Of(jsii.String("value"))
//
// Experimental.
type InputEndAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputEndAction
type jsiiProxy_InputEndAction struct {
	_ byte // padding
}

func (j *jsiiProxy_InputEndAction) Value() *string {
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
func InputEndAction_Of(value *string) InputEndAction {
	_init_.Initialize()

	if err := validateInputEndAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputEndAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputEndAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputEndAction_NONE() InputEndAction {
	_init_.Initialize()
	var returns InputEndAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputEndAction",
		"NONE",
		&returns,
	)
	return returns
}

func InputEndAction_SWITCH_AND_LOOP_INPUTS() InputEndAction {
	_init_.Initialize()
	var returns InputEndAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputEndAction",
		"SWITCH_AND_LOOP_INPUTS",
		&returns,
	)
	return returns
}

