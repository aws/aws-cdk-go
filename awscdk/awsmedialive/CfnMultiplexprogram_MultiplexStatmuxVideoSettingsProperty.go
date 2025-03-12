package awsmedialive


// Statmux rate control settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexStatmuxVideoSettingsProperty := &MultiplexStatmuxVideoSettingsProperty{
//   	MaximumBitrate: jsii.Number(123),
//   	MinimumBitrate: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexstatmuxvideosettings.html
//
type CfnMultiplexprogram_MultiplexStatmuxVideoSettingsProperty struct {
	// Maximum statmux bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexstatmuxvideosettings.html#cfn-medialive-multiplexprogram-multiplexstatmuxvideosettings-maximumbitrate
	//
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Minimum statmux bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexstatmuxvideosettings.html#cfn-medialive-multiplexprogram-multiplexstatmuxvideosettings-minimumbitrate
	//
	MinimumBitrate *float64 `field:"optional" json:"minimumBitrate" yaml:"minimumBitrate"`
	// The purpose of the priority is to use a combination of the\nmultiplex rate control algorithm and the QVBR capability of the\nencoder to prioritize the video quality of some channels in a\nmultiplex over others.
	//
	// Channels that have a higher priority will\nget higher video quality at the expense of the video quality of\nother channels in the multiplex with lower priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexstatmuxvideosettings.html#cfn-medialive-multiplexprogram-multiplexstatmuxvideosettings-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

