package awscognito


// Attributes that will be kept until the user verifies the changed attribute.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	signInAliases: &signInAliases{
//   		username: jsii.Boolean(true),
//   	},
//   	autoVerify: &autoVerifiedAttrs{
//   		email: jsii.Boolean(true),
//   		phone: jsii.Boolean(true),
//   	},
//   	keepOriginal: &keepOriginalAttrs{
//   		email: jsii.Boolean(true),
//   		phone: jsii.Boolean(true),
//   	},
//   })
//
type KeepOriginalAttrs struct {
	// Whether the email address of the user should remain the original value until the new email address is verified.
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether the phone number of the user should remain the original value until the new phone number is verified.
	Phone *bool `field:"optional" json:"phone" yaml:"phone"`
}

