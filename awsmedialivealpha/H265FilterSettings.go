package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter settings for H.265 video. Supports temporal filter and bandwidth reduction filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var bandwidthReductionPostFilterSharpening BandwidthReductionPostFilterSharpening
//   var bandwidthReductionStrength BandwidthReductionStrength
//
//   h265FilterSettings := medialive_alpha.H265FilterSettings_BandwidthReductionFilter(&BandwidthReductionFilterProps{
//   	PostFilterSharpening: bandwidthReductionPostFilterSharpening,
//   	Strength: bandwidthReductionStrength,
//   })
//
// Experimental.
type H265FilterSettings interface {
}

// The jsii proxy struct for H265FilterSettings
type jsiiProxy_H265FilterSettings struct {
	_ byte // padding
}

// Apply a bandwidth reduction filter.
// Experimental.
func H265FilterSettings_BandwidthReductionFilter(props *BandwidthReductionFilterProps) H265FilterSettings {
	_init_.Initialize()

	if err := validateH265FilterSettings_BandwidthReductionFilterParameters(props); err != nil {
		panic(err)
	}
	var returns H265FilterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265FilterSettings",
		"bandwidthReductionFilter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Apply a temporal filter.
// Experimental.
func H265FilterSettings_TemporalFilter(props *TemporalFilterProps) H265FilterSettings {
	_init_.Initialize()

	if err := validateH265FilterSettings_TemporalFilterParameters(props); err != nil {
		panic(err)
	}
	var returns H265FilterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265FilterSettings",
		"temporalFilter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

