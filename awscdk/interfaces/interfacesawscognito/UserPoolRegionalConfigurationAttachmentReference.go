package interfacesawscognito


// A reference to a UserPoolRegionalConfigurationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolRegionalConfigurationAttachmentReference := &UserPoolRegionalConfigurationAttachmentReference{
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolRegionalConfigurationAttachmentReference struct {
	// The UserPoolId of the UserPoolRegionalConfigurationAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

