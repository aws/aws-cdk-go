package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth timestamp offset mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothTimestampOffsetMode := medialive_alpha.MsSmoothTimestampOffsetMode_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothTimestampOffsetMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothTimestampOffsetMode
type jsiiProxy_MsSmoothTimestampOffsetMode struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothTimestampOffsetMode) Value() *string {
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
func MsSmoothTimestampOffsetMode_Of(value *string) MsSmoothTimestampOffsetMode {
	_init_.Initialize()

	if err := validateMsSmoothTimestampOffsetMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothTimestampOffsetMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothTimestampOffsetMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothTimestampOffsetMode_USE_CONFIGURED_OFFSET() MsSmoothTimestampOffsetMode {
	_init_.Initialize()
	var returns MsSmoothTimestampOffsetMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothTimestampOffsetMode",
		"USE_CONFIGURED_OFFSET",
		&returns,
	)
	return returns
}

func MsSmoothTimestampOffsetMode_USE_EVENT_START_DATE() MsSmoothTimestampOffsetMode {
	_init_.Initialize()
	var returns MsSmoothTimestampOffsetMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothTimestampOffsetMode",
		"USE_EVENT_START_DATE",
		&returns,
	)
	return returns
}

