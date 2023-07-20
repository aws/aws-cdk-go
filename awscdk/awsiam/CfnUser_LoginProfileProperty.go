package awsiam


// Creates a password for the specified user, giving the user the ability to access AWS services through the AWS Management Console .
//
// For more information about managing passwords, see [Managing Passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loginProfileProperty := &LoginProfileProperty{
//   	Password: jsii.String("password"),
//
//   	// the properties below are optional
//   	PasswordResetRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
//
type CfnUser_LoginProfileProperty struct {
	// The user's password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies whether the user is required to set a new password on next sign-in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-passwordresetrequired
	//
	PasswordResetRequired interface{} `field:"optional" json:"passwordResetRequired" yaml:"passwordResetRequired"`
}

