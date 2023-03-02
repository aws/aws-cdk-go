package awsmedialive


// Settings to enable and configure the motion graphics overlay feature in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   motionGraphicsConfigurationProperty := &motionGraphicsConfigurationProperty{
//   	motionGraphicsInsertion: jsii.String("motionGraphicsInsertion"),
//   	motionGraphicsSettings: &motionGraphicsSettingsProperty{
//   		htmlMotionGraphicsSettings: &htmlMotionGraphicsSettingsProperty{
//   		},
//   	},
//   }
//
type CfnChannel_MotionGraphicsConfigurationProperty struct {
	// Enables or disables the motion graphics overlay feature in the channel.
	MotionGraphicsInsertion *string `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	MotionGraphicsSettings interface{} `field:"optional" json:"motionGraphicsSettings" yaml:"motionGraphicsSettings"`
}

