package awss3


// A reference to a StorageLens resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLensReference := &StorageLensReference{
//   	StorageLensId: jsii.String("storageLensId"),
//   }
//
type StorageLensReference struct {
	// The StorageLensConfiguration/Id of the StorageLens resource.
	StorageLensId *string `field:"required" json:"storageLensId" yaml:"storageLensId"`
}

