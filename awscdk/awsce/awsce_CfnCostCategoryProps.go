package awsce


// Properties for defining a `CfnCostCategory`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCostCategoryProps := &cfnCostCategoryProps{
//   	name: jsii.String("name"),
//   	rules: jsii.String("rules"),
//   	ruleVersion: jsii.String("ruleVersion"),
//
//   	// the properties below are optional
//   	defaultValue: jsii.String("defaultValue"),
//   	splitChargeRules: jsii.String("splitChargeRules"),
//   }
//
type CfnCostCategoryProps struct {
	// The unique name of the Cost Category.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The array of CostCategoryRule in JSON array format.
	//
	// > Rules are processed in order. If there are multiple rules that match the line item, then the first rule to match is used to determine that Cost Category value.
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// The rule schema version in this particular Cost Category.
	RuleVersion *string `field:"required" json:"ruleVersion" yaml:"ruleVersion"`
	// The default value for the cost category.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The split charge rules that are used to allocate your charges between your Cost Category values.
	SplitChargeRules *string `field:"optional" json:"splitChargeRules" yaml:"splitChargeRules"`
}

