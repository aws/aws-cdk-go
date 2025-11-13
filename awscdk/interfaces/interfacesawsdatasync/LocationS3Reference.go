package interfacesawsdatasync


// A reference to a LocationS3 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationS3Reference := &LocationS3Reference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationS3Reference struct {
	// The LocationArn of the LocationS3 resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

