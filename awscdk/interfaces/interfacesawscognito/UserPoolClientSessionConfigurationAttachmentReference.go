package interfacesawscognito


// A reference to a UserPoolClientSessionConfigurationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolClientSessionConfigurationAttachmentReference := &UserPoolClientSessionConfigurationAttachmentReference{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolClientSessionConfigurationAttachmentReference struct {
	// The ClientId of the UserPoolClientSessionConfigurationAttachment resource.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The UserPoolId of the UserPoolClientSessionConfigurationAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

