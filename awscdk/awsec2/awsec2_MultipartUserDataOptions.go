package awsec2


// Options for creating {@link MultipartUserData}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multipartUserDataOptions := &multipartUserDataOptions{
//   	partsSeparator: jsii.String("partsSeparator"),
//   }
//
type MultipartUserDataOptions struct {
	// The string used to separate parts in multipart user data archive (it's like MIME boundary).
	//
	// This string should contain [a-zA-Z0-9()+,-./:=?] characters only, and should not be present in any part, or in text content of archive.
	PartsSeparator *string `field:"optional" json:"partsSeparator" yaml:"partsSeparator"`
}

