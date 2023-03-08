package awsrobomaker


// Properties for defining a `CfnRobotApplicationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotApplicationVersionProps := &cfnRobotApplicationVersionProps{
//   	application: jsii.String("application"),
//
//   	// the properties below are optional
//   	currentRevisionId: jsii.String("currentRevisionId"),
//   }
//
type CfnRobotApplicationVersionProps struct {
	// The application information for the robot application.
	Application *string `field:"required" json:"application" yaml:"application"`
	// The current revision id for the robot application.
	//
	// If you provide a value and it matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
}

