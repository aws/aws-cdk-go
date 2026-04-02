package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Enables you to create filters for your Origin Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   manifestFilter := mediapackagev2_alpha.ManifestFilter_AudioCodec(mediapackagev2_alpha.AudioCodec_AACL)
//
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filtering.html
//
// Experimental.
type ManifestFilter interface {
	// Normalised manifest filter string.
	// Experimental.
	FilterString() *string
}

// The jsii proxy struct for ManifestFilter
type jsiiProxy_ManifestFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_ManifestFilter) FilterString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterString",
		&returns,
	)
	return returns
}


// Experimental.
func NewManifestFilter(filterString *string) ManifestFilter {
	_init_.Initialize()

	if err := validateNewManifestFilterParameters(filterString); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManifestFilter{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		[]interface{}{filterString},
		&j,
	)

	return &j
}

// Experimental.
func NewManifestFilter_Override(m ManifestFilter, filterString *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		[]interface{}{filterString},
		m,
	)
}

// Filter by a single audio codec.
// Experimental.
func ManifestFilter_AudioCodec(value AudioCodec) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_AudioCodecParameters(value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"audioCodec",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Filter by multiple audio codecs.
// Experimental.
func ManifestFilter_AudioCodecList(values *[]AudioCodec) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_AudioCodecListParameters(values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"audioCodecList",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Filter by a single bitrate value.
// Experimental.
func ManifestFilter_Bitrate(key BitrateFilterKey, value awscdk.Bitrate) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_BitrateParameters(key, value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"bitrate",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Filter by a combination of bitrate ranges and individual values.
//
// Example:
//   awsmediapackagev2alpha.ManifestFilter_BitrateCombo(awsmediapackagev2alpha.BitrateFilterKey_VIDEO_BITRATE, []BitrateExpression{
//   	awsmediapackagev2alpha.BitrateExpression_Range(awscdk.Bitrate_Mbps(jsii.Number(1)), awscdk.Bitrate_Mbps(jsii.Number(3))),
//   	awsmediapackagev2alpha.BitrateExpression_Value(awscdk.Bitrate_Mbps(jsii.Number(5))),
//   })
//
// Experimental.
func ManifestFilter_BitrateCombo(key BitrateFilterKey, expressions *[]BitrateExpression) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_BitrateComboParameters(key, expressions); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"bitrateCombo",
		[]interface{}{key, expressions},
		&returns,
	)

	return returns
}

// Filter by a bitrate range (inclusive).
// Experimental.
func ManifestFilter_BitrateRange(key BitrateFilterKey, start awscdk.Bitrate, end awscdk.Bitrate) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_BitrateRangeParameters(key, start, end); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"bitrateRange",
		[]interface{}{key, start, end},
		&returns,
	)

	return returns
}

// Specify a custom manifest filter string.
// Experimental.
func ManifestFilter_Custom(custom *string) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_CustomParameters(custom); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"custom",
		[]interface{}{custom},
		&returns,
	)

	return returns
}

// Specify a single numeric filter value.
// Experimental.
func ManifestFilter_Numeric(key NumericFilterKey, value *float64) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_NumericParameters(key, value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"numeric",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Filter by a combination of numeric ranges and individual values.
//
// Use this for advanced patterns like `video_height:240-360,720-1080,1440`.
//
// Example:
//   awsmediapackagev2alpha.ManifestFilter_NumericCombo(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, []NumericExpression{
//   	awsmediapackagev2alpha.NumericExpression_Range(jsii.Number(240), jsii.Number(360)),
//   	awsmediapackagev2alpha.NumericExpression_Range(jsii.Number(720), jsii.Number(1080)),
//   	awsmediapackagev2alpha.NumericExpression_Value(jsii.Number(1440)),
//   })
//
// Experimental.
func ManifestFilter_NumericCombo(key NumericFilterKey, expressions *[]NumericExpression) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_NumericComboParameters(key, expressions); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"numericCombo",
		[]interface{}{key, expressions},
		&returns,
	)

	return returns
}

// Specify multiple numeric filter values.
// Experimental.
func ManifestFilter_NumericList(key NumericFilterKey, values *[]*float64) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_NumericListParameters(key, values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"numericList",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Specify a numeric filter range (inclusive).
// Experimental.
func ManifestFilter_NumericRange(key NumericFilterKey, start *float64, end *float64) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_NumericRangeParameters(key, start, end); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"numericRange",
		[]interface{}{key, start, end},
		&returns,
	)

	return returns
}

// Specify a free-form text filter value (for language keys).
// Experimental.
func ManifestFilter_Text(key TextFilterKey, value *string) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_TextParameters(key, value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"text",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Specify multiple free-form text filter values (for language keys).
// Experimental.
func ManifestFilter_TextList(key TextFilterKey, values *[]*string) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_TextListParameters(key, values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"textList",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Filter by a single trickplay type.
// Experimental.
func ManifestFilter_TrickplayType(value TrickplayType) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_TrickplayTypeParameters(value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"trickplayType",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Filter by multiple trickplay types.
// Experimental.
func ManifestFilter_TrickplayTypeList(values *[]TrickplayType) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_TrickplayTypeListParameters(values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"trickplayTypeList",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Filter by a single video codec.
// Experimental.
func ManifestFilter_VideoCodec(value VideoCodec) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_VideoCodecParameters(value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"videoCodec",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Filter by multiple video codecs.
// Experimental.
func ManifestFilter_VideoCodecList(values *[]VideoCodec) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_VideoCodecListParameters(values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"videoCodecList",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Filter by a single video dynamic range.
// Experimental.
func ManifestFilter_VideoDynamicRange(value VideoDynamicRange) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_VideoDynamicRangeParameters(value); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"videoDynamicRange",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Filter by multiple video dynamic ranges.
// Experimental.
func ManifestFilter_VideoDynamicRangeList(values *[]VideoDynamicRange) ManifestFilter {
	_init_.Initialize()

	if err := validateManifestFilter_VideoDynamicRangeListParameters(values); err != nil {
		panic(err)
	}
	var returns ManifestFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		"videoDynamicRangeList",
		[]interface{}{values},
		&returns,
	)

	return returns
}

