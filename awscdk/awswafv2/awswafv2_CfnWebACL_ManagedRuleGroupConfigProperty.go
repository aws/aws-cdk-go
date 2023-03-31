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
//   	awsManagedRulesAtpRuleSet: &aWSManagedRulesATPRuleSetProperty{
//   		loginPath: jsii.String("loginPath"),
//
//   		// the properties below are optional
//   		requestInspection: &requestInspectionProperty{
//   			passwordField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   			payloadType: jsii.String("payloadType"),
//   			usernameField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   		},
//   		responseInspection: &responseInspectionProperty{
//   			bodyContains: &responseInspectionBodyContainsProperty{
//   				failureStrings: []*string{
//   					jsii.String("failureStrings"),
//   				},
//   				successStrings: []*string{
//   					jsii.String("successStrings"),
//   				},
//   			},
//   			header: &responseInspectionHeaderProperty{
//   				failureValues: []*string{
//   					jsii.String("failureValues"),
//   				},
//   				name: jsii.String("name"),
//   				successValues: []*string{
//   					jsii.String("successValues"),
//   				},
//   			},
//   			json: &responseInspectionJsonProperty{
//   				failureValues: []*string{
//   					jsii.String("failureValues"),
//   				},
//   				identifier: jsii.String("identifier"),
//   				successValues: []*string{
//   					jsii.String("successValues"),
//   				},
//   			},
//   			statusCode: &responseInspectionStatusCodeProperty{
//   				failureCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   				successCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
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
	// Additional configuration for using the account takeover prevention (ATP) managed rule group, `AWSManagedRulesATPRuleSet` .
	//
	// Use this to provide login request information to the rule group. For web ACLs that protect CloudFront distributions, use this to also provide the information about how your distribution responds to login requests.
	//
	// This configuration replaces the individual configuration fields in `ManagedRuleGroupConfig` and provides additional feature configuration.
	//
	// For information about using the ATP managed rule group, see [AWS WAF Fraud Control account takeover prevention (ATP) rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-atp.html) and [AWS WAF Fraud Control account takeover prevention (ATP)](https://docs.aws.amazon.com/waf/latest/developerguide/waf-atp.html) in the *AWS WAF Developer Guide* .
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
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

