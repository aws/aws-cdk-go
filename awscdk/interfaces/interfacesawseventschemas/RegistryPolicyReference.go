package interfacesawseventschemas


// A reference to a RegistryPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryPolicyReference := &RegistryPolicyReference{
//   	RegistryPolicyId: jsii.String("registryPolicyId"),
//   }
//
type RegistryPolicyReference struct {
	// The Id of the RegistryPolicy resource.
	RegistryPolicyId *string `field:"required" json:"registryPolicyId" yaml:"registryPolicyId"`
}

