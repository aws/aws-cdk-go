package awsrobomaker


// Properties for defining a `CfnSimulationApplicationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimulationApplicationVersionProps := &CfnSimulationApplicationVersionProps{
//   	Application: jsii.String("application"),
//
//   	// the properties below are optional
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html
//
type CfnSimulationApplicationVersionProps struct {
	// The application information for the simulation application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html#cfn-robomaker-simulationapplicationversion-application
	//
	Application interface{} `field:"required" json:"application" yaml:"application"`
	// The current revision id for the simulation application.
	//
	// If you provide a value and it matches the latest revision ID, a new version will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html#cfn-robomaker-simulationapplicationversion-currentrevisionid
	//
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
}

