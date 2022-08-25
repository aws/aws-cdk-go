package awscognito


// The different ways in which users of this pool can sign up or sign in.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	// ...
//   	signInAliases: &signInAliases{
//   		username: jsii.Boolean(true),
//   		email: jsii.Boolean(true),
//   	},
//   	autoVerify: &autoVerifiedAttrs{
//   		email: jsii.Boolean(true),
//   		phone: jsii.Boolean(true),
//   	},
//   })
//
type SignInAliases struct {
	// Whether a user is allowed to sign up or sign in with an email address.
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether a user is allowed to sign up or sign in with a phone number.
	Phone *bool `field:"optional" json:"phone" yaml:"phone"`
	// Whether a user is allowed to sign in with a secondary username, that can be set and modified after sign up.
	//
	// Can only be used in conjunction with `USERNAME`.
	PreferredUsername *bool `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// Whether user is allowed to sign up or sign in with a username.
	Username *bool `field:"optional" json:"username" yaml:"username"`
}

