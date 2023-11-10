package awswafv2


// A rule statement used to run the rules that are defined in a managed rule group.
//
// To use this, provide the vendor name and the name of the rule group in this statement. You can retrieve the required names through the API call `ListAvailableManagedRuleGroups` .
//
// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . You cannot use a managed rule group statement inside another rule group. You can only use a managed rule group statement as a top-level statement in a rule that you define in a web ACL.
//
// > You are charged additional fees when you use the AWS WAF Bot Control managed rule group `AWSManagedRulesBotControlRuleSet` , the AWS WAF Fraud Control account takeover prevention (ATP) managed rule group `AWSManagedRulesATPRuleSet` , or the AWS WAF Fraud Control account creation fraud prevention (ACFP) managed rule group `AWSManagedRulesACFPRuleSet` . For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html
//
type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// The name of the managed rule group.
	//
	// You use this, along with the vendor name, to identify the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the managed rule group vendor.
	//
	// You use this, along with the rule group name, to identify a rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-vendorname
	//
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-excludedrules
	//
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
	//
	// The rule groups used for intelligent threat mitigation require additional configuration:
	//
	// - Use the `AWSManagedRulesACFPRuleSet` configuration object to configure the account creation fraud prevention managed rule group. The configuration includes the registration and sign-up pages of your application and the locations in the account creation request payload of data, such as the user email and phone number fields.
	// - Use the `AWSManagedRulesATPRuleSet` configuration object to configure the account takeover prevention managed rule group. The configuration includes the sign-in page of your application and the locations in the login request payload of data such as the username and password.
	// - Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-managedrulegroupconfigs
	//
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// Action settings to use in the place of the rule actions that are configured inside the rule group.
	//
	// You specify one override for each rule whose action you want to change.
	//
	// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-ruleactionoverrides
	//
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the managed rule group.
	//
	// Requests are only evaluated by the rule group if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-scopedownstatement
	//
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// The version of the managed rule group to use.
	//
	// If you specify this, the version setting is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's default version, and then keeps the version at the vendor's default when the vendor updates the managed rule group settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

