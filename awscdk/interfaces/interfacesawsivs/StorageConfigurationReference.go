package interfacesawsivs


// A reference to a StorageConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConfigurationReference := &StorageConfigurationReference{
//   	StorageConfigurationArn: jsii.String("storageConfigurationArn"),
//   }
//
type StorageConfigurationReference struct {
	// The Arn of the StorageConfiguration resource.
	StorageConfigurationArn *string `field:"required" json:"storageConfigurationArn" yaml:"storageConfigurationArn"`
}

