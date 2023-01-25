package awsec2


// Configuration options for GenericLinuxImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   genericLinuxImageProps := &genericLinuxImageProps{
//   	userData: userData,
//   }
//
// Experimental.
type GenericLinuxImageProps struct {
	// Initial user data.
	// Experimental.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

