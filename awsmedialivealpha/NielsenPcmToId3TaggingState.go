package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether Nielsen PCM to ID3 tagging is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   nielsenPcmToId3TaggingState := medialive_alpha.NielsenPcmToId3TaggingState_Of(jsii.String("value"))
//
// Experimental.
type NielsenPcmToId3TaggingState interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NielsenPcmToId3TaggingState
type jsiiProxy_NielsenPcmToId3TaggingState struct {
	_ byte // padding
}

func (j *jsiiProxy_NielsenPcmToId3TaggingState) Value() *string {
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
func NielsenPcmToId3TaggingState_Of(value *string) NielsenPcmToId3TaggingState {
	_init_.Initialize()

	if err := validateNielsenPcmToId3TaggingState_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NielsenPcmToId3TaggingState

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NielsenPcmToId3TaggingState",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NielsenPcmToId3TaggingState_DISABLED() NielsenPcmToId3TaggingState {
	_init_.Initialize()
	var returns NielsenPcmToId3TaggingState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenPcmToId3TaggingState",
		"DISABLED",
		&returns,
	)
	return returns
}

func NielsenPcmToId3TaggingState_ENABLED() NielsenPcmToId3TaggingState {
	_init_.Initialize()
	var returns NielsenPcmToId3TaggingState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenPcmToId3TaggingState",
		"ENABLED",
		&returns,
	)
	return returns
}

