package awsmedialivealpha


// Properties for fMP4 HLS settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var m3u8NielsenId3Behavior M3u8NielsenId3Behavior
//   var m3u8TimedMetadataBehavior M3u8TimedMetadataBehavior
//
//   fmp4HlsSettingsProps := &Fmp4HlsSettingsProps{
//   	AudioRenditionSets: jsii.String("audioRenditionSets"),
//   	NielsenId3Behavior: m3u8NielsenId3Behavior,
//   	TimedMetadataBehavior: m3u8TimedMetadataBehavior,
//   }
//
// Experimental.
type Fmp4HlsSettingsProps struct {
	// The audio GROUP-IDs used with this video output stream, comma-separated.
	// Default: - service default.
	//
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Nielsen ID3 passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	NielsenId3Behavior M3u8NielsenId3Behavior `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Timed-metadata passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	TimedMetadataBehavior M3u8TimedMetadataBehavior `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
}

