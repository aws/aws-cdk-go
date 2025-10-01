package awscognito


// The email verification style.
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
type VerificationEmailStyle string

const (
	// Verify email via code.
	VerificationEmailStyle_CODE VerificationEmailStyle = "CODE"
	// Verify email via link.
	VerificationEmailStyle_LINK VerificationEmailStyle = "LINK"
)

