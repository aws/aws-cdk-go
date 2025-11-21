package awscognito


// Properties for defining a `CfnUserPoolUserToGroupAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolUserToGroupAttachmentProps := &CfnUserPoolUserToGroupAttachmentProps{
//   	GroupName: jsii.String("groupName"),
//   	Username: jsii.String("username"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html
//
type CfnUserPoolUserToGroupAttachmentProps struct {
	// The name of the group that you want to add your user to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-groupname
	//
	GroupName interface{} `field:"required" json:"groupName" yaml:"groupName"`
	// The user's username.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-username
	//
	Username interface{} `field:"required" json:"username" yaml:"username"`
	// The ID of the user pool that contains the group that you want to add the user to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-userpoolid
	//
	UserPoolId interface{} `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

