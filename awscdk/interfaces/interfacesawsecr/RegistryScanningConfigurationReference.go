package interfacesawsecr


// A reference to a RegistryScanningConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryScanningConfigurationReference := &RegistryScanningConfigurationReference{
//   	RegistryId: jsii.String("registryId"),
//   }
//
type RegistryScanningConfigurationReference struct {
	// The RegistryId of the RegistryScanningConfiguration resource.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

