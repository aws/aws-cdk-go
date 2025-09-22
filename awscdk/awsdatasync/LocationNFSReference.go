package awsdatasync


// A reference to a LocationNFS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationNFSReference := &LocationNFSReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationNFSReference struct {
	// The LocationArn of the LocationNFS resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

