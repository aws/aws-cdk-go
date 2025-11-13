package interfacesawsdatasync


// A reference to a LocationFSxWindows resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationFSxWindowsReference := &LocationFSxWindowsReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationFSxWindowsReference struct {
	// The LocationArn of the LocationFSxWindows resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

