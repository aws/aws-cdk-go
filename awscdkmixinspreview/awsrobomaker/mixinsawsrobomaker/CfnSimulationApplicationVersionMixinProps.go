package mixinsawsrobomaker


// Properties for CfnSimulationApplicationVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSimulationApplicationVersionMixinProps := &CfnSimulationApplicationVersionMixinProps{
//   	Application: jsii.String("application"),
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html
//
type CfnSimulationApplicationVersionMixinProps struct {
	// The application information for the simulation application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html#cfn-robomaker-simulationapplicationversion-application
	//
	Application *string `field:"optional" json:"application" yaml:"application"`
	// The current revision id for the simulation application.
	//
	// If you provide a value and it matches the latest revision ID, a new version will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html#cfn-robomaker-simulationapplicationversion-currentrevisionid
	//
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
}

