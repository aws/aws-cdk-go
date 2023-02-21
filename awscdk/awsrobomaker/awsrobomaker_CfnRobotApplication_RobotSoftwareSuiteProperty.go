package awsrobomaker


// Information about a robot software suite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   robotSoftwareSuiteProperty := &robotSoftwareSuiteProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type CfnRobotApplication_RobotSoftwareSuiteProperty struct {
	// The name of the robot software suite.
	//
	// `General` is the only supported value.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the robot software suite.
	//
	// Not applicable for General software suite.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

