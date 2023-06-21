package awswafv2


// A rule statement used to run the rules that are defined in a managed rule group.
//
// To use this, provide the vendor name and the name of the rule group in this statement.
//
// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
//
// Example:
//
//
type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// The name of the managed rule group.
	//
	// You use this, along with the vendor name, to identify the rule group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the managed rule group vendor.
	//
	// You use this, along with the rule group name, to identify a rule group.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
	//
	// Use the `AWSManagedRulesATPRuleSet` configuration object for the account takeover prevention managed rule group, to provide information such as the sign-in page of your application and the type of content to accept or reject from the client.
	//
	// Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// Action settings to use in the place of the rule actions that are configured inside the rule group.
	//
	// You specify one override for each rule whose action you want to change.
	//
	// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the managed rule group.
	//
	// Requests are only evaluated by the rule group if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// The version of the managed rule group to use.
	//
	// If you specify this, the version setting is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's default version, and then keeps the version at the vendor's default when the vendor updates the managed rule group settings.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

