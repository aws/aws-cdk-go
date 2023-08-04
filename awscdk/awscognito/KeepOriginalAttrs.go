package awscognito


// Attributes that will be kept until the user verifies the changed attribute.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	SignInAliases: &SignInAliases{
//   		Username: jsii.Boolean(true),
//   	},
//   	AutoVerify: &AutoVerifiedAttrs{
//   		Email: jsii.Boolean(true),
//   		Phone: jsii.Boolean(true),
//   	},
//   	KeepOriginal: &KeepOriginalAttrs{
//   		Email: jsii.Boolean(true),
//   		Phone: jsii.Boolean(true),
//   	},
//   })
//
type KeepOriginalAttrs struct {
	// Whether the email address of the user should remain the original value until the new email address is verified.
	// Default: - false.
	//
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether the phone number of the user should remain the original value until the new phone number is verified.
	// Default: - false.
	//
	Phone *bool `field:"optional" json:"phone" yaml:"phone"`
}

