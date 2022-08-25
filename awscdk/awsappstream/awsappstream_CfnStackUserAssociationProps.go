package awsappstream


// Properties for defining a `CfnStackUserAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackUserAssociationProps := &cfnStackUserAssociationProps{
//   	authenticationType: jsii.String("authenticationType"),
//   	stackName: jsii.String("stackName"),
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	sendEmailNotification: jsii.Boolean(false),
//   }
//
type CfnStackUserAssociationProps struct {
	// The authentication type for the user who is associated with the stack.
	//
	// You must specify USERPOOL.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The name of the stack that is associated with the user.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The email address of the user who is associated with the stack.
	//
	// > Users' email addresses are case-sensitive.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Specifies whether a welcome email is sent to a user after the user is created in the user pool.
	SendEmailNotification interface{} `field:"optional" json:"sendEmailNotification" yaml:"sendEmailNotification"`
}

