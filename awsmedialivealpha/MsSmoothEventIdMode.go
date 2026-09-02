package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth event ID mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothEventIdMode := medialive_alpha.MsSmoothEventIdMode_NO_EVENT_ID()
//
// Experimental.
type MsSmoothEventIdMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothEventIdMode
type jsiiProxy_MsSmoothEventIdMode struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothEventIdMode) Value() *string {
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
func MsSmoothEventIdMode_Of(value *string) MsSmoothEventIdMode {
	_init_.Initialize()

	if err := validateMsSmoothEventIdMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothEventIdMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventIdMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothEventIdMode_NO_EVENT_ID() MsSmoothEventIdMode {
	_init_.Initialize()
	var returns MsSmoothEventIdMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventIdMode",
		"NO_EVENT_ID",
		&returns,
	)
	return returns
}

func MsSmoothEventIdMode_USE_CONFIGURED() MsSmoothEventIdMode {
	_init_.Initialize()
	var returns MsSmoothEventIdMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventIdMode",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

func MsSmoothEventIdMode_USE_TIMESTAMP() MsSmoothEventIdMode {
	_init_.Initialize()
	var returns MsSmoothEventIdMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothEventIdMode",
		"USE_TIMESTAMP",
		&returns,
	)
	return returns
}

