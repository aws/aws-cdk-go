package awsecs


// A reference to a CapacityProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderReference := &CapacityProviderReference{
//   	CapacityProviderName: jsii.String("capacityProviderName"),
//   }
//
type CapacityProviderReference struct {
	// The Name of the CapacityProvider resource.
	CapacityProviderName *string `field:"required" json:"capacityProviderName" yaml:"capacityProviderName"`
}

