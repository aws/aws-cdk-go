package interfacesawswafv2


// A reference to a RuleGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReference := &RuleGroupReference{
//   	RuleGroupArn: jsii.String("ruleGroupArn"),
//   	RuleGroupId: jsii.String("ruleGroupId"),
//   	RuleGroupName: jsii.String("ruleGroupName"),
//   	Scope: jsii.String("scope"),
//   }
//
type RuleGroupReference struct {
	// The ARN of the RuleGroup resource.
	RuleGroupArn *string `field:"required" json:"ruleGroupArn" yaml:"ruleGroupArn"`
	// The Id of the RuleGroup resource.
	RuleGroupId *string `field:"required" json:"ruleGroupId" yaml:"ruleGroupId"`
	// The Name of the RuleGroup resource.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// The Scope of the RuleGroup resource.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

