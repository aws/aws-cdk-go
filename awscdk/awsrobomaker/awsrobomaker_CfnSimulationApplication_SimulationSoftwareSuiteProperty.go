package awsrobomaker


// Information about a simulation software suite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simulationSoftwareSuiteProperty := &simulationSoftwareSuiteProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type CfnSimulationApplication_SimulationSoftwareSuiteProperty struct {
	// The name of the simulation software suite.
	//
	// `SimulationRuntime` is the only supported value.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the simulation software suite.
	//
	// Not applicable for `SimulationRuntime` .
	Version *string `field:"optional" json:"version" yaml:"version"`
}

