package interfacesawscognito


// A reference to a UserPoolSessionConfigurationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolSessionConfigurationAttachmentReference := &UserPoolSessionConfigurationAttachmentReference{
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolSessionConfigurationAttachmentReference struct {
	// The UserPoolId of the UserPoolSessionConfigurationAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

