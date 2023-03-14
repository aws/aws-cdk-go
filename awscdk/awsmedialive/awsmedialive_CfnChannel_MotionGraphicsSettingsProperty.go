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
//   motionGraphicsSettingsProperty := &motionGraphicsSettingsProperty{
//   	htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   	},
//   }
//
type CfnChannel_MotionGraphicsSettingsProperty struct {
	// Settings to configure the motion graphics overlay to use an HTML asset.
	HtmlMotionGraphicsSettings interface{} `field:"optional" json:"htmlMotionGraphicsSettings" yaml:"htmlMotionGraphicsSettings"`
}

