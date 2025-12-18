package interfacesawsrobomaker


// A reference to a SimulationApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simulationApplicationReference := &SimulationApplicationReference{
//   	SimulationApplicationArn: jsii.String("simulationApplicationArn"),
//   }
//
type SimulationApplicationReference struct {
	// The ARN of the SimulationApplication resource.
	SimulationApplicationArn *string `field:"required" json:"simulationApplicationArn" yaml:"simulationApplicationArn"`
}

