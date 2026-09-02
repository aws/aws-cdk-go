package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls placement of EBP markers on audio PIDs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsEbpPlacement := medialive_alpha.M2tsEbpPlacement_Of(jsii.String("value"))
//
// Experimental.
type M2tsEbpPlacement interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsEbpPlacement
type jsiiProxy_M2tsEbpPlacement struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsEbpPlacement) Value() *string {
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
func M2tsEbpPlacement_Of(value *string) M2tsEbpPlacement {
	_init_.Initialize()

	if err := validateM2tsEbpPlacement_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsEbpPlacement

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpPlacement",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsEbpPlacement_VIDEO_AND_AUDIO_PIDS() M2tsEbpPlacement {
	_init_.Initialize()
	var returns M2tsEbpPlacement
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpPlacement",
		"VIDEO_AND_AUDIO_PIDS",
		&returns,
	)
	return returns
}

func M2tsEbpPlacement_VIDEO_PID() M2tsEbpPlacement {
	_init_.Initialize()
	var returns M2tsEbpPlacement
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpPlacement",
		"VIDEO_PID",
		&returns,
	)
	return returns
}

