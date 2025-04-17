package awswafv2


// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
//
// The rule groups used for intelligent threat mitigation require additional configuration:
//
// - Use the `AWSManagedRulesACFPRuleSet` configuration object to configure the account creation fraud prevention managed rule group. The configuration includes the registration and sign-up pages of your application and the locations in the account creation request payload of data, such as the user email and phone number fields.
// - Use the `AWSManagedRulesATPRuleSet` configuration object to configure the account takeover prevention managed rule group. The configuration includes the sign-in page of your application and the locations in the login request payload of data such as the username and password.
// - Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedRuleGroupConfigProperty := &ManagedRuleGroupConfigProperty{
//   	AwsManagedRulesAcfpRuleSet: &AWSManagedRulesACFPRuleSetProperty{
//   		CreationPath: jsii.String("creationPath"),
//   		RegistrationPagePath: jsii.String("registrationPagePath"),
//   		RequestInspection: &RequestInspectionACFPProperty{
//   			PayloadType: jsii.String("payloadType"),
//
//   			// the properties below are optional
//   			AddressFields: []interface{}{
//   				&FieldIdentifierProperty{
//   					Identifier: jsii.String("identifier"),
//   				},
//   			},
//   			EmailField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   			PasswordField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   			PhoneNumberFields: []interface{}{
//   				&FieldIdentifierProperty{
//   					Identifier: jsii.String("identifier"),
//   				},
//   			},
//   			UsernameField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		EnableRegexInPath: jsii.Boolean(false),
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
//   	AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   		LoginPath: jsii.String("loginPath"),
//
//   		// the properties below are optional
//   		EnableRegexInPath: jsii.Boolean(false),
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
//
//   		// the properties below are optional
//   		EnableMachineLearning: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html
//
type CfnWebACL_ManagedRuleGroupConfigProperty struct {
	// Additional configuration for using the account creation fraud prevention (ACFP) managed rule group, `AWSManagedRulesACFPRuleSet` .
	//
	// Use this to provide account creation request information to the rule group. For web ACLs that protect CloudFront distributions, use this to also provide the information about how your distribution responds to account creation requests.
	//
	// For information about using the ACFP managed rule group, see [AWS WAF Fraud Control account creation fraud prevention (ACFP) rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-acfp.html) and [AWS WAF Fraud Control account creation fraud prevention (ACFP)](https://docs.aws.amazon.com/waf/latest/developerguide/waf-acfp.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesacfpruleset
	//
	AwsManagedRulesAcfpRuleSet interface{} `field:"optional" json:"awsManagedRulesAcfpRuleSet" yaml:"awsManagedRulesAcfpRuleSet"`
	// Additional configuration for using the account takeover prevention (ATP) managed rule group, `AWSManagedRulesATPRuleSet` .
	//
	// Use this to provide login request information to the rule group. For web ACLs that protect CloudFront distributions, use this to also provide the information about how your distribution responds to login requests.
	//
	// This configuration replaces the individual configuration fields in `ManagedRuleGroupConfig` and provides additional feature configuration.
	//
	// For information about using the ATP managed rule group, see [AWS WAF Fraud Control account takeover prevention (ATP) rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-atp.html) and [AWS WAF Fraud Control account takeover prevention (ATP)](https://docs.aws.amazon.com/waf/latest/developerguide/waf-atp.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesatpruleset
	//
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// Additional configuration for using the Bot Control managed rule group.
	//
	// Use this to specify the inspection level that you want to use. For information about using the Bot Control managed rule group, see [AWS WAF Bot Control rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html) and [AWS WAF Bot Control](https://docs.aws.amazon.com/waf/latest/developerguide/waf-bot-control.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesbotcontrolruleset
	//
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// > Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-loginpath
	//
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// > Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-passwordfield
	//
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// > Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-payloadtype
	//
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// > Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-usernamefield
	//
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

