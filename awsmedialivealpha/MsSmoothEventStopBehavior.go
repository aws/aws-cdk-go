package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth event stop behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothEventStopBehavior := medialive_alpha.MsSmoothEventStopBehavior_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothEventStopBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothEventStopBehavior
type jsiiProxy_MsSmoothEventStopBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothEventStopBehavior) Value() *string {
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
func MsSmoothEventStopBehavior_Of(value *string) MsSmoothEventStopBehavior {
	_init_.Initialize()

	if err := validateMsSmoothEventStopBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothEventStopBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventStopBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothEventStopBehavior_NONE() MsSmoothEventStopBehavior {
	_init_.Initialize()
	var returns MsSmoothEventStopBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventStopBehavior",
		"NONE",
		&returns,
	)
	return returns
}

func MsSmoothEventStopBehavior_SEND_EOS() MsSmoothEventStopBehavior {
	_init_.Initialize()
	var returns MsSmoothEventStopBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventStopBehavior",
		"SEND_EOS",
		&returns,
	)
	return returns
}

