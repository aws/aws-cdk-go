package awscognito


// User pool configuration for user self sign up.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	SelfSignUpEnabled: jsii.Boolean(true),
//   	UserVerification: &UserVerificationConfig{
//   		EmailSubject: jsii.String("Verify your email for our awesome app!"),
//   		EmailBody: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   		EmailStyle: cognito.VerificationEmailStyle_CODE,
//   		SmsMessage: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   	},
//   })
//
type UserVerificationConfig struct {
	// The email body template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Default: - 'The verification code to your new account is {####}' if VerificationEmailStyle.CODE is chosen,
	// 'Verify your account by clicking on {##Verify Email##}' if VerificationEmailStyle.LINK is chosen.
	//
	EmailBody *string `field:"optional" json:"emailBody" yaml:"emailBody"`
	// Emails can be verified either using a code or a link.
	//
	// Learn more at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-email-verification-message-customization.html
	// Default: VerificationEmailStyle.CODE
	//
	EmailStyle VerificationEmailStyle `field:"optional" json:"emailStyle" yaml:"emailStyle"`
	// The email subject template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Default: 'Verify your new account'.
	//
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The message template for the verification SMS sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Default: - 'The verification code to your new account is {####}' if VerificationEmailStyle.CODE is chosen,
	// not configured if VerificationEmailStyle.LINK is chosen
	//
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

