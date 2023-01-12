package awsmediapackage


// Parameters for an HLS manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsManifestProperty := &hlsManifestProperty{
//   	adMarkers: jsii.String("adMarkers"),
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	manifestName: jsii.String("manifestName"),
//   	programDateTimeIntervalSeconds: jsii.Number(123),
//   	repeatExtXKey: jsii.Boolean(false),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnPackagingConfiguration_HlsManifestProperty struct {
	// This setting controls ad markers in the packaged content.
	//
	// `NONE` omits SCTE-35 ad markers from the output. `PASSTHROUGH` copies SCTE-35 ad markers from the source content to the output. `SCTE35_ENHANCED` generates ad markers and blackout tags in the output, based on SCTE-35 messages in the source content.
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Applies to stream sets with a single video track only.
	//
	// When enabled, the output includes an additional I-frame only stream, along with the other tracks.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Additionally, ID3Timed metadata messages are generated every 5 seconds starting when the content was ingested.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Repeat the `EXT-X-KEY` directive for every media segment.
	//
	// This might result in an increase in client requests to the DRM server.
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
	// Video bitrate limitations for outputs from this packaging configuration.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

