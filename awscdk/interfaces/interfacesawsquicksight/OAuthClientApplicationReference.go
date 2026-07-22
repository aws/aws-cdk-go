package interfacesawsquicksight


// A reference to a OAuthClientApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthClientApplicationReference := &OAuthClientApplicationReference{
//   	OAuthClientApplicationArn: jsii.String("oAuthClientApplicationArn"),
//   }
//
type OAuthClientApplicationReference struct {
	// The Arn of the OAuthClientApplication resource.
	OAuthClientApplicationArn *string `field:"required" json:"oAuthClientApplicationArn" yaml:"oAuthClientApplicationArn"`
}

