package interfacesawsdatasync


// A reference to a LocationObjectStorage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationObjectStorageReference := &LocationObjectStorageReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationObjectStorageReference struct {
	// The LocationArn of the LocationObjectStorage resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

