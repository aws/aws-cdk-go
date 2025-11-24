package mixinsawsrobomaker


// Information about a simulation software suite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   simulationSoftwareSuiteProperty := &SimulationSoftwareSuiteProperty{
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-simulationsoftwaresuite.html
//
type CfnSimulationApplicationPropsMixin_SimulationSoftwareSuiteProperty struct {
	// The name of the simulation software suite.
	//
	// `SimulationRuntime` is the only supported value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-simulationsoftwaresuite.html#cfn-robomaker-simulationapplication-simulationsoftwaresuite-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the simulation software suite.
	//
	// Not applicable for `SimulationRuntime` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-simulationsoftwaresuite.html#cfn-robomaker-simulationapplication-simulationsoftwaresuite-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

