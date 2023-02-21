package awswafregional


// A combination of `ByteMatchSet` , `IPSet` , and/or `SqlInjectionMatchSet` objects that identify the web requests that you want to allow, block, or count.
//
// For example, you might create a `Rule` that includes the following predicates:
//
// - An `IPSet` that causes AWS WAF to search for web requests that originate from the IP address `192.0.2.44`
// - A `ByteMatchSet` that causes AWS WAF to search for web requests for which the value of the `User-Agent` header is `BadBot` .
//
// To match the settings in this `Rule` , a request must originate from `192.0.2.44` AND include a `User-Agent` header for which the value is `BadBot` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	action: &actionProperty{
//   		type: jsii.String("type"),
//   	},
//   	priority: jsii.Number(123),
//   	ruleId: jsii.String("ruleId"),
//   }
//
type CfnWebACL_RuleProperty struct {
	// The action that AWS WAF takes when a web request matches all conditions in the rule, such as allow, block, or count the request.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The order in which AWS WAF evaluates the rules in a web ACL.
	//
	// AWS WAF evaluates rules with a lower value before rules with a higher value. The value must be a unique integer. If you have multiple rules in a web ACL, the priority numbers do not need to be consecutive.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The ID of an AWS WAF Regional rule to associate with a web ACL.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
}

