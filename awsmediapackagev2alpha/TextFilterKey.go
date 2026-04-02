package awsmediapackagev2alpha


// Text manifest filter keys for free-form string values.
//
// Use with `ManifestFilter.text()` and `ManifestFilter.textList()`.
//
// For keys with fixed accepted values, use the dedicated methods on `ManifestFilter` instead:
// - Audio codec: `ManifestFilter.audioCodec()` / `ManifestFilter.audioCodecList()`
// - Video codec: `ManifestFilter.videoCodec()` / `ManifestFilter.videoCodecList()`
// - Video dynamic range: `ManifestFilter.videoDynamicRange()` / `ManifestFilter.videoDynamicRangeList()`
// - Trickplay type: `ManifestFilter.trickplayType()` / `ManifestFilter.trickplayTypeList()`
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
type TextFilterKey string

const (
	// Audio languages or functional codes derived from encoder passthrough.
	//
	// Accepted values: arbitrary strings such as ISO-639-1 language codes (not case-sensitive).
	// Experimental.
	TextFilterKey_AUDIO_LANGUAGE TextFilterKey = "AUDIO_LANGUAGE"
	// The subtitle language or functional codes derived from encoder passthrough.
	//
	// Accepted values: arbitrary strings such as ISO-639-1 language codes (not case-sensitive).
	// Experimental.
	TextFilterKey_SUBTITLE_LANGUAGE TextFilterKey = "SUBTITLE_LANGUAGE"
)

