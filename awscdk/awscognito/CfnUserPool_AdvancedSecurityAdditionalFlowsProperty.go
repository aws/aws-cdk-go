package awscognito


// Advanced security configuration options for additional authentication types in your user pool, including custom authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedSecurityAdditionalFlowsProperty := &AdvancedSecurityAdditionalFlowsProperty{
//   	CustomAuthMode: jsii.String("customAuthMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-advancedsecurityadditionalflows.html
//
type CfnUserPool_AdvancedSecurityAdditionalFlowsProperty struct {
	// The operating mode of advanced security features in custom authentication with [Custom authentication challenge Lambda triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-advancedsecurityadditionalflows.html#cfn-cognito-userpool-advancedsecurityadditionalflows-customauthmode
	//
	CustomAuthMode *string `field:"optional" json:"customAuthMode" yaml:"customAuthMode"`
}

