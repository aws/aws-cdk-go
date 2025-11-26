package interfacesawsecr


// A reference to a SigningConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingConfigurationReference := &SigningConfigurationReference{
//   	RegistryId: jsii.String("registryId"),
//   }
//
type SigningConfigurationReference struct {
	// The RegistryId of the SigningConfiguration resource.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

