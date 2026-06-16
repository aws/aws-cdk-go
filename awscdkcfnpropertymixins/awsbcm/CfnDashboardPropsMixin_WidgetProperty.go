package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var costAndUsageExpressionProperty_ CostAndUsageExpressionProperty
//   var expressionProperty_ ExpressionProperty
//   var table interface{}
//
//   widgetProperty := &WidgetProperty{
//   	Configs: []interface{}{
//   		&WidgetConfigProperty{
//   			DisplayConfig: &DisplayConfigProperty{
//   				Graph: map[string]interface{}{
//   					"graphKey": &GraphDisplayConfigProperty{
//   						"visualType": jsii.String("visualType"),
//   					},
//   				},
//   				Table: table,
//   			},
//   			QueryParameters: &QueryParametersProperty{
//   				CostAndUsage: &CostAndUsageQueryProperty{
//   					Filter: &CostAndUsageExpressionProperty{
//   						And: []interface{}{
//   							costAndUsageExpressionProperty_,
//   						},
//   						CostCategories: &CostCategoryValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Dimensions: &DimensionValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Not: costAndUsageExpressionProperty_,
//   						Or: []interface{}{
//   							costAndUsageExpressionProperty_,
//   						},
//   						Tags: &TagValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Granularity: jsii.String("granularity"),
//   					GroupBy: []interface{}{
//   						&GroupDefinitionProperty{
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Metrics: []*string{
//   						jsii.String("metrics"),
//   					},
//   					TimeRange: &DateTimeRangeProperty{
//   						EndTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   						StartTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				ReservationCoverage: &ReservationCoverageQueryProperty{
//   					Filter: &ExpressionProperty{
//   						And: []interface{}{
//   							expressionProperty_,
//   						},
//   						CostCategories: &CostCategoryValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Dimensions: &DimensionValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Not: expressionProperty_,
//   						Tags: &TagValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Granularity: jsii.String("granularity"),
//   					GroupBy: []interface{}{
//   						&GroupDefinitionProperty{
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Metrics: []*string{
//   						jsii.String("metrics"),
//   					},
//   					TimeRange: &DateTimeRangeProperty{
//   						EndTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   						StartTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				ReservationUtilization: &ReservationUtilizationQueryProperty{
//   					Filter: &ExpressionProperty{
//   						And: []interface{}{
//   							expressionProperty_,
//   						},
//   						CostCategories: &CostCategoryValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Dimensions: &DimensionValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Not: expressionProperty_,
//   						Tags: &TagValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Granularity: jsii.String("granularity"),
//   					GroupBy: []interface{}{
//   						&GroupDefinitionProperty{
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					TimeRange: &DateTimeRangeProperty{
//   						EndTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   						StartTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				SavingsPlansCoverage: &SavingsPlansCoverageQueryProperty{
//   					Filter: &ExpressionProperty{
//   						And: []interface{}{
//   							expressionProperty_,
//   						},
//   						CostCategories: &CostCategoryValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Dimensions: &DimensionValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Not: expressionProperty_,
//   						Tags: &TagValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Granularity: jsii.String("granularity"),
//   					GroupBy: []interface{}{
//   						&GroupDefinitionProperty{
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Metrics: []*string{
//   						jsii.String("metrics"),
//   					},
//   					TimeRange: &DateTimeRangeProperty{
//   						EndTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   						StartTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				SavingsPlansUtilization: &SavingsPlansUtilizationQueryProperty{
//   					Filter: &ExpressionProperty{
//   						And: []interface{}{
//   							expressionProperty_,
//   						},
//   						CostCategories: &CostCategoryValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Dimensions: &DimensionValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Not: expressionProperty_,
//   						Tags: &TagValuesProperty{
//   							Key: jsii.String("key"),
//   							MatchOptions: []*string{
//   								jsii.String("matchOptions"),
//   							},
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Granularity: jsii.String("granularity"),
//   					TimeRange: &DateTimeRangeProperty{
//   						EndTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   						StartTime: &DateTimeValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Height: jsii.Number(123),
//   	HorizontalOffset: jsii.Number(123),
//   	Title: jsii.String("title"),
//   	Width: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html
//
type CfnDashboardPropsMixin_WidgetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-configs
	//
	Configs interface{} `field:"optional" json:"configs" yaml:"configs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-height
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-horizontaloffset
	//
	HorizontalOffset *float64 `field:"optional" json:"horizontalOffset" yaml:"horizontalOffset"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widget.html#cfn-bcm-dashboard-widget-width
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

