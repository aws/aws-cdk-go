package awswafv2


// List of labels used by one or more of the rules of a `RuleGroup` .
//
// This summary object is used for the following rule group lists:
//
// - `AvailableLabels` - Labels that rules add to matching requests. These labels are defined in the `RuleLabels` for a rule.
// - `ConsumedLabels` - Labels that rules match against. These labels are defined in a `LabelMatchStatement` specification, in the `Statement` definition of a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelSummaryProperty := &labelSummaryProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnRuleGroup_LabelSummaryProperty struct {
	// An individual label specification.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

