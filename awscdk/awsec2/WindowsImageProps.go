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
//   windowsImageProps := &WindowsImageProps{
//   	UserData: userData,
//   }
//
type WindowsImageProps struct {
	// Initial user data.
	// Default: - Empty UserData for Windows machines.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

