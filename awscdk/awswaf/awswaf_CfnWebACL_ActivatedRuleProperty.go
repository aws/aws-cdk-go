package awswaf


// The `ActivatedRule` object in an `UpdateWebACL` request specifies a `Rule` that you want to insert or delete, the priority of the `Rule` in the `WebACL` , and the action that you want AWS WAF to take when a web request matches the `Rule` ( `ALLOW` , `BLOCK` , or `COUNT` ).
//
// To specify whether to insert or delete a `Rule` , use the `Action` parameter in the `WebACLUpdate` data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activatedRuleProperty := &activatedRuleProperty{
//   	priority: jsii.Number(123),
//   	ruleId: jsii.String("ruleId"),
//
//   	// the properties below are optional
//   	action: &wafActionProperty{
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnWebACL_ActivatedRuleProperty struct {
	// Specifies the order in which the `Rules` in a `WebACL` are evaluated.
	//
	// Rules with a lower value for `Priority` are evaluated before `Rules` with a higher value. The value must be a unique integer. If you add multiple `Rules` to a `WebACL` , the values don't need to be consecutive.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The `RuleId` for a `Rule` .
	//
	// You use `RuleId` to get more information about a `Rule` , update a `Rule` , insert a `Rule` into a `WebACL` or delete a one from a `WebACL` , or delete a `Rule` from AWS WAF .
	//
	// `RuleId` is returned by `CreateRule` and by `ListRules` .
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Specifies the action that Amazon CloudFront or AWS WAF takes when a web request matches the conditions in the `Rule` .
	//
	// Valid values for `Action` include the following:
	//
	// - `ALLOW` : CloudFront responds with the requested object.
	// - `BLOCK` : CloudFront responds with an HTTP 403 (Forbidden) status code.
	// - `COUNT` : AWS WAF increments a counter of requests that match the conditions in the rule and then continues to inspect the web request based on the remaining rules in the web ACL.
	//
	// `ActivatedRule|OverrideAction` applies only when updating or adding a `RuleGroup` to a `WebACL` . In this case, you do not use `ActivatedRule|Action` . For all other update requests, `ActivatedRule|Action` is used instead of `ActivatedRule|OverrideAction` .
	Action interface{} `field:"optional" json:"action" yaml:"action"`
}

