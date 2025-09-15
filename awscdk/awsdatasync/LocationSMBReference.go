package awsdatasync


// A reference to a LocationSMB resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationSMBReference := &LocationSMBReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationSMBReference struct {
	// The LocationArn of the LocationSMB resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

