package awsappstream


// Properties for defining a `CfnStackUserAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackUserAssociationProps := &CfnStackUserAssociationProps{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	StackName: jsii.String("stackName"),
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	SendEmailNotification: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html
//
type CfnStackUserAssociationProps struct {
	// The authentication type for the user who is associated with the stack.
	//
	// You must specify USERPOOL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The name of the stack that is associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-stackname
	//
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The email address of the user who is associated with the stack.
	//
	// > Users' email addresses are case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-username
	//
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Specifies whether a welcome email is sent to a user after the user is created in the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-sendemailnotification
	//
	SendEmailNotification interface{} `field:"optional" json:"sendEmailNotification" yaml:"sendEmailNotification"`
}

