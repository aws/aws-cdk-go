package awsce


// Properties for defining a `CfnCostCategory`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCostCategoryProps := &CfnCostCategoryProps{
//   	Name: jsii.String("name"),
//   	Rules: jsii.String("rules"),
//   	RuleVersion: jsii.String("ruleVersion"),
//
//   	// the properties below are optional
//   	DefaultValue: jsii.String("defaultValue"),
//   	SplitChargeRules: jsii.String("splitChargeRules"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html
//
type CfnCostCategoryProps struct {
	// The unique name of the Cost Category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The array of CostCategoryRule in JSON array format.
	//
	// > Rules are processed in order. If there are multiple rules that match the line item, then the first rule to match is used to determine that Cost Category value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-rules
	//
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// The rule schema version in this particular Cost Category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-ruleversion
	//
	RuleVersion *string `field:"required" json:"ruleVersion" yaml:"ruleVersion"`
	// The default value for the cost category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The split charge rules that are used to allocate your charges between your Cost Category values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-splitchargerules
	//
	SplitChargeRules *string `field:"optional" json:"splitChargeRules" yaml:"splitChargeRules"`
}

