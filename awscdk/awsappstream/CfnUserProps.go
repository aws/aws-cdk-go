package awsappstream


// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProps := &CfnUserProps{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	FirstName: jsii.String("firstName"),
//   	LastName: jsii.String("lastName"),
//   	MessageAction: jsii.String("messageAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html
//
type CfnUserProps struct {
	// The authentication type for the user.
	//
	// You must specify USERPOOL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The email address of the user.
	//
	// Users' email addresses are case-sensitive. During login, if they specify an email address that doesn't use the same capitalization as the email address specified when their user pool account was created, a "user does not exist" error message displays.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-username
	//
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// The first name, or given name, of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name, or surname, of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The action to take for the welcome email that is sent to a user after the user is created in the user pool.
	//
	// If you specify SUPPRESS, no email is sent. If you specify RESEND, do not specify the first name or last name of the user. If the value is null, the email is sent.
	//
	// > The temporary password in the welcome email is valid for only 7 days. If users don’t set their passwords within 7 days, you must send them a new welcome email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-messageaction
	//
	MessageAction *string `field:"optional" json:"messageAction" yaml:"messageAction"`
}

