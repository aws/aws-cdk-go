package awsec2


// Configuration options for GenericWindowsImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   genericWindowsImageProps := &genericWindowsImageProps{
//   	userData: userData,
//   }
//
type GenericWindowsImageProps struct {
	// Initial user data.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

