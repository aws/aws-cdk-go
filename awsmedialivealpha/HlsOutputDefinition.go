package awsmedialivealpha


// Output definition for an HLS output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//   var h265PackagingType H265PackagingType
//   var hlsSettings HlsSettings
//
//   hlsOutputDefinition := &HlsOutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	H265PackagingType: h265PackagingType,
//   	HlsSettings: hlsSettings,
//   	NameModifier: jsii.String("nameModifier"),
//   	SegmentModifier: jsii.String("segmentModifier"),
//   }
//
// Experimental.
type HlsOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// For H.265 video, whether to package as HEV1 or HVC1.
	// Default: - service default.
	//
	// Experimental.
	H265PackagingType H265PackagingType `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// The per-output HLS settings (standard, audio-only, fMP4, or frame-capture).
	//
	// Use the
	// `HlsSettings` factory methods. Standard outputs additionally configure the M3U8 container
	// via `HlsSettings.standard({ m3u8Settings })`.
	// Default: - HlsSettings.standard() with service-default M3U8 settings
	//
	// Experimental.
	HlsSettings HlsSettings `field:"optional" json:"hlsSettings" yaml:"hlsSettings"`
	// A string concatenated to the end of the destination file name.
	// Default: - service default.
	//
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// A string concatenated to the end of segment file names.
	// Default: - no segment modifier.
	//
	// Experimental.
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

