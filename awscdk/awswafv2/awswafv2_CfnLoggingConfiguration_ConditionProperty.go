package awswafv2


// A single match condition for a `Filter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &conditionProperty{
//   	actionCondition: &actionConditionProperty{
//   		action: jsii.String("action"),
//   	},
//   	labelNameCondition: &labelNameConditionProperty{
//   		labelName: jsii.String("labelName"),
//   	},
//   }
//
type CfnLoggingConfiguration_ConditionProperty struct {
	// A single action condition.
	//
	// This is the action setting that a log record must contain in order to meet the condition.
	ActionCondition interface{} `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// A single label name condition.
	//
	// This is the fully qualified label name that a log record must contain in order to meet the condition. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.
	LabelNameCondition interface{} `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

