package awswafv2


// The action to use in the place of the action that results from the rule group evaluation.
//
// Set the override action to none to leave the result of the rule group alone. Set it to count to override the result to count only.
//
// You can only use this for rule statements that reference a rule group, like `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
//
// > This option is usually set to none. It does not affect how the rules in the rule group are evaluated. If you want the rules in the rule group to only count matches, do not use this and instead exclude those rules in your rule group reference statement settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var count interface{}
//   var none interface{}
//
//   overrideActionProperty := &overrideActionProperty{
//   	count: count,
//   	none: none,
//   }
//
type CfnWebACL_OverrideActionProperty struct {
	// Override the rule group evaluation result to count only.
	//
	// > This option is usually set to none. It does not affect how the rules in the rule group are evaluated. If you want the rules in the rule group to only count matches, do not use this and instead exclude those rules in your rule group reference statement settings.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Don't override the rule group evaluation result.
	//
	// This is the most common setting.
	None interface{} `field:"optional" json:"none" yaml:"none"`
}

