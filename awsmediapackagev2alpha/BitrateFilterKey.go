package awsmediapackagev2alpha


// Bitrate manifest filter keys.
//
// Use with `ManifestFilter.bitrate()` and `ManifestFilter.bitrateRange()`.
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
//   					awsmediapackagev2alpha.ManifestFilter_BitrateRange(awsmediapackagev2alpha.BitrateFilterKey_VIDEO_BITRATE, awscdk.Bitrate_Mbps(jsii.Number(1)), awscdk.Bitrate_*Mbps(jsii.Number(5))),
//   					awsmediapackagev2alpha.ManifestFilter_NumericRange(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, jsii.Number(720), jsii.Number(1080)),
//   					awsmediapackagev2alpha.ManifestFilter_VideoCodecList([]VideoCodec{
//   						awsmediapackagev2alpha.VideoCodec_H264,
//   						awsmediapackagev2alpha.VideoCodec_H265,
//   					}),
//   					awsmediapackagev2alpha.ManifestFilter_Numeric(awsmediapackagev2alpha.NumericFilterKey_AUDIO_CHANNELS, jsii.Number(2)),
//   					awsmediapackagev2alpha.ManifestFilter_TextList(awsmediapackagev2alpha.TextFilterKey_AUDIO_LANGUAGE, []*string{
//   						jsii.String("en-US"),
//   						jsii.String("fr"),
//   					}),
//   				},
//   				TimeDelay: awscdk.Duration_Seconds(jsii.Number(30)),
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type BitrateFilterKey string

const (
	// The audio bitrate in bits per second.
	//
	// Accepted values: range 0–2147483647, or individual integers.
	// Experimental.
	BitrateFilterKey_AUDIO_BITRATE BitrateFilterKey = "AUDIO_BITRATE"
	// The video bitrate in bits per second.
	//
	// We recommend using only this filter parameter to set the video bitrate.
	// Do not also set the minimum and maximum video bitrate via the MediaPackage
	// console or AWS CLI, as your output might be skewed.
	//
	// This parameter cannot be used with trick-play streams.
	//
	// Accepted values: range 0–2147483647, or individual integers.
	// Experimental.
	BitrateFilterKey_VIDEO_BITRATE BitrateFilterKey = "VIDEO_BITRATE"
)

