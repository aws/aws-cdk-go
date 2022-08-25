package awscognito


// Properties for defining a `CfnUserPoolUserToGroupAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolUserToGroupAttachmentProps := &cfnUserPoolUserToGroupAttachmentProps{
//   	groupName: jsii.String("groupName"),
//   	username: jsii.String("username"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnUserPoolUserToGroupAttachmentProps struct {
	// The group name.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The username for the user.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The user pool ID for the user pool.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

