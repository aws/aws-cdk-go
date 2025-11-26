package previewawsmediapackagemixins


// Parameters for an HLS manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsManifestProperty := &HlsManifestProperty{
//   	AdMarkers: jsii.String("adMarkers"),
//   	IncludeIframeOnlyStream: jsii.Boolean(false),
//   	ManifestName: jsii.String("manifestName"),
//   	ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   	RepeatExtXKey: jsii.Boolean(false),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html
//
type CfnPackagingConfigurationPropsMixin_HlsManifestProperty struct {
	// This setting controls ad markers in the packaged content.
	//
	// Valid values:
	//
	// - `NONE` - Omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - Creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - Generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-admarkers
	//
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Applies to stream sets with a single video track only.
	//
	// When enabled, the output includes an additional I-frame only stream, along with the other tracks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-includeiframeonlystream
	//
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-manifestname
	//
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-programdatetimeintervalseconds
	//
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Repeat the `EXT-X-KEY` directive for every media segment.
	//
	// This might result in an increase in client requests to the DRM server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-repeatextxkey
	//
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
	// Video bitrate limitations for outputs from this packaging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsmanifest.html#cfn-mediapackage-packagingconfiguration-hlsmanifest-streamselection
	//
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

