package awssecurityhub


// A reference to a AutomationRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automationRuleReference := &AutomationRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type AutomationRuleReference struct {
	// The RuleArn of the AutomationRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

