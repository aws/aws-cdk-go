package awswafv2


// A single label container.
//
// This is used as an element of a label array in `RuleLabels` inside a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &labelProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnRuleGroup_LabelProperty struct {
	// The label string.
	Name *string `field:"required" json:"name" yaml:"name"`
}

