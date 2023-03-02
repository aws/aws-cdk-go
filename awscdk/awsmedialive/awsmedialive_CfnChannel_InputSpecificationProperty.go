package awsmedialive


// The input specification for this channel.
//
// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSpecificationProperty := &inputSpecificationProperty{
//   	codec: jsii.String("codec"),
//   	maximumBitrate: jsii.String("maximumBitrate"),
//   	resolution: jsii.String("resolution"),
//   }
//
type CfnChannel_InputSpecificationProperty struct {
	// The codec to include in the input specification for this channel.
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The maximum input bitrate for any input attached to this channel.
	MaximumBitrate *string `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The resolution for any input attached to this channel.
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
}

