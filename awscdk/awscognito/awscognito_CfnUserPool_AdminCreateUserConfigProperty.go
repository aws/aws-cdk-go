package awscognito


// The configuration for `AdminCreateUser` requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adminCreateUserConfigProperty := &adminCreateUserConfigProperty{
//   	allowAdminCreateUserOnly: jsii.Boolean(false),
//   	inviteMessageTemplate: &inviteMessageTemplateProperty{
//   		emailMessage: jsii.String("emailMessage"),
//   		emailSubject: jsii.String("emailSubject"),
//   		smsMessage: jsii.String("smsMessage"),
//   	},
//   	unusedAccountValidityDays: jsii.Number(123),
//   }
//
type CfnUserPool_AdminCreateUserConfigProperty struct {
	// Set to `True` if only the administrator is allowed to create user profiles.
	//
	// Set to `False` if users can sign themselves up via an app.
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// The message template to be used for the welcome message to new users.
	//
	// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
	InviteMessageTemplate interface{} `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
	// The user account expiration limit, in days, after which a new account that hasn't signed in is no longer usable.
	//
	// To reset the account after that time limit, you must call `AdminCreateUser` again, specifying `"RESEND"` for the `MessageAction` parameter. The default value for this parameter is 7.
	//
	// > If you set a value for `TemporaryPasswordValidityDays` in `PasswordPolicy` , that value will be used, and `UnusedAccountValidityDays` will be no longer be an available parameter for that user pool.
	UnusedAccountValidityDays *float64 `field:"optional" json:"unusedAccountValidityDays" yaml:"unusedAccountValidityDays"`
}

