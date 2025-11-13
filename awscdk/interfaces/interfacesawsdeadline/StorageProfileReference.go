package interfacesawsdeadline


// A reference to a StorageProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageProfileReference := &StorageProfileReference{
//   	FarmId: jsii.String("farmId"),
//   	StorageProfileId: jsii.String("storageProfileId"),
//   }
//
type StorageProfileReference struct {
	// The FarmId of the StorageProfile resource.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The StorageProfileId of the StorageProfile resource.
	StorageProfileId *string `field:"required" json:"storageProfileId" yaml:"storageProfileId"`
}

