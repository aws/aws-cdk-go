package awslambda


// A destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfig := &DestinationConfig{
//   	Destination: jsii.String("destination"),
//   }
//
type DestinationConfig struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

