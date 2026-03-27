package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encodingConfigProperty := &EncodingConfigProperty{
//   	EncodingProfile: jsii.String("encodingProfile"),
//   	VideoMaxBitrate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-encodingconfig.html
//
type CfnFlowPropsMixin_EncodingConfigProperty struct {
	// The encoding profile to use when transcoding the NDI source to a Transport Stream.
	//
	// You can change this value while a flow is running.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-encodingconfig.html#cfn-mediaconnect-flow-encodingconfig-encodingprofile
	//
	EncodingProfile *string `field:"optional" json:"encodingProfile" yaml:"encodingProfile"`
	// The maximum video bitrate to use when transcoding the NDI source to a Transport Stream.
	//
	// This parameter enables you to override the default video bitrate within the encoding profile's supported range. The supported range is 10,000,000 - 50,000,000 bits per second (bps). If you do not specify a value, MediaConnect uses the default value of 20,000,000 bps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-encodingconfig.html#cfn-mediaconnect-flow-encodingconfig-videomaxbitrate
	//
	VideoMaxBitrate *float64 `field:"optional" json:"videoMaxBitrate" yaml:"videoMaxBitrate"`
}

