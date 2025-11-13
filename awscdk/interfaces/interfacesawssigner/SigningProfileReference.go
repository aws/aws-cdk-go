package interfacesawssigner


// A reference to a SigningProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingProfileReference := &SigningProfileReference{
//   	SigningProfileArn: jsii.String("signingProfileArn"),
//   }
//
type SigningProfileReference struct {
	// The Arn of the SigningProfile resource.
	SigningProfileArn *string `field:"required" json:"signingProfileArn" yaml:"signingProfileArn"`
}

