package awscognito


// The policy associated with a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policiesProperty := &policiesProperty{
//   	passwordPolicy: &passwordPolicyProperty{
//   		minimumLength: jsii.Number(123),
//   		requireLowercase: jsii.Boolean(false),
//   		requireNumbers: jsii.Boolean(false),
//   		requireSymbols: jsii.Boolean(false),
//   		requireUppercase: jsii.Boolean(false),
//   		temporaryPasswordValidityDays: jsii.Number(123),
//   	},
//   }
//
type CfnUserPool_PoliciesProperty struct {
	// The password policy.
	PasswordPolicy interface{} `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
}

