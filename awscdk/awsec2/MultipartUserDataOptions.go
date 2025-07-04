package awsec2


// Options for creating `MultipartUserData`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multipartUserDataOptions := &MultipartUserDataOptions{
//   	PartsSeparator: jsii.String("partsSeparator"),
//   }
//
type MultipartUserDataOptions struct {
	// The string used to separate parts in multipart user data archive (it's like MIME boundary).
	//
	// This string should contain [a-zA-Z0-9()+,-./:=?] characters only, and should not be present in any part, or in text content of archive.
	// Default: `+AWS+CDK+User+Data+Separator==`.
	//
	PartsSeparator *string `field:"optional" json:"partsSeparator" yaml:"partsSeparator"`
}

