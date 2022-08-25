package awsec2


// Configuration options for WindowsImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   windowsImageProps := &windowsImageProps{
//   	userData: userData,
//   }
//
type WindowsImageProps struct {
	// Initial user data.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

