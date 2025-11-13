package interfacesawsecr


// A reference to a ReplicationConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationReference := &ReplicationConfigurationReference{
//   	RegistryId: jsii.String("registryId"),
//   }
//
type ReplicationConfigurationReference struct {
	// The RegistryId of the ReplicationConfiguration resource.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

