package mixinsawscognito


// The policy for allowed types of authentication in a user pool.
//
// To activate this setting, your user pool must be in the [Essentials tier](https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-essentials.html) or higher.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   signInPolicyProperty := &SignInPolicyProperty{
//   	AllowedFirstAuthFactors: []*string{
//   		jsii.String("allowedFirstAuthFactors"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html
//
type CfnUserPoolPropsMixin_SignInPolicyProperty struct {
	// The sign-in methods that a user pool supports as the first factor.
	//
	// You can permit users to start authentication with a standard username and password, or with other one-time password and hardware factors.
	//
	// Supports values of `EMAIL_OTP` , `SMS_OTP` , `WEB_AUTHN` and `PASSWORD` ,.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html#cfn-cognito-userpool-signinpolicy-allowedfirstauthfactors
	//
	AllowedFirstAuthFactors *[]*string `field:"optional" json:"allowedFirstAuthFactors" yaml:"allowedFirstAuthFactors"`
}

