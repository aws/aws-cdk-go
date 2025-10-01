package awslogs


// A reference to a Destination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationReference := &DestinationReference{
//   	DestinationArn: jsii.String("destinationArn"),
//   	DestinationName: jsii.String("destinationName"),
//   }
//
type DestinationReference struct {
	// The ARN of the Destination resource.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The DestinationName of the Destination resource.
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
}

