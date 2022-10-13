package awssigner


// A reference to a Signing Profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingProfileAttributes := &signingProfileAttributes{
//   	signingProfileName: jsii.String("signingProfileName"),
//   	signingProfileVersion: jsii.String("signingProfileVersion"),
//   }
//
type SigningProfileAttributes struct {
	// The name of signing profile.
	SigningProfileName *string `field:"required" json:"signingProfileName" yaml:"signingProfileName"`
	// The version of signing profile.
	SigningProfileVersion *string `field:"required" json:"signingProfileVersion" yaml:"signingProfileVersion"`
}

