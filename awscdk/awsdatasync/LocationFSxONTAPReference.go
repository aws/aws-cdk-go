package awsdatasync


// A reference to a LocationFSxONTAP resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationFSxONTAPReference := &LocationFSxONTAPReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationFSxONTAPReference struct {
	// The LocationArn of the LocationFSxONTAP resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

