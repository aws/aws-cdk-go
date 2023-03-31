package awswafregional


// Specifies the action AWS WAF takes when a web request matches or doesn't match all rule conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnWebACL_ActionProperty struct {
	// For actions that are associated with a rule, the action that AWS WAF takes when a web request matches all conditions in a rule.
	//
	// For the default action of a web access control list (ACL), the action that AWS WAF takes when a web request doesn't match all conditions in any rule.
	//
	// Valid settings include the following:
	//
	// - `ALLOW` : AWS WAF allows requests
	// - `BLOCK` : AWS WAF blocks requests
	// - `COUNT` : AWS WAF increments a counter of the requests that match all of the conditions in the rule. AWS WAF then continues to inspect the web request based on the remaining rules in the web ACL. You can't specify `COUNT` for the default action for a WebACL.
	Type *string `field:"required" json:"type" yaml:"type"`
}

