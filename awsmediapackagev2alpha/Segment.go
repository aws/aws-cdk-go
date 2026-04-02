package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper class for creating segment configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   segment := mediapackagev2_alpha.NewSegment()
//
// Experimental.
type Segment interface {
}

// The jsii proxy struct for Segment
type jsiiProxy_Segment struct {
	_ byte // padding
}

// Experimental.
func NewSegment() Segment {
	_init_.Initialize()

	j := jsiiProxy_Segment{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSegment_Override(s Segment) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		nil, // no parameters
		s,
	)
}

// Create a CMAF segment configuration.
//
// Use this for endpoints with ContainerType.CMAF.
// Experimental.
func Segment_Cmaf(props *CmafSegmentProps) *SegmentConfiguration {
	_init_.Initialize()

	if err := validateSegment_CmafParameters(props); err != nil {
		panic(err)
	}
	var returns *SegmentConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		"cmaf",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an ISM (Microsoft Smooth Streaming) segment configuration.
//
// Use this for endpoints with ContainerType.ISM.
// Experimental.
func Segment_Ism(props *IsmSegmentProps) *SegmentConfiguration {
	_init_.Initialize()

	if err := validateSegment_IsmParameters(props); err != nil {
		panic(err)
	}
	var returns *SegmentConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		"ism",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a TS (Transport Stream) segment configuration.
//
// Use this for endpoints with ContainerType.TS.
// Experimental.
func Segment_Ts(props *TsSegmentProps) *SegmentConfiguration {
	_init_.Initialize()

	if err := validateSegment_TsParameters(props); err != nil {
		panic(err)
	}
	var returns *SegmentConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		"ts",
		[]interface{}{props},
		&returns,
	)

	return returns
}

