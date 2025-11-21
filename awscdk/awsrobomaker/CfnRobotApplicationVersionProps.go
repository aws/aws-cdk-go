package awsrobomaker


// Properties for defining a `CfnRobotApplicationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotApplicationVersionProps := &CfnRobotApplicationVersionProps{
//   	Application: jsii.String("application"),
//
//   	// the properties below are optional
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html
//
type CfnRobotApplicationVersionProps struct {
	// The application information for the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-application
	//
	Application interface{} `field:"required" json:"application" yaml:"application"`
	// The current revision id for the robot application.
	//
	// If you provide a value and it matches the latest revision ID, a new version will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-currentrevisionid
	//
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
}

