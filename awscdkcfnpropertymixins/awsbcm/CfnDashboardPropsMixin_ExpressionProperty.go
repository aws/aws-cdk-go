package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
//   	Dimensions: &DimensionValuesProperty{
//   		Key: jsii.String("key"),
//   		MatchOptions: []*string{
//   			jsii.String("matchOptions"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Not: expressionProperty_,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html
//
type CfnDashboardPropsMixin_ExpressionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html#cfn-bcm-dashboard-expression-and
	//
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html#cfn-bcm-dashboard-expression-costcategories
	//
	CostCategories interface{} `field:"optional" json:"costCategories" yaml:"costCategories"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html#cfn-bcm-dashboard-expression-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html#cfn-bcm-dashboard-expression-not
	//
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-expression.html#cfn-bcm-dashboard-expression-tags
	//
	Tags *CfnDashboardPropsMixin_TagValuesProperty `field:"optional" json:"tags" yaml:"tags"`
}

