package awscognito


// The different ways in which a user pool's MFA enforcement can be configured.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	Mfa: cognito.Mfa_REQUIRED,
//   	MfaSecondFactor: &MfaSecondFactor{
//   		Sms: jsii.Boolean(true),
//   		Otp: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html
//
type Mfa string

const (
	// Users are not required to use MFA for sign in, and cannot configure one.
	Mfa_OFF Mfa = "OFF"
	// Users are not required to use MFA for sign in, but can configure one if they so choose to.
	Mfa_OPTIONAL Mfa = "OPTIONAL"
	// Users are required to configure an MFA, and have to use it to sign in.
	Mfa_REQUIRED Mfa = "REQUIRED"
)

