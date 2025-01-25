package awscognito


// The password policy settings for a user pool, including complexity, history, and length requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   passwordPolicyProperty := &PasswordPolicyProperty{
//   	MinimumLength: jsii.Number(123),
//   	PasswordHistorySize: jsii.Number(123),
//   	RequireLowercase: jsii.Boolean(false),
//   	RequireNumbers: jsii.Boolean(false),
//   	RequireSymbols: jsii.Boolean(false),
//   	RequireUppercase: jsii.Boolean(false),
//   	TemporaryPasswordValidityDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html
//
type CfnUserPool_PasswordPolicyProperty struct {
	// The minimum length of the password in the policy that you have set.
	//
	// This value can't be less than 6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-minimumlength
	//
	MinimumLength *float64 `field:"optional" json:"minimumLength" yaml:"minimumLength"`
	// The number of previous passwords that you want Amazon Cognito to restrict each user from reusing.
	//
	// Users can't set a password that matches any of `n` previous passwords, where `n` is the value of `PasswordHistorySize` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-passwordhistorysize
	//
	PasswordHistorySize *float64 `field:"optional" json:"passwordHistorySize" yaml:"passwordHistorySize"`
	// The requirement in a password policy that users must include at least one lowercase letter in their password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirelowercase
	//
	RequireLowercase interface{} `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// The requirement in a password policy that users must include at least one number in their password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirenumbers
	//
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// The requirement in a password policy that users must include at least one symbol in their password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requiresymbols
	//
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// The requirement in a password policy that users must include at least one uppercase letter in their password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requireuppercase
	//
	RequireUppercase interface{} `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// The number of days a temporary password is valid in the password policy.
	//
	// If the user doesn't sign in during this time, an administrator must reset their password. Defaults to `7` . If you submit a value of `0` , Amazon Cognito treats it as a null value and sets `TemporaryPasswordValidityDays` to its default value.
	//
	// > When you set `TemporaryPasswordValidityDays` for a user pool, you can no longer set a value for the legacy `UnusedAccountValidityDays` parameter in that user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-temporarypasswordvaliditydays
	//
	TemporaryPasswordValidityDays *float64 `field:"optional" json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

