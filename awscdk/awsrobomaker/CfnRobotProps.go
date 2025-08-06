package awsrobomaker


// Properties for defining a `CfnRobot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotProps := &CfnRobotProps{
//   	Architecture: jsii.String("architecture"),
//   	GreengrassGroupId: jsii.String("greengrassGroupId"),
//
//   	// the properties below are optional
//   	Fleet: jsii.String("fleet"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html
//
type CfnRobotProps struct {
	// The architecture of the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-architecture
	//
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Greengrass group associated with the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-greengrassgroupid
	//
	GreengrassGroupId *string `field:"required" json:"greengrassGroupId" yaml:"greengrassGroupId"`
	// The Amazon Resource Name (ARN) of the fleet to which the robot will be registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-fleet
	//
	Fleet *string `field:"optional" json:"fleet" yaml:"fleet"`
	// The name of the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map that contains tag keys and tag values that are attached to the robot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html#cfn-robomaker-robot-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

