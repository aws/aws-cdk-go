package awscognito


// The template for verification messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verificationMessageTemplateProperty := &verificationMessageTemplateProperty{
//   	defaultEmailOption: jsii.String("defaultEmailOption"),
//   	emailMessage: jsii.String("emailMessage"),
//   	emailMessageByLink: jsii.String("emailMessageByLink"),
//   	emailSubject: jsii.String("emailSubject"),
//   	emailSubjectByLink: jsii.String("emailSubjectByLink"),
//   	smsMessage: jsii.String("smsMessage"),
//   }
//
type CfnUserPool_VerificationMessageTemplateProperty struct {
	// The default email option.
	DefaultEmailOption *string `field:"optional" json:"defaultEmailOption" yaml:"defaultEmailOption"`
	// The template for email messages that Amazon Cognito sends to your users.
	//
	// You can set an `EmailMessage` template only if the value of [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` . When your [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` , your user pool sends email messages with your own Amazon SES configuration.
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// The email message template for sending a confirmation link to the user.
	//
	// You can set an `EmailMessageByLink` template only if the value of [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` . When your [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` , your user pool sends email messages with your own Amazon SES configuration.
	EmailMessageByLink *string `field:"optional" json:"emailMessageByLink" yaml:"emailMessageByLink"`
	// The subject line for the email message template.
	//
	// You can set an `EmailSubject` template only if the value of [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` . When your [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` , your user pool sends email messages with your own Amazon SES configuration.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The subject line for the email message template for sending a confirmation link to the user.
	//
	// You can set an `EmailSubjectByLink` template only if the value of [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` . When your [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER` , your user pool sends email messages with your own Amazon SES configuration.
	EmailSubjectByLink *string `field:"optional" json:"emailSubjectByLink" yaml:"emailSubjectByLink"`
	// The template for SMS messages that Amazon Cognito sends to your users.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

