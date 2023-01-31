package awscognito


// The email verification style.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	selfSignUpEnabled: jsii.Boolean(true),
//   	userVerification: &userVerificationConfig{
//   		emailSubject: jsii.String("Verify your email for our awesome app!"),
//   		emailBody: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   		emailStyle: cognito.verificationEmailStyle_CODE,
//   		smsMessage: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   	},
//   })
//
// Experimental.
type VerificationEmailStyle string

const (
	// Verify email via code.
	// Experimental.
	VerificationEmailStyle_CODE VerificationEmailStyle = "CODE"
	// Verify email via link.
	// Experimental.
	VerificationEmailStyle_LINK VerificationEmailStyle = "LINK"
)

