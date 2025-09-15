package awssecurityhub


// A reference to a AutomationRuleV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automationRuleV2Reference := &AutomationRuleV2Reference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type AutomationRuleV2Reference struct {
	// The RuleArn of the AutomationRuleV2 resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

