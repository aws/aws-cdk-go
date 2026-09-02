package awsmedialivealpha


// Properties for a CDI input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var cdiInputResolution CdiInputResolution
//   var inputCodec InputCodec
//   var inputMaximumBitrate InputMaximumBitrate
//   var inputResolution InputResolution
//
//   cdiInputSpecificationProps := &CdiInputSpecificationProps{
//   	CdiResolution: cdiInputResolution,
//   	Codec: inputCodec,
//   	MaximumBitrate: inputMaximumBitrate,
//   	Resolution: inputResolution,
//   }
//
// Experimental.
type CdiInputSpecificationProps struct {
	// The codec of the input.
	//
	// This should match the codec of your source content, not the output codec.
	// Default: InputCodec.AVC
	//
	// Experimental.
	Codec InputCodec `field:"optional" json:"codec" yaml:"codec"`
	// The maximum bitrate of the input.
	// Default: InputMaximumBitrate.MAX_20_MBPS
	//
	// Experimental.
	MaximumBitrate InputMaximumBitrate `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The resolution of the input.
	// Default: InputResolution.HD
	//
	// Experimental.
	Resolution InputResolution `field:"optional" json:"resolution" yaml:"resolution"`
	// The maximum resolution of the most demanding CDI input.
	// Default: CdiInputResolution.HD
	//
	// Experimental.
	CdiResolution CdiInputResolution `field:"optional" json:"cdiResolution" yaml:"cdiResolution"`
}

