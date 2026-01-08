package interfacesawseventschemas


// A reference to a RegistryPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryPolicyReference := &RegistryPolicyReference{
//   	RegistryName: jsii.String("registryName"),
//   }
//
type RegistryPolicyReference struct {
	// The RegistryName of the RegistryPolicy resource.
	RegistryName *string `field:"required" json:"registryName" yaml:"registryName"`
}

