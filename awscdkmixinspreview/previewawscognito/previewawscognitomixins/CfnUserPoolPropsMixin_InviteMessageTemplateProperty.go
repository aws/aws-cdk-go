package previewawscognitomixins


// The template for the welcome message to new users.
//
// This template must include the `{####}` temporary password placeholder if you are creating users with passwords. If your users don't have passwords, you can omit the placeholder.
//
// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inviteMessageTemplateProperty := &InviteMessageTemplateProperty{
//   	EmailMessage: jsii.String("emailMessage"),
//   	EmailSubject: jsii.String("emailSubject"),
//   	SmsMessage: jsii.String("smsMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html
//
type CfnUserPoolPropsMixin_InviteMessageTemplateProperty struct {
	// The message template for email messages.
	//
	// EmailMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailmessage
	//
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// The subject line for email messages.
	//
	// EmailSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailsubject
	//
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The message template for SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-smsmessage
	//
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

