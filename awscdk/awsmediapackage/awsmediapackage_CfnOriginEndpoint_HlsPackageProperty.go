package awsmediapackage


// Parameters for Apple HLS packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPackageProperty := &hlsPackageProperty{
//   	adMarkers: jsii.String("adMarkers"),
//   	adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	adTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	encryption: &hlsEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		encryptionMethod: jsii.String("encryptionMethod"),
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   		repeatExtXKey: jsii.Boolean(false),
//   	},
//   	includeDvbSubtitles: jsii.Boolean(false),
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	playlistType: jsii.String("playlistType"),
//   	playlistWindowSeconds: jsii.Number(123),
//   	programDateTimeIntervalSeconds: jsii.Number(123),
//   	segmentDurationSeconds: jsii.Number(123),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   	useAudioRenditionGroup: jsii.Boolean(false),
//   }
//
type CfnOriginEndpoint_HlsPackageProperty struct {
	// Controls how ad markers are included in the packaged endpoint.
	//
	// Valid values are `none` , `passthrough` , or `scte35_enhanced` .
	//
	// - `NONE` - omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The flags on SCTE-35 segmentation descriptors that have to be present for MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that MediaPackage treats as ad markers in the output manifest.
	//
	// Valid values:
	//
	// - `BREAK`
	// - `DISTRIBUTOR_ADVERTISEMENT`
	// - `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `DISTRIBUTOR_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_ADVERTISEMENT`
	// - `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_PLACEMENT_OPPORTUNITY`
	// - `SPLICE_INSERT`.
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// `CfnOriginEndpoint.HlsPackageProperty.IncludeDvbSubtitles`.
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// Only applies to stream sets with a single video track.
	//
	// When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// When specified as either `event` or `vod` , a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
	//
	// Indicates if the playlist is live-to-VOD content.
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Additionally, ID3Timed metadata messages are generated every 5 seconds starting when the content was ingested.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

