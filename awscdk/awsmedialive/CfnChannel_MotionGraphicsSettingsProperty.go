package awsmedialive


// Settings to enable and configure the motion graphics overlay feature in the channel.
//
// The parent of this entity is MotionGraphicsConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   motionGraphicsSettingsProperty := &MotionGraphicsSettingsProperty{
//   	HtmlMotionGraphicsSettings: &HtmlMotionGraphicsSettingsProperty{
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-motiongraphicssettings.html
//
type CfnChannel_MotionGraphicsSettingsProperty struct {
	// Settings to configure the motion graphics overlay to use an HTML asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-motiongraphicssettings.html#cfn-medialive-channel-motiongraphicssettings-htmlmotiongraphicssettings
	//
	HtmlMotionGraphicsSettings interface{} `field:"optional" json:"htmlMotionGraphicsSettings" yaml:"htmlMotionGraphicsSettings"`
}

