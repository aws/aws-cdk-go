package interfacesawsecr


// A reference to a RegistryPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryPolicyReference := &RegistryPolicyReference{
//   	RegistryId: jsii.String("registryId"),
//   }
//
type RegistryPolicyReference struct {
	// The RegistryId of the RegistryPolicy resource.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

