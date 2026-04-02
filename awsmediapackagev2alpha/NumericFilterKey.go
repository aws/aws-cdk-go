package awsmediapackagev2alpha


// Numeric manifest filter keys.
//
// Use with `ManifestFilter.numeric()`, `ManifestFilter.numericList()`, and `ManifestFilter.numericRange()`.
//
// Audio and video bitrate filters are not included here because they use the
// `Bitrate` class directly via dedicated methods on `ManifestFilter`.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			FilterConfiguration: &FilterConfiguration{
//   				ManifestFilter: []ManifestFilter{
//   					awsmediapackagev2alpha.ManifestFilter_NumericCombo(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, []NumericExpression{
//   						awsmediapackagev2alpha.NumericExpression_Range(jsii.Number(240), jsii.Number(360)),
//   						awsmediapackagev2alpha.NumericExpression_*Range(jsii.Number(720), jsii.Number(1080)),
//   						awsmediapackagev2alpha.NumericExpression_Value(jsii.Number(1440)),
//   					}),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type NumericFilterKey string

const (
	// The number of audio channels.
	//
	// Accepted values: range 1–32767, or individual integers.
	// Experimental.
	NumericFilterKey_AUDIO_CHANNELS NumericFilterKey = "AUDIO_CHANNELS"
	// The audio sample rate in Hz.
	//
	// Accepted values: range 0–2147483647, or individual integers.
	// Experimental.
	NumericFilterKey_AUDIO_SAMPLE_RATE NumericFilterKey = "AUDIO_SAMPLE_RATE"
	// The height of the trick-play image in pixels (I-frame and image-based trick-play).
	//
	// If using with I-frame only trick-play, `TRICKPLAY_HEIGHT` and `VIDEO_HEIGHT`
	// should have similar values. If they differ, I-frame only tracks might be
	// removed from the manifest.
	//
	// Accepted values: range 1–2147483647, or individual integers.
	// Experimental.
	NumericFilterKey_TRICKPLAY_HEIGHT NumericFilterKey = "TRICKPLAY_HEIGHT"
	// The video frame rate range in NTSC format.
	//
	// When filtering for a single value, MediaPackage uses an approximate equals
	// comparison with an epsilon tolerance of 0.0005.
	//
	// Accepted values: range 1–999.999 (up to three decimal places), or individual values.
	// Experimental.
	NumericFilterKey_VIDEO_FRAMERATE NumericFilterKey = "VIDEO_FRAMERATE"
	// The height of the video in pixels.
	//
	// If using with I-frame only trick-play, `TRICKPLAY_HEIGHT` and `VIDEO_HEIGHT`
	// should have similar values. If they differ, I-frame only tracks might be
	// removed from the manifest.
	//
	// Accepted values: range 1–32767, or individual integers.
	// Experimental.
	NumericFilterKey_VIDEO_HEIGHT NumericFilterKey = "VIDEO_HEIGHT"
)

