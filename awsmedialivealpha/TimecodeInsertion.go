package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Timecode insertion mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   timecodeInsertion := medialive_alpha.TimecodeInsertion_Of(jsii.String("value"))
//
// Experimental.
type TimecodeInsertion interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimecodeInsertion
type jsiiProxy_TimecodeInsertion struct {
	_ byte // padding
}

func (j *jsiiProxy_TimecodeInsertion) Value() *string {
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
func TimecodeInsertion_Of(value *string) TimecodeInsertion {
	_init_.Initialize()

	if err := validateTimecodeInsertion_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimecodeInsertion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimecodeInsertion",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimecodeInsertion_DISABLED() TimecodeInsertion {
	_init_.Initialize()
	var returns TimecodeInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeInsertion",
		"DISABLED",
		&returns,
	)
	return returns
}

func TimecodeInsertion_PIC_TIMING_SEI() TimecodeInsertion {
	_init_.Initialize()
	var returns TimecodeInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeInsertion",
		"PIC_TIMING_SEI",
		&returns,
	)
	return returns
}

