package awsmediaconnect


// A set of parameters that define the media stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fmtpProperty := &FmtpProperty{
//   	ChannelOrder: jsii.String("channelOrder"),
//   	Colorimetry: jsii.String("colorimetry"),
//   	ExactFramerate: jsii.String("exactFramerate"),
//   	Par: jsii.String("par"),
//   	Range: jsii.String("range"),
//   	ScanMode: jsii.String("scanMode"),
//   	Tcs: jsii.String("tcs"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html
//
type CfnFlow_FmtpProperty struct {
	// The format of the audio channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-channelorder
	//
	ChannelOrder *string `field:"optional" json:"channelOrder" yaml:"channelOrder"`
	// The format used for the representation of color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-colorimetry
	//
	Colorimetry *string `field:"optional" json:"colorimetry" yaml:"colorimetry"`
	// The frame rate for the video stream, in frames/second.
	//
	// For example: 60000/1001.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-exactframerate
	//
	ExactFramerate *string `field:"optional" json:"exactFramerate" yaml:"exactFramerate"`
	// The pixel aspect ratio (PAR) of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-par
	//
	Par *string `field:"optional" json:"par" yaml:"par"`
	// The encoding range of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-range
	//
	Range *string `field:"optional" json:"range" yaml:"range"`
	// The type of compression that was used to smooth the video’s appearance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-scanmode
	//
	ScanMode *string `field:"optional" json:"scanMode" yaml:"scanMode"`
	// The transfer characteristic system (TCS) that is used in the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-tcs
	//
	Tcs *string `field:"optional" json:"tcs" yaml:"tcs"`
}

