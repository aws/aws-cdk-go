package awscognito


// User pool add-ons.
//
// Contains settings for activation of advanced security features. To log user security information but take no action, set to `AUDIT` . To configure automatic security responses to risky traffic to your user pool, set to `ENFORCED` .
//
// For more information, see [Adding advanced security to a user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html) .
//
// This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecurityadditionalflows
	//
	AdvancedSecurityAdditionalFlows interface{} `field:"optional" json:"advancedSecurityAdditionalFlows" yaml:"advancedSecurityAdditionalFlows"`
	// The operating mode of advanced security features for standard authentication types in your user pool, including username-password and secure remote password (SRP) authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecuritymode
	//
	AdvancedSecurityMode *string `field:"optional" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

