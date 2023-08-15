package awscognito


// Attributes that can be automatically verified for users in a user pool.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	// ...
//   	SignInAliases: &SignInAliases{
//   		Username: jsii.Boolean(true),
//   		Email: jsii.Boolean(true),
//   	},
//   	AutoVerify: &AutoVerifiedAttrs{
//   		Email: jsii.Boolean(true),
//   		Phone: jsii.Boolean(true),
//   	},
//   })
//
type AutoVerifiedAttrs struct {
	// Whether the email address of the user should be auto verified at sign up.
	//
	// Note: If both `email` and `phone` is set, Cognito only verifies the phone number. To also verify email, see here -
	// https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html
	// Default: - true, if email is turned on for `signIn`. false, otherwise.
	//
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether the phone number of the user should be auto verified at sign up.
	// Default: - true, if phone is turned on for `signIn`. false, otherwise.
	//
	Phone *bool `field:"optional" json:"phone" yaml:"phone"`
}

