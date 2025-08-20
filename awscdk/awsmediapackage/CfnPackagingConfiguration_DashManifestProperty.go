package awsmediapackage


// Parameters for a DASH manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashManifestProperty := &DashManifestProperty{
//   	ManifestLayout: jsii.String("manifestLayout"),
//   	ManifestName: jsii.String("manifestName"),
//   	MinBufferTimeSeconds: jsii.Number(123),
//   	Profile: jsii.String("profile"),
//   	ScteMarkersSource: jsii.String("scteMarkersSource"),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html
//
type CfnPackagingConfiguration_DashManifestProperty struct {
	// Determines the position of some tags in the Media Presentation Description (MPD).
	//
	// When set to `FULL` , elements like `SegmentTemplate` and `ContentProtection` are included in each `Representation` . When set to `COMPACT` , duplicate elements are combined and presented at the AdaptationSet level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-manifestlayout
	//
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-manifestname
	//
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-minbuffertimeseconds
	//
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// The DASH profile type.
	//
	// When set to `HBBTV_1_5` , the content is compliant with HbbTV 1.5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-profile
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The source of scte markers used.
	//
	// Value description:
	//
	// - `SEGMENTS` - The scte markers are sourced from the segments of the ingested content.
	// - `MANIFEST` - the scte markers are sourced from the manifest of the ingested content. The MANIFEST value is compatible with source HLS playlists using the SCTE-35 Enhanced syntax ( `EXT-OATCLS-SCTE35` tags). SCTE-35 Elemental and SCTE-35 Daterange syntaxes are not supported with this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-sctemarkerssource
	//
	ScteMarkersSource *string `field:"optional" json:"scteMarkersSource" yaml:"scteMarkersSource"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-streamselection
	//
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

