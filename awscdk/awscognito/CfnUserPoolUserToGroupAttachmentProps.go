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
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The username of the user that you want to query or modify.
	//
	// The value of this parameter is typically your user's
	// username, but it can be any of their alias attributes. If `username` isn't an alias attribute in your user pool, you can also use their `sub` in this request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
	// The user pool ID for the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

