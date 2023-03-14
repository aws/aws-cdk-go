package awswafv2


// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
//
// Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedRuleGroupConfigProperty := &managedRuleGroupConfigProperty{
//   	awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   		inspectionLevel: jsii.String("inspectionLevel"),
//   	},
//   	loginPath: jsii.String("loginPath"),
//   	passwordField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   	payloadType: jsii.String("payloadType"),
//   	usernameField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   }
//
type CfnWebACL_ManagedRuleGroupConfigProperty struct {
	// Additional configuration for using the Bot Control managed rule group.
	//
	// Use this to specify the inspection level that you want to use. For information about using the Bot Control managed rule group, see [AWS WAF Bot Control rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html) and [AWS WAF Bot Control](https://docs.aws.amazon.com/waf/latest/developerguide/waf-bot-control.html) in the *AWS WAF Developer Guide* .
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// > Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet` .
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// > Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet` `RequestInspection` .
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// > Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet` `RequestInspection` .
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// > Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet` `RequestInspection` .
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

