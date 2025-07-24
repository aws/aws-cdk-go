package awscognito


// The types of authentication that you want to allow for users' first authentication prompt.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	SignInPolicy: &SignInPolicy{
//   		AllowedFirstAuthFactors: &AllowedFirstAuthFactors{
//   			Password: jsii.Boolean(true),
//   			Passkey: jsii.Boolean(true),
//   		},
//   	},
//   	PasskeyRelyingPartyId: jsii.String("auth.example.com"),
//   	PasskeyUserVerification: cognito.PasskeyUserVerification_REQUIRED,
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flows-selection-sdk.html#authentication-flows-selection-choice
//
type AllowedFirstAuthFactors struct {
	// Whether the password authentication is allowed.
	//
	// This must be true.
	Password *bool `field:"required" json:"password" yaml:"password"`
	// Whether the email message one-time password is allowed.
	// Default: false.
	//
	EmailOtp *bool `field:"optional" json:"emailOtp" yaml:"emailOtp"`
	// Whether the Passkey (WebAuthn) is allowed.
	// Default: false.
	//
	Passkey *bool `field:"optional" json:"passkey" yaml:"passkey"`
	// Whether the SMS message one-time password is allowed.
	// Default: false.
	//
	SmsOtp *bool `field:"optional" json:"smsOtp" yaml:"smsOtp"`
}

