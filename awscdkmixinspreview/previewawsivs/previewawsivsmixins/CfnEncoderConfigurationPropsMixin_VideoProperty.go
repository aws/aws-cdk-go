package previewawsivsmixins


// The Video property type describes a stream's video configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoProperty := &VideoProperty{
//   	Bitrate: jsii.Number(123),
//   	Framerate: jsii.Number(123),
//   	Height: jsii.Number(123),
//   	Width: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html
//
type CfnEncoderConfigurationPropsMixin_VideoProperty struct {
	// Bitrate for generated output, in bps.
	//
	// Default: 2500000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html#cfn-ivs-encoderconfiguration-video-bitrate
	//
	// Default: - 2500000.
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Video frame rate, in fps.
	//
	// Default: 30.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html#cfn-ivs-encoderconfiguration-video-framerate
	//
	// Default: - 30.
	//
	Framerate *float64 `field:"optional" json:"framerate" yaml:"framerate"`
	// Video-resolution height.
	//
	// Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html#cfn-ivs-encoderconfiguration-video-height
	//
	// Default: - 720.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Video-resolution width.
	//
	// Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html#cfn-ivs-encoderconfiguration-video-width
	//
	// Default: - 1280.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

