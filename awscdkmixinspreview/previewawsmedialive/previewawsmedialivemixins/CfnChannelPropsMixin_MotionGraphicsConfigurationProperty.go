package previewawsmedialivemixins


// Settings to enable and configure the motion graphics overlay feature in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   motionGraphicsConfigurationProperty := &MotionGraphicsConfigurationProperty{
//   	MotionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   	MotionGraphicsSettings: &MotionGraphicsSettingsProperty{
//   		HtmlMotionGraphicsSettings: &HtmlMotionGraphicsSettingsProperty{
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-motiongraphicsconfiguration.html
//
type CfnChannelPropsMixin_MotionGraphicsConfigurationProperty struct {
	// Enables or disables the motion graphics overlay feature in the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-motiongraphicsconfiguration.html#cfn-medialive-channel-motiongraphicsconfiguration-motiongraphicsinsertion
	//
	MotionGraphicsInsertion *string `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-motiongraphicsconfiguration.html#cfn-medialive-channel-motiongraphicsconfiguration-motiongraphicssettings
	//
	MotionGraphicsSettings interface{} `field:"optional" json:"motionGraphicsSettings" yaml:"motionGraphicsSettings"`
}

