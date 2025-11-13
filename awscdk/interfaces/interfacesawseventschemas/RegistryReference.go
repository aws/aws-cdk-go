package interfacesawseventschemas


// A reference to a Registry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryReference := &RegistryReference{
//   	RegistryArn: jsii.String("registryArn"),
//   }
//
type RegistryReference struct {
	// The RegistryArn of the Registry resource.
	RegistryArn *string `field:"required" json:"registryArn" yaml:"registryArn"`
}

