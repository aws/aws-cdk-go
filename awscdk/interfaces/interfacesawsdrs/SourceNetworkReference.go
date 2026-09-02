package interfacesawsdrs


// A reference to a SourceNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceNetworkReference := &SourceNetworkReference{
//   	SourceNetworkArn: jsii.String("sourceNetworkArn"),
//   }
//
type SourceNetworkReference struct {
	// The Arn of the SourceNetwork resource.
	SourceNetworkArn *string `field:"required" json:"sourceNetworkArn" yaml:"sourceNetworkArn"`
}

