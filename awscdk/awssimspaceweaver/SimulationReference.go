package awssimspaceweaver


// A reference to a Simulation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simulationReference := &SimulationReference{
//   	SimulationName: jsii.String("simulationName"),
//   }
//
type SimulationReference struct {
	// The Name of the Simulation resource.
	SimulationName *string `field:"required" json:"simulationName" yaml:"simulationName"`
}

