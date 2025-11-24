package mixinsawsmedialive


// The input specification for this channel.
//
// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputSpecificationProperty := &InputSpecificationProperty{
//   	Codec: jsii.String("codec"),
//   	MaximumBitrate: jsii.String("maximumBitrate"),
//   	Resolution: jsii.String("resolution"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputspecification.html
//
type CfnChannelPropsMixin_InputSpecificationProperty struct {
	// The codec to include in the input specification for this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputspecification.html#cfn-medialive-channel-inputspecification-codec
	//
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The maximum input bitrate for any input attached to this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputspecification.html#cfn-medialive-channel-inputspecification-maximumbitrate
	//
	MaximumBitrate *string `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The resolution for any input attached to this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputspecification.html#cfn-medialive-channel-inputspecification-resolution
	//
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
}

