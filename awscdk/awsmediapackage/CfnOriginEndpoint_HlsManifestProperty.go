package awsmediapackage


// An HTTP Live Streaming (HLS) manifest configuration on a CMAF endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsManifestProperty := &HlsManifestProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AdMarkers: jsii.String("adMarkers"),
//   	AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	AdTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	IncludeIframeOnlyStream: jsii.Boolean(false),
//   	ManifestName: jsii.String("manifestName"),
//   	PlaylistType: jsii.String("playlistType"),
//   	PlaylistWindowSeconds: jsii.Number(123),
//   	ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html
//
type CfnOriginEndpoint_HlsManifestProperty struct {
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Controls how ad markers are included in the packaged endpoint.
	//
	// Valid values:
	//
	// - `NONE` - Omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - Creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - Generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-admarkers
	//
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The flags on SCTE-35 segmentation descriptors that have to be present for AWS Elemental MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in AWS Elemental MediaPackage , see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-adsondeliveryrestrictions
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-adtriggers
	//
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Applies to stream sets with a single video track only.
	//
	// When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-includeiframeonlystream
	//
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	//
	// The manifestName on the HLSManifest object overrides the manifestName that you provided on the originEndpoint object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-manifestname
	//
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// When specified as either `event` or `vod` , a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
	//
	// Indicates if the playlist is live-to-VOD content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-playlisttype
	//
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-playlistwindowseconds
	//
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-programdatetimeintervalseconds
	//
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// The URL that's used to request this manifest from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsmanifest.html#cfn-mediapackage-originendpoint-hlsmanifest-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

