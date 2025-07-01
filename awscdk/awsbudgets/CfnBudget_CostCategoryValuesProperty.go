package awsbudgets


// The cost category values used for filtering the costs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   costCategoryValuesProperty := &CostCategoryValuesProperty{
//   	Key: jsii.String("key"),
//   	MatchOptions: []*string{
//   		jsii.String("matchOptions"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costcategoryvalues.html
//
type CfnBudget_CostCategoryValuesProperty struct {
	// The unique name of the cost category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costcategoryvalues.html#cfn-budgets-budget-costcategoryvalues-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The match options that you can use to filter your results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costcategoryvalues.html#cfn-budgets-budget-costcategoryvalues-matchoptions
	//
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// The specific value of the cost category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costcategoryvalues.html#cfn-budgets-budget-costcategoryvalues-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

