package interfacesawsbedrockagentcore


// A reference to a CapacityProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderReference := &CapacityProviderReference{
//   	CapacityProviderArn: jsii.String("capacityProviderArn"),
//   }
//
type CapacityProviderReference struct {
	// The Arn of the CapacityProvider resource.
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
}

