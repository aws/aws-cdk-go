package interfacesawscognito


// A reference to a UserPoolUICustomizationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolUICustomizationAttachmentReference := &UserPoolUICustomizationAttachmentReference{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolUICustomizationAttachmentReference struct {
	// The ClientId of the UserPoolUICustomizationAttachment resource.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The UserPoolId of the UserPoolUICustomizationAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

