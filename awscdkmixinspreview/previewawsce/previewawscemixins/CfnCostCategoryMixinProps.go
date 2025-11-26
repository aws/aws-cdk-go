package previewawscemixins


// Properties for CfnCostCategoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCostCategoryMixinProps := &CfnCostCategoryMixinProps{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Name: jsii.String("name"),
//   	Rules: jsii.String("rules"),
//   	RuleVersion: jsii.String("ruleVersion"),
//   	SplitChargeRules: jsii.String("splitChargeRules"),
//   	Tags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html
//
type CfnCostCategoryMixinProps struct {
	// The default value for the cost category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The unique name of the Cost Category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The array of CostCategoryRule in JSON array format.
	//
	// > Rules are processed in order. If there are multiple rules that match the line item, then the first rule to match is used to determine that Cost Category value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-rules
	//
	Rules *string `field:"optional" json:"rules" yaml:"rules"`
	// The rule schema version in this particular Cost Category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-ruleversion
	//
	RuleVersion *string `field:"optional" json:"ruleVersion" yaml:"ruleVersion"`
	// The split charge rules that are used to allocate your charges between your Cost Category values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-splitchargerules
	//
	SplitChargeRules *string `field:"optional" json:"splitChargeRules" yaml:"splitChargeRules"`
	// Tags to assign to the cost category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html#cfn-ce-costcategory-tags
	//
	Tags *[]*CfnCostCategoryPropsMixin_ResourceTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

