package awsmedialivealpha


// Output definition for a CMAF Ingest output group.
//
// CMAF Ingest requires one media track (video or audio) per output. In-band captions (burn-in,
// embedded) can ride alongside the primary encode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//
//   cmafIngestOutputDefinition := &CmafIngestOutputDefinition{
//   	Encode: encodeConfiguration,
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	Captions: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// Experimental.
type CmafIngestOutputDefinition struct {
	// The primary encode for this output — one video or one audio track.
	// Experimental.
	Encode EncodeConfiguration `field:"required" json:"encode" yaml:"encode"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// Caption encodes that ride alongside the primary encode.
	//
	// Only in-band caption types are
	// allowed (burn-in, embedded) — out-of-band captions must go in their own output.
	// Default: - no captions on this output.
	//
	// Experimental.
	Captions *[]EncodeConfiguration `field:"optional" json:"captions" yaml:"captions"`
	// A string concatenated to the end of the destination file name.
	// Default: - no name modifier.
	//
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

