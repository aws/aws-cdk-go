package awscognito


// A list of user pool policies. Contains the policy that sets password-complexity requirements.
//
// This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policiesProperty := &PoliciesProperty{
//   	PasswordPolicy: &PasswordPolicyProperty{
//   		MinimumLength: jsii.Number(123),
//   		PasswordHistorySize: jsii.Number(123),
//   		RequireLowercase: jsii.Boolean(false),
//   		RequireNumbers: jsii.Boolean(false),
//   		RequireSymbols: jsii.Boolean(false),
//   		RequireUppercase: jsii.Boolean(false),
//   		TemporaryPasswordValidityDays: jsii.Number(123),
//   	},
//   	SignInPolicy: &SignInPolicyProperty{
//   		AllowedFirstAuthFactors: []*string{
//   			jsii.String("allowedFirstAuthFactors"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
//
type CfnUserPool_PoliciesProperty struct {
	// The password policy settings for a user pool, including complexity, history, and length requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	//
	PasswordPolicy interface{} `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-signinpolicy
	//
	SignInPolicy interface{} `field:"optional" json:"signInPolicy" yaml:"signInPolicy"`
}

