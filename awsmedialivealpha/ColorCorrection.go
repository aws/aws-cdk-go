package awsmedialivealpha


// A color space correction rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var colorSpace ColorSpace
//   var lut Lut
//
//   colorCorrection := &ColorCorrection{
//   	InputColorSpace: colorSpace,
//   	OutputColorSpace: colorSpace,
//
//   	// the properties below are optional
//   	Lut: lut,
//   }
//
// Experimental.
type ColorCorrection struct {
	// The input color space to match.
	// Experimental.
	InputColorSpace ColorSpace `field:"required" json:"inputColorSpace" yaml:"inputColorSpace"`
	// The output color space to convert to.
	// Experimental.
	OutputColorSpace ColorSpace `field:"required" json:"outputColorSpace" yaml:"outputColorSpace"`
	// The 3D LUT file for the color correction.
	//
	// MediaLive reads the LUT from S3 at runtime, so it
	// must be an S3 location — provide it via `Lut.fromBucket()` (which uses the secure `s3ssl://`
	// form and auto-grants the channel role read access) or `Lut.url()` with an `s3://`/`s3ssl://` URL.
	// Default: - no LUT file.
	//
	// Experimental.
	Lut Lut `field:"optional" json:"lut" yaml:"lut"`
}

