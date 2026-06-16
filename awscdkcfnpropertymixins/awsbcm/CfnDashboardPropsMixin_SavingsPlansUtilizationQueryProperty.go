package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var expressionProperty_ ExpressionProperty
//
//   savingsPlansUtilizationQueryProperty := &SavingsPlansUtilizationQueryProperty{
//   	Filter: &ExpressionProperty{
//   		And: []interface{}{
//   			expressionProperty_,
//   		},
//   		CostCategories: &CostCategoryValuesProperty{
//   			Key: jsii.String("key"),
//   			MatchOptions: []*string{
//   				jsii.String("matchOptions"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Dimensions: &DimensionValuesProperty{
//   			Key: jsii.String("key"),
//   			MatchOptions: []*string{
//   				jsii.String("matchOptions"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Not: expressionProperty_,
//   		Tags: &TagValuesProperty{
//   			Key: jsii.String("key"),
//   			MatchOptions: []*string{
//   				jsii.String("matchOptions"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Granularity: jsii.String("granularity"),
//   	TimeRange: &DateTimeRangeProperty{
//   		EndTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   		StartTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-savingsplansutilizationquery.html
//
type CfnDashboardPropsMixin_SavingsPlansUtilizationQueryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-savingsplansutilizationquery.html#cfn-bcm-dashboard-savingsplansutilizationquery-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-savingsplansutilizationquery.html#cfn-bcm-dashboard-savingsplansutilizationquery-granularity
	//
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-savingsplansutilizationquery.html#cfn-bcm-dashboard-savingsplansutilizationquery-timerange
	//
	TimeRange interface{} `field:"optional" json:"timeRange" yaml:"timeRange"`
}

