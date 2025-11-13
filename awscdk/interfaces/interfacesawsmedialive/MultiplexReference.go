package interfacesawsmedialive


// A reference to a Multiplex resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexReference := &MultiplexReference{
//   	MultiplexArn: jsii.String("multiplexArn"),
//   	MultiplexId: jsii.String("multiplexId"),
//   }
//
type MultiplexReference struct {
	// The ARN of the Multiplex resource.
	MultiplexArn *string `field:"required" json:"multiplexArn" yaml:"multiplexArn"`
	// The Id of the Multiplex resource.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
}

