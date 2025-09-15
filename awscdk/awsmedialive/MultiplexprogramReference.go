package awsmedialive


// A reference to a Multiplexprogram resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexprogramReference := &MultiplexprogramReference{
//   	MultiplexId: jsii.String("multiplexId"),
//   	ProgramName: jsii.String("programName"),
//   }
//
type MultiplexprogramReference struct {
	// The MultiplexId of the Multiplexprogram resource.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
	// The ProgramName of the Multiplexprogram resource.
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
}

