package awscognito


// The settings for administrator creation of users in a user pool.
//
// Contains settings for allowing user sign-up, customizing invitation messages to new users, and the amount of time before temporary passwords expire.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adminCreateUserConfigProperty := &AdminCreateUserConfigProperty{
//   	AllowAdminCreateUserOnly: jsii.Boolean(false),
//   	InviteMessageTemplate: &InviteMessageTemplateProperty{
//   		EmailMessage: jsii.String("emailMessage"),
//   		EmailSubject: jsii.String("emailSubject"),
//   		SmsMessage: jsii.String("smsMessage"),
//   	},
//   	UnusedAccountValidityDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
//
type CfnUserPool_AdminCreateUserConfigProperty struct {
	// The setting for allowing self-service sign-up.
	//
	// When `true` , only administrators can create new user profiles. When `false` , users can register themselves and create a new user profile with the `SignUp` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	//
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// The template for the welcome message to new users.
	//
	// This template must include the `{####}` temporary password placeholder if you are creating users with passwords. If your users don't have passwords, you can omit the placeholder.
	//
	// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	//
	InviteMessageTemplate interface{} `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
	// This parameter is no longer in use.
	//
	// The password expiration limit in days for administrator-created users. When this time expires, the user can't sign in with their temporary password. To reset the account after that time limit, you must call `AdminCreateUser` again, specifying `RESEND` for the `MessageAction` parameter.
	//
	// The default value for this parameter is 7.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	//
	UnusedAccountValidityDays *float64 `field:"optional" json:"unusedAccountValidityDays" yaml:"unusedAccountValidityDays"`
}

