package awscognito


// The user-pool treatment for MFA with a passkey.
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
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html#amazon-cognito-user-pools-authentication-flow-methods-passkey
//
type PasskeyUserVerification string

const (
	// Passkey MFA is preferred.
	PasskeyUserVerification_PREFERRED PasskeyUserVerification = "PREFERRED"
	// Passkey MFA is required.
	PasskeyUserVerification_REQUIRED PasskeyUserVerification = "REQUIRED"
)

