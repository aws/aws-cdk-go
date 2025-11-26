package previewawswafv2mixins


// A single rule, which you can use in a `WebACL` or `RuleGroup` to identify web requests that you want to manage in some way.
//
// Each rule includes one top-level `Statement` that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html
//
type CfnWebACLPropsMixin_RuleProperty struct {
	// The action that AWS WAF should take on a web request when it matches the rule's statement.
	//
	// Settings at the web ACL level can override the rule action setting.
	//
	// This is used only for rules whose statements don't reference a rule group. Rule statements that reference a rule group are `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
	//
	// You must set either this `Action` setting or the rule's `OverrideAction` , but not both:
	//
	// - If the rule statement doesn't reference a rule group, you must set this rule action setting and you must not set the rule's override action setting.
	// - If the rule statement references a rule group, you must not set this action setting, because the actions are already set on the rules inside the rule group. You must set the rule's override action setting to indicate specifically whether to override the actions that are set on the rules in the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
	//
	// If you don't specify this, AWS WAF uses the `CAPTCHA` configuration that's defined for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-captchaconfig
	//
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Specifies how AWS WAF should handle `Challenge` evaluations.
	//
	// If you don't specify this, AWS WAF uses the challenge configuration that's defined for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-challengeconfig
	//
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// The name of the rule.
	//
	// If you change the name of a `Rule` after you create it and you want the rule's metric name to reflect the change, update the metric name in the rule's `VisibilityConfig` settings. AWS WAF doesn't automatically update the metric name when you update the rule name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The override action to apply to the rules in a rule group, instead of the individual rule action settings.
	//
	// This is used only for rules whose statements reference a rule group. Rule statements that reference a rule group are `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
	//
	// Set the override action to none to leave the rule group rule actions in effect. Set it to count to only count matches, regardless of the rule action settings.
	//
	// You must set either this `OverrideAction` setting or the `Action` setting, but not both:
	//
	// - If the rule statement references a rule group, you must set this override action setting and you must not set the rule's action setting.
	// - If the rule statement doesn't reference a rule group, you must set the rule action setting and you must not set the rule's override action setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-overrideaction
	//
	OverrideAction interface{} `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// If you define more than one `Rule` in a `WebACL` , AWS WAF evaluates each request against the `Rules` in order based on the value of `Priority` .
	//
	// AWS WAF processes rules with lower priority first. The priorities don't need to be consecutive, but they must all be different.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Labels to apply to web requests that match the rule match statement.
	//
	// AWS WAF applies fully qualified labels to matching web requests. A fully qualified label is the concatenation of a label namespace and a rule label. The rule's rule group or web ACL defines the label namespace.
	//
	// > Any rule that isn't a rule group reference statement or managed rule group statement can add labels to matching web requests.
	//
	// Rules that run after this rule in the web ACL can match against these labels using a `LabelMatchStatement` .
	//
	// For each label, provide a case-sensitive string containing optional namespaces and a label name, according to the following guidelines:
	//
	// - Separate each component of the label with a colon.
	// - Each namespace or name can have up to 128 characters.
	// - You can specify up to 5 namespaces in a label.
	// - Don't use the following reserved words in your label specification: `aws` , `waf` , `managed` , `rulegroup` , `webacl` , `regexpatternset` , or `ipset` .
	//
	// For example, `myLabelName` or `nameSpace1:nameSpace2:myLabelName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-rulelabels
	//
	RuleLabels interface{} `field:"optional" json:"ruleLabels" yaml:"ruleLabels"`
	// The AWS WAF processing statement for the rule, for example `ByteMatchStatement` or `SizeConstraintStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-statement
	//
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// If you change the name of a `Rule` after you create it and you want the rule's metric name to reflect the change, update the metric name as well. AWS WAF doesn't automatically update the metric name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-visibilityconfig
	//
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

