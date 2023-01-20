package awswafv2


// A rule statement used to run the rules that are defined in a `RuleGroup` .
//
// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
//
// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReferenceStatementProperty := &ruleGroupReferenceStatementProperty{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	excludedRules: []interface{}{
//   		&excludedRuleProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	ruleActionOverrides: []interface{}{
//   		&ruleActionOverrideProperty{
//   			actionToUse: &ruleActionProperty{
//   				allow: &allowActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				block: &blockActionProperty{
//   					customResponse: &customResponseProperty{
//   						responseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						responseHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				captcha: &captchaActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				challenge: &challengeActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				count: &countActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnWebACL_RuleGroupReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the entity.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The rules in the referenced rule group whose actions are set to `Count` .
	//
	// When you exclude a rule, AWS WAF evaluates it exactly as it would if the rule action setting were `Count` . This is a useful option for testing the rules in a rule group without modifying how they handle your web traffic.
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// `CfnWebACL.RuleGroupReferenceStatementProperty.RuleActionOverrides`.
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
}

