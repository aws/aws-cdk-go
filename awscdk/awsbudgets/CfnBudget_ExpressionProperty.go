package awsbudgets


// Use Expression to filter in various Budgets APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var expressionProperty_ ExpressionProperty
//
//   expressionProperty := &ExpressionProperty{
//   	And: []interface{}{
//   		expressionProperty_,
//   	},
//   	CostCategories: &CostCategoryValuesProperty{
//   		Key: jsii.String("key"),
//   		MatchOptions: []*string{
//   			jsii.String("matchOptions"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Dimensions: &ExpressionDimensionValuesProperty{
//   		Key: jsii.String("key"),
//   		MatchOptions: []*string{
//   			jsii.String("matchOptions"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Not: expressionProperty_,
//   	Or: []interface{}{
//   		expressionProperty_,
//   	},
//   	Tags: &TagValuesProperty{
//   		Key: jsii.String("key"),
//   		MatchOptions: []*string{
//   			jsii.String("matchOptions"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html
//
type CfnBudget_ExpressionProperty struct {
	// Return results that match both Dimension objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-and
	//
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// The filter that's based on CostCategoryValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-costcategories
	//
	CostCategories interface{} `field:"optional" json:"costCategories" yaml:"costCategories"`
	// The specific Dimension to use for Expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Return results that don't match a Dimension object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-not
	//
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// Return results that match either Dimension object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-or
	//
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// The specific Tag to use for Expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expression.html#cfn-budgets-budget-expression-tags
	//
	Tags *CfnBudget_TagValuesProperty `field:"optional" json:"tags" yaml:"tags"`
}

