package awswafv2


// Specifies a single rule in a rule group whose action you want to override to `Count` .
//
// When you exclude a rule, AWS WAF evaluates it exactly as it would if the rule action setting were `Count` . This is a useful option for testing the rules in a rule group without modifying how they handle your web traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   excludedRuleProperty := &excludedRuleProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnWebACL_ExcludedRuleProperty struct {
	// The name of the rule whose action you want to override to `Count` .
	Name *string `field:"required" json:"name" yaml:"name"`
}

