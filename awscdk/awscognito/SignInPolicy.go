package awscognito


// Sign-in policy for User Pools.
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
type SignInPolicy struct {
	// The types of authentication that you want to allow for users' first authentication prompt.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flows-selection-sdk.html#authentication-flows-selection-choice
	//
	// Default: - Password only.
	//
	AllowedFirstAuthFactors *AllowedFirstAuthFactors `field:"optional" json:"allowedFirstAuthFactors" yaml:"allowedFirstAuthFactors"`
}

