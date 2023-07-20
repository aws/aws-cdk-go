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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
//
type CfnUserPool_PoliciesProperty struct {
	// The password policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	//
	PasswordPolicy interface{} `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
}

