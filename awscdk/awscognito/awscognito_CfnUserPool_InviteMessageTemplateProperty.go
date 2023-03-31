package awscognito


// The message template to be used for the welcome message to new users.
//
// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inviteMessageTemplateProperty := &inviteMessageTemplateProperty{
//   	emailMessage: jsii.String("emailMessage"),
//   	emailSubject: jsii.String("emailSubject"),
//   	smsMessage: jsii.String("smsMessage"),
//   }
//
type CfnUserPool_InviteMessageTemplateProperty struct {
	// The message template for email messages.
	//
	// EmailMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// The subject line for email messages.
	//
	// EmailSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The message template for SMS messages.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

