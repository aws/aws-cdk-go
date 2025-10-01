package awsdatasync


// A reference to a LocationFSxLustre resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationFSxLustreReference := &LocationFSxLustreReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationFSxLustreReference struct {
	// The LocationArn of the LocationFSxLustre resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

