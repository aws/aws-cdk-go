package awscognito


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-advancedsecurityadditionalflows.html#cfn-cognito-userpool-advancedsecurityadditionalflows-customauthmode
	//
	CustomAuthMode *string `field:"optional" json:"customAuthMode" yaml:"customAuthMode"`
}

