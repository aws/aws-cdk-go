package awsdatasync


// A reference to a LocationFSxOpenZFS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationFSxOpenZFSReference := &LocationFSxOpenZFSReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationFSxOpenZFSReference struct {
	// The LocationArn of the LocationFSxOpenZFS resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

