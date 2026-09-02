package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter settings for H.264 video. Supports temporal filter and bandwidth reduction filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var bandwidthReductionPostFilterSharpening BandwidthReductionPostFilterSharpening
//   var bandwidthReductionStrength BandwidthReductionStrength
//
//   h264FilterSettings := medialive_alpha.H264FilterSettings_BandwidthReductionFilter(&BandwidthReductionFilterProps{
//   	PostFilterSharpening: bandwidthReductionPostFilterSharpening,
//   	Strength: bandwidthReductionStrength,
//   })
//
// Experimental.
type H264FilterSettings interface {
}

// The jsii proxy struct for H264FilterSettings
type jsiiProxy_H264FilterSettings struct {
	_ byte // padding
}

// Apply a bandwidth reduction filter.
// Experimental.
func H264FilterSettings_BandwidthReductionFilter(props *BandwidthReductionFilterProps) H264FilterSettings {
	_init_.Initialize()

	if err := validateH264FilterSettings_BandwidthReductionFilterParameters(props); err != nil {
		panic(err)
	}
	var returns H264FilterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264FilterSettings",
		"bandwidthReductionFilter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Apply a temporal filter.
// Experimental.
func H264FilterSettings_TemporalFilter(props *TemporalFilterProps) H264FilterSettings {
	_init_.Initialize()

	if err := validateH264FilterSettings_TemporalFilterParameters(props); err != nil {
		panic(err)
	}
	var returns H264FilterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264FilterSettings",
		"temporalFilter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

