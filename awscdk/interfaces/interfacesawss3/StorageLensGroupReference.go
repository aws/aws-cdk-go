package interfacesawss3


// A reference to a StorageLensGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLensGroupReference := &StorageLensGroupReference{
//   	StorageLensGroupArn: jsii.String("storageLensGroupArn"),
//   	StorageLensGroupName: jsii.String("storageLensGroupName"),
//   }
//
type StorageLensGroupReference struct {
	// The ARN of the StorageLensGroup resource.
	StorageLensGroupArn *string `field:"required" json:"storageLensGroupArn" yaml:"storageLensGroupArn"`
	// The Name of the StorageLensGroup resource.
	StorageLensGroupName *string `field:"required" json:"storageLensGroupName" yaml:"storageLensGroupName"`
}

