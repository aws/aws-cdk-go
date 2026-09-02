package awsmedialivealpha


// Output definition for a MediaPackage V2 output group.
//
// MediaPackage V2 uses CMAF ingest which requires one media track (video or audio) per output.
// In-band captions (burn-in, embedded) can ride alongside the primary encode because they do not
// produce a separate track.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//   var mediaPackageV2HlsSetting MediaPackageV2HlsSetting
//
//   mediaPackageV2OutputDefinition := &MediaPackageV2OutputDefinition{
//   	Encode: encodeConfiguration,
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	AudioGroupId: jsii.String("audioGroupId"),
//   	AudioRenditionSets: jsii.String("audioRenditionSets"),
//   	Captions: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	HlsAutoSelect: mediaPackageV2HlsSetting,
//   	HlsDefault: mediaPackageV2HlsSetting,
//   }
//
// Experimental.
type MediaPackageV2OutputDefinition struct {
	// The primary encode for this output — one video or one audio track.
	// Experimental.
	Encode EncodeConfiguration `field:"required" json:"encode" yaml:"encode"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The audio group ID for audio outputs.
	// Default: - service default.
	//
	// Experimental.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// The audio rendition sets for video outputs.
	// Default: - service default.
	//
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Caption encodes that ride alongside the primary encode.
	//
	// Only in-band caption types are
	// allowed (burn-in, embedded) — out-of-band captions must go in their own output.
	// Default: - no captions on this output.
	//
	// Experimental.
	Captions *[]EncodeConfiguration `field:"optional" json:"captions" yaml:"captions"`
	// For audio outputs, whether MediaPackage sets this rendition as the auto-select rendition in the HLS manifest.
	// Default: MediaPackageV2HlsSetting.OMIT
	//
	// Experimental.
	HlsAutoSelect MediaPackageV2HlsSetting `field:"optional" json:"hlsAutoSelect" yaml:"hlsAutoSelect"`
	// For audio outputs, whether MediaPackage sets this rendition as the default rendition in the HLS manifest.
	// Default: MediaPackageV2HlsSetting.OMIT
	//
	// Experimental.
	HlsDefault MediaPackageV2HlsSetting `field:"optional" json:"hlsDefault" yaml:"hlsDefault"`
}

