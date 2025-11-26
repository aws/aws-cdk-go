package previewawsrobomakermixins


// Properties for CfnRobotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRobotMixinProps := &CfnRobotMixinProps{
//   	Architecture: jsii.String("architecture"),
//   	Fleet: jsii.String("fleet"),
//   	GreengrassGroupId: jsii.String("greengrassGroupId"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html
//
type CfnRobotMixinProps struct {
	// The architecture of the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-architecture
	//
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The Amazon Resource Name (ARN) of the fleet to which the robot will be registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-fleet
	//
	Fleet *string `field:"optional" json:"fleet" yaml:"fleet"`
	// The Greengrass group associated with the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-greengrassgroupid
	//
	GreengrassGroupId *string `field:"optional" json:"greengrassGroupId" yaml:"greengrassGroupId"`
	// The name of the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map that contains tag keys and tag values that are attached to the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

