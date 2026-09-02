package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS program date time clock.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsProgramDateTimeClock := medialive_alpha.HlsProgramDateTimeClock_Of(jsii.String("value"))
//
// Experimental.
type HlsProgramDateTimeClock interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsProgramDateTimeClock
type jsiiProxy_HlsProgramDateTimeClock struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsProgramDateTimeClock) Value() *string {
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
func HlsProgramDateTimeClock_Of(value *string) HlsProgramDateTimeClock {
	_init_.Initialize()

	if err := validateHlsProgramDateTimeClock_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsProgramDateTimeClock

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTimeClock",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsProgramDateTimeClock_INITIALIZE_FROM_OUTPUT_TIMECODE() HlsProgramDateTimeClock {
	_init_.Initialize()
	var returns HlsProgramDateTimeClock
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTimeClock",
		"INITIALIZE_FROM_OUTPUT_TIMECODE",
		&returns,
	)
	return returns
}

func HlsProgramDateTimeClock_SYSTEM_CLOCK() HlsProgramDateTimeClock {
	_init_.Initialize()
	var returns HlsProgramDateTimeClock
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTimeClock",
		"SYSTEM_CLOCK",
		&returns,
	)
	return returns
}

