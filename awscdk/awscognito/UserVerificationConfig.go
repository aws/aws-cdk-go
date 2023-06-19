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
// Experimental.
type UserVerificationConfig struct {
	// The email body template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	EmailBody *string `field:"optional" json:"emailBody" yaml:"emailBody"`
	// Emails can be verified either using a code or a link.
	//
	// Learn more at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-email-verification-message-customization.html
	// Experimental.
	EmailStyle VerificationEmailStyle `field:"optional" json:"emailStyle" yaml:"emailStyle"`
	// The email subject template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The message template for the verification SMS sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

