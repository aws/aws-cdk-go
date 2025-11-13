package interfacesawsdatasync


// A reference to a LocationEFS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationEFSReference := &LocationEFSReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationEFSReference struct {
	// The LocationArn of the LocationEFS resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

