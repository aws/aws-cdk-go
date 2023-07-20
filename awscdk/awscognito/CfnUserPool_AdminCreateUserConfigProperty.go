package awscognito


// The configuration for `AdminCreateUser` requests.
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
	// Set to `True` if only the administrator is allowed to create user profiles.
	//
	// Set to `False` if users can sign themselves up via an app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	//
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// The message template to be used for the welcome message to new users.
	//
	// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	//
	InviteMessageTemplate interface{} `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
	// The user account expiration limit, in days, after which a new account that hasn't signed in is no longer usable.
	//
	// To reset the account after that time limit, you must call `AdminCreateUser` again, specifying `"RESEND"` for the `MessageAction` parameter. The default value for this parameter is 7.
	//
	// > If you set a value for `TemporaryPasswordValidityDays` in `PasswordPolicy` , that value will be used, and `UnusedAccountValidityDays` will be no longer be an available parameter for that user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	//
	UnusedAccountValidityDays *float64 `field:"optional" json:"unusedAccountValidityDays" yaml:"unusedAccountValidityDays"`
}

