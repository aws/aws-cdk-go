package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var costAndUsageExpressionProperty_ CostAndUsageExpressionProperty
//
//   costAndUsageExpressionProperty := &CostAndUsageExpressionProperty{
//   	And: []interface{}{
//   		costAndUsageExpressionProperty_,
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
//   	Not: costAndUsageExpressionProperty_,
//   	Or: []interface{}{
//   		costAndUsageExpressionProperty_,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html
//
type CfnDashboardPropsMixin_CostAndUsageExpressionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-and
	//
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-costcategories
	//
	CostCategories interface{} `field:"optional" json:"costCategories" yaml:"costCategories"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-not
	//
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-or
	//
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costandusageexpression.html#cfn-bcm-dashboard-costandusageexpression-tags
	//
	Tags *CfnDashboardPropsMixin_TagValuesProperty `field:"optional" json:"tags" yaml:"tags"`
}

