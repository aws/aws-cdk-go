package awslambda


// Options when binding a destination to a function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationOptions := &destinationOptions{
//   	type: awscdk.Aws_lambda.destinationType_FAILURE,
//   }
//
type DestinationOptions struct {
	// The destination type.
	Type DestinationType `field:"required" json:"type" yaml:"type"`
}

