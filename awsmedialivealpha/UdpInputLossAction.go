package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// UDP input loss action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   udpInputLossAction := medialive_alpha.UdpInputLossAction_DROP_PROGRAM()
//
// Experimental.
type UdpInputLossAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for UdpInputLossAction
type jsiiProxy_UdpInputLossAction struct {
	_ byte // padding
}

func (j *jsiiProxy_UdpInputLossAction) Value() *string {
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
func UdpInputLossAction_Of(value *string) UdpInputLossAction {
	_init_.Initialize()

	if err := validateUdpInputLossAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns UdpInputLossAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.UdpInputLossAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func UdpInputLossAction_DROP_PROGRAM() UdpInputLossAction {
	_init_.Initialize()
	var returns UdpInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpInputLossAction",
		"DROP_PROGRAM",
		&returns,
	)
	return returns
}

func UdpInputLossAction_DROP_TS() UdpInputLossAction {
	_init_.Initialize()
	var returns UdpInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpInputLossAction",
		"DROP_TS",
		&returns,
	)
	return returns
}

func UdpInputLossAction_EMIT_PROGRAM() UdpInputLossAction {
	_init_.Initialize()
	var returns UdpInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpInputLossAction",
		"EMIT_PROGRAM",
		&returns,
	)
	return returns
}

