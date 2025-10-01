package awscognito


// A reference to a UserPoolRiskConfigurationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolRiskConfigurationAttachmentReference := &UserPoolRiskConfigurationAttachmentReference{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolRiskConfigurationAttachmentReference struct {
	// The ClientId of the UserPoolRiskConfigurationAttachment resource.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The UserPoolId of the UserPoolRiskConfigurationAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

