package awscognito


// The password policy type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   passwordPolicyProperty := &passwordPolicyProperty{
//   	minimumLength: jsii.Number(123),
//   	requireLowercase: jsii.Boolean(false),
//   	requireNumbers: jsii.Boolean(false),
//   	requireSymbols: jsii.Boolean(false),
//   	requireUppercase: jsii.Boolean(false),
//   	temporaryPasswordValidityDays: jsii.Number(123),
//   }
//
type CfnUserPool_PasswordPolicyProperty struct {
	// The minimum length of the password in the policy that you have set.
	//
	// This value can't be less than 6.
	MinimumLength *float64 `field:"optional" json:"minimumLength" yaml:"minimumLength"`
	// In the password policy that you have set, refers to whether you have required users to use at least one lowercase letter in their password.
	RequireLowercase interface{} `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// In the password policy that you have set, refers to whether you have required users to use at least one number in their password.
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// In the password policy that you have set, refers to whether you have required users to use at least one symbol in their password.
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// In the password policy that you have set, refers to whether you have required users to use at least one uppercase letter in their password.
	RequireUppercase interface{} `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// The number of days a temporary password is valid in the password policy.
	//
	// If the user doesn't sign in during this time, an administrator must reset their password.
	//
	// > When you set `TemporaryPasswordValidityDays` for a user pool, you can no longer set a value for the legacy `UnusedAccountValidityDays` parameter in that user pool.
	TemporaryPasswordValidityDays *float64 `field:"optional" json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

