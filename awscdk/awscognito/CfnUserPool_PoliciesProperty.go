package awscognito


// The policy associated with a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policiesProperty := &PoliciesProperty{
//   	PasswordPolicy: &PasswordPolicyProperty{
//   		MinimumLength: jsii.Number(123),
//   		RequireLowercase: jsii.Boolean(false),
//   		RequireNumbers: jsii.Boolean(false),
//   		RequireSymbols: jsii.Boolean(false),
//   		RequireUppercase: jsii.Boolean(false),
//   		TemporaryPasswordValidityDays: jsii.Number(123),
//   	},
//   }
//
type CfnUserPool_PoliciesProperty struct {
	// The password policy.
	PasswordPolicy interface{} `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
}

