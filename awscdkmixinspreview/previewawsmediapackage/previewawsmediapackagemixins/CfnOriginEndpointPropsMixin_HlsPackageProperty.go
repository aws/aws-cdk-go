package previewawsmediapackagemixins


// Parameters for Apple HLS packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsPackageProperty := &HlsPackageProperty{
//   	AdMarkers: jsii.String("adMarkers"),
//   	AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	AdTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	Encryption: &HlsEncryptionProperty{
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		EncryptionMethod: jsii.String("encryptionMethod"),
//   		KeyRotationIntervalSeconds: jsii.Number(123),
//   		RepeatExtXKey: jsii.Boolean(false),
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			SystemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	IncludeDvbSubtitles: jsii.Boolean(false),
//   	IncludeIframeOnlyStream: jsii.Boolean(false),
//   	PlaylistType: jsii.String("playlistType"),
//   	PlaylistWindowSeconds: jsii.Number(123),
//   	ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   	SegmentDurationSeconds: jsii.Number(123),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   	UseAudioRenditionGroup: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html
//
type CfnOriginEndpointPropsMixin_HlsPackageProperty struct {
	// Controls how ad markers are included in the packaged endpoint.
	//
	// Valid values:
	//
	// - `NONE` - Omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - Creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - Generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-admarkers
	//
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The flags on SCTE-35 segmentation descriptors that have to be present for AWS Elemental MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in AWS Elemental MediaPackage , see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-adsondeliveryrestrictions
	//
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that AWS Elemental MediaPackage treats as ad markers in the output manifest.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-adtriggers
	//
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Parameters for encrypting content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-includedvbsubtitles
	//
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// Only applies to stream sets with a single video track.
	//
	// When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-includeiframeonlystream
	//
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// When specified as either `event` or `vod` , a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
	//
	// Indicates if the playlist is live-to-VOD content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-playlisttype
	//
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-playlistwindowseconds
	//
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-programdatetimeintervalseconds
	//
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-streamselection
	//
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-useaudiorenditiongroup
	//
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

