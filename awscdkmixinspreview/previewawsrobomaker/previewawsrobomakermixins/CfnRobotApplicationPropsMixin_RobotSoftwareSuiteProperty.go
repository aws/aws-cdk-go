package previewawsrobomakermixins


// Information about a robot software suite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   robotSoftwareSuiteProperty := &RobotSoftwareSuiteProperty{
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-robotsoftwaresuite.html
//
type CfnRobotApplicationPropsMixin_RobotSoftwareSuiteProperty struct {
	// The name of the robot software suite.
	//
	// `General` is the only supported value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-robotsoftwaresuite.html#cfn-robomaker-robotapplication-robotsoftwaresuite-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the robot software suite.
	//
	// Not applicable for General software suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-robotsoftwaresuite.html#cfn-robomaker-robotapplication-robotsoftwaresuite-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

