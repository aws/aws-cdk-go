package interfacesawscognito


// A reference to a UserPoolUserToGroupAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolUserToGroupAttachmentReference := &UserPoolUserToGroupAttachmentReference{
//   	GroupName: jsii.String("groupName"),
//   	Username: jsii.String("username"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolUserToGroupAttachmentReference struct {
	// The GroupName of the UserPoolUserToGroupAttachment resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The Username of the UserPoolUserToGroupAttachment resource.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The UserPoolId of the UserPoolUserToGroupAttachment resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

