package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS ID3 segment tagging state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsId3SegmentTaggingState := medialive_alpha.HlsId3SegmentTaggingState_Of(jsii.String("value"))
//
// Experimental.
type HlsId3SegmentTaggingState interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsId3SegmentTaggingState
type jsiiProxy_HlsId3SegmentTaggingState struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsId3SegmentTaggingState) Value() *string {
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
func HlsId3SegmentTaggingState_Of(value *string) HlsId3SegmentTaggingState {
	_init_.Initialize()

	if err := validateHlsId3SegmentTaggingState_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsId3SegmentTaggingState

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsId3SegmentTaggingState",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsId3SegmentTaggingState_DISABLED() HlsId3SegmentTaggingState {
	_init_.Initialize()
	var returns HlsId3SegmentTaggingState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsId3SegmentTaggingState",
		"DISABLED",
		&returns,
	)
	return returns
}

func HlsId3SegmentTaggingState_ENABLED() HlsId3SegmentTaggingState {
	_init_.Initialize()
	var returns HlsId3SegmentTaggingState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsId3SegmentTaggingState",
		"ENABLED",
		&returns,
	)
	return returns
}

