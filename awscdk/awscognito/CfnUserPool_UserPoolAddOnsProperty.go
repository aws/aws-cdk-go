package awscognito


// User pool add-ons.
//
// Contains settings for activation of threat protection. To log user security information but take no action, set to `AUDIT` . To configure automatic security responses to risky traffic to your user pool, set to `ENFORCED` .
//
// For more information, see [Adding advanced security to a user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html) . To activate this setting, your user pool must be on the [Plus tier](https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-plus.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolAddOnsProperty := &UserPoolAddOnsProperty{
//   	AdvancedSecurityAdditionalFlows: &AdvancedSecurityAdditionalFlowsProperty{
//   		CustomAuthMode: jsii.String("customAuthMode"),
//   	},
//   	AdvancedSecurityMode: jsii.String("advancedSecurityMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html
//
type CfnUserPool_UserPoolAddOnsProperty struct {
	// Threat protection configuration options for additional authentication types in your user pool, including custom authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecurityadditionalflows
	//
	AdvancedSecurityAdditionalFlows interface{} `field:"optional" json:"advancedSecurityAdditionalFlows" yaml:"advancedSecurityAdditionalFlows"`
	// The operating mode of threat protection for standard authentication types in your user pool, including username-password and secure remote password (SRP) authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecuritymode
	//
	AdvancedSecurityMode *string `field:"optional" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

