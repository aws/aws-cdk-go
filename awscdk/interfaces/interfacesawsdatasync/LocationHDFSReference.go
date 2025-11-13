package interfacesawsdatasync


// A reference to a LocationHDFS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationHDFSReference := &LocationHDFSReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationHDFSReference struct {
	// The LocationArn of the LocationHDFS resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

