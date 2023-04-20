package awslambda


// Options when binding a destination to a function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationOptions := &DestinationOptions{
//   	Type: awscdk.Aws_lambda.DestinationType_FAILURE,
//   }
//
// Experimental.
type DestinationOptions struct {
	// The destination type.
	// Experimental.
	Type DestinationType `field:"required" json:"type" yaml:"type"`
}

