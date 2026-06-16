package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var expressionProperty_ ExpressionProperty
//
//   reservationCoverageQueryProperty := &ReservationCoverageQueryProperty{
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
//   	GroupBy: []interface{}{
//   		&GroupDefinitionProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Metrics: []*string{
//   		jsii.String("metrics"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html
//
type CfnDashboardPropsMixin_ReservationCoverageQueryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html#cfn-bcm-dashboard-reservationcoveragequery-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html#cfn-bcm-dashboard-reservationcoveragequery-granularity
	//
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html#cfn-bcm-dashboard-reservationcoveragequery-groupby
	//
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html#cfn-bcm-dashboard-reservationcoveragequery-metrics
	//
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-reservationcoveragequery.html#cfn-bcm-dashboard-reservationcoveragequery-timerange
	//
	TimeRange interface{} `field:"optional" json:"timeRange" yaml:"timeRange"`
}

