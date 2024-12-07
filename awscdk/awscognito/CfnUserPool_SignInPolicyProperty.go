package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signInPolicyProperty := &SignInPolicyProperty{
//   	AllowedFirstAuthFactors: []*string{
//   		jsii.String("allowedFirstAuthFactors"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html
//
type CfnUserPool_SignInPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-signinpolicy.html#cfn-cognito-userpool-signinpolicy-allowedfirstauthfactors
	//
	AllowedFirstAuthFactors *[]*string `field:"optional" json:"allowedFirstAuthFactors" yaml:"allowedFirstAuthFactors"`
}

