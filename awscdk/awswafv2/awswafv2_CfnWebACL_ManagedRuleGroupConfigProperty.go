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
//   managedRuleGroupConfigProperty := &ManagedRuleGroupConfigProperty{
//   	AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   		LoginPath: jsii.String("loginPath"),
//
//   		// the properties below are optional
//   		RequestInspection: &RequestInspectionProperty{
//   			PasswordField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   			PayloadType: jsii.String("payloadType"),
//   			UsernameField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   		ResponseInspection: &ResponseInspectionProperty{
//   			BodyContains: &ResponseInspectionBodyContainsProperty{
//   				FailureStrings: []*string{
//   					jsii.String("failureStrings"),
//   				},
//   				SuccessStrings: []*string{
//   					jsii.String("successStrings"),
//   				},
//   			},
//   			Header: &ResponseInspectionHeaderProperty{
//   				FailureValues: []*string{
//   					jsii.String("failureValues"),
//   				},
//   				Name: jsii.String("name"),
//   				SuccessValues: []*string{
//   					jsii.String("successValues"),
//   				},
//   			},
//   			Json: &ResponseInspectionJsonProperty{
//   				FailureValues: []*string{
//   					jsii.String("failureValues"),
//   				},
//   				Identifier: jsii.String("identifier"),
//   				SuccessValues: []*string{
//   					jsii.String("successValues"),
//   				},
//   			},
//   			StatusCode: &ResponseInspectionStatusCodeProperty{
//   				FailureCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   				SuccessCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   		InspectionLevel: jsii.String("inspectionLevel"),
//   	},
//   	LoginPath: jsii.String("loginPath"),
//   	PasswordField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   	PayloadType: jsii.String("payloadType"),
//   	UsernameField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
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

