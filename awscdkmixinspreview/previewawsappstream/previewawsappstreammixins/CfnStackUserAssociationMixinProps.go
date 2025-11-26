package previewawsappstreammixins


// Properties for CfnStackUserAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStackUserAssociationMixinProps := &CfnStackUserAssociationMixinProps{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	SendEmailNotification: jsii.Boolean(false),
//   	StackName: jsii.String("stackName"),
//   	UserName: jsii.String("userName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html
//
type CfnStackUserAssociationMixinProps struct {
	// The authentication type for the user who is associated with the stack.
	//
	// You must specify USERPOOL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Specifies whether a welcome email is sent to a user after the user is created in the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-sendemailnotification
	//
	SendEmailNotification interface{} `field:"optional" json:"sendEmailNotification" yaml:"sendEmailNotification"`
	// The name of the stack that is associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-stackname
	//
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// The email address of the user who is associated with the stack.
	//
	// > Users' email addresses are case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

