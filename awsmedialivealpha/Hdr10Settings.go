package awsmedialivealpha


// HDR10 color space metadata for the input video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hdr10Settings := &Hdr10Settings{
//   	MaxContentLightLevel: jsii.Number(123),
//   	MaxFrameAverageLightLevel: jsii.Number(123),
//   }
//
// Experimental.
type Hdr10Settings struct {
	// Maximum Content Light Level (MaxCLL) — the maximum light level, in nits, of any single pixel in the stream.
	// Default: - not set.
	//
	// Experimental.
	MaxContentLightLevel *float64 `field:"optional" json:"maxContentLightLevel" yaml:"maxContentLightLevel"`
	// Maximum Frame Average Light Level (MaxFALL) — the maximum average light level, in nits, of any single frame in the stream.
	// Default: - not set.
	//
	// Experimental.
	MaxFrameAverageLightLevel *float64 `field:"optional" json:"maxFrameAverageLightLevel" yaml:"maxFrameAverageLightLevel"`
}

