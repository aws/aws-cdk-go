package interfacesawsrobomaker


// A reference to a SimulationApplicationVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simulationApplicationVersionReference := &SimulationApplicationVersionReference{
//   	SimulationApplicationVersionArn: jsii.String("simulationApplicationVersionArn"),
//   }
//
type SimulationApplicationVersionReference struct {
	// The Arn of the SimulationApplicationVersion resource.
	SimulationApplicationVersionArn *string `field:"required" json:"simulationApplicationVersionArn" yaml:"simulationApplicationVersionArn"`
}

