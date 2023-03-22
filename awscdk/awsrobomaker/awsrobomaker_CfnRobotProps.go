package awsrobomaker


// Properties for defining a `CfnRobot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotProps := &cfnRobotProps{
//   	architecture: jsii.String("architecture"),
//   	greengrassGroupId: jsii.String("greengrassGroupId"),
//
//   	// the properties below are optional
//   	fleet: jsii.String("fleet"),
//   	name: jsii.String("name"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnRobotProps struct {
	// The architecture of the robot.
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Greengrass group associated with the robot.
	GreengrassGroupId *string `field:"required" json:"greengrassGroupId" yaml:"greengrassGroupId"`
	// The Amazon Resource Name (ARN) of the fleet to which the robot will be registered.
	Fleet *string `field:"optional" json:"fleet" yaml:"fleet"`
	// The name of the robot.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map that contains tag keys and tag values that are attached to the robot.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

