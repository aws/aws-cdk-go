package awswafv2


// A rule statement used to run the rules that are defined in a `RuleGroup` .
//
// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
//
// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You cannot use a rule group reference statement inside another rule group. You can only reference a rule group as a top-level statement within a rule that you define in a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReferenceStatementProperty := &RuleGroupReferenceStatementProperty{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	ExcludedRules: []interface{}{
//   		&ExcludedRuleProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	RuleActionOverrides: []interface{}{
//   		&RuleActionOverrideProperty{
//   			ActionToUse: &RuleActionProperty{
//   				Allow: &AllowActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Block: &BlockActionProperty{
//   					CustomResponse: &CustomResponseProperty{
//   						ResponseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						ResponseHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Captcha: &CaptchaActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Challenge: &ChallengeActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Count: &CountActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html
//
type CfnWebACL_RuleGroupReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-excludedrules
	//
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Action settings to use in the place of the rule actions that are configured inside the rule group.
	//
	// You specify one override for each rule whose action you want to change.
	//
	// > Verify the rule names in your overrides carefully. With managed rule groups, AWS WAF silently ignores any override that uses an invalid rule name. With customer-owned rule groups, invalid rule names in your overrides will cause web ACL updates to fail. An invalid rule name is any name that doesn't exactly match the case-sensitive name of an existing rule in the rule group.
	//
	// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-ruleactionoverrides
	//
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
}

