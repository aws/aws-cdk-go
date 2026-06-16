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
//   widgetConfigProperty := &WidgetConfigProperty{
//   	DisplayConfig: &DisplayConfigProperty{
//   		Graph: map[string]interface{}{
//   			"graphKey": &GraphDisplayConfigProperty{
//   				"visualType": jsii.String("visualType"),
//   			},
//   		},
//   		Table: table,
//   	},
//   	QueryParameters: &QueryParametersProperty{
//   		CostAndUsage: &CostAndUsageQueryProperty{
//   			Filter: &CostAndUsageExpressionProperty{
//   				And: []interface{}{
//   					costAndUsageExpressionProperty_,
//   				},
//   				CostCategories: &CostCategoryValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Dimensions: &DimensionValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Not: costAndUsageExpressionProperty_,
//   				Or: []interface{}{
//   					costAndUsageExpressionProperty_,
//   				},
//   				Tags: &TagValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Granularity: jsii.String("granularity"),
//   			GroupBy: []interface{}{
//   				&GroupDefinitionProperty{
//   					Key: jsii.String("key"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   			TimeRange: &DateTimeRangeProperty{
//   				EndTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   				StartTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		ReservationCoverage: &ReservationCoverageQueryProperty{
//   			Filter: &ExpressionProperty{
//   				And: []interface{}{
//   					expressionProperty_,
//   				},
//   				CostCategories: &CostCategoryValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Dimensions: &DimensionValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Not: expressionProperty_,
//   				Tags: &TagValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Granularity: jsii.String("granularity"),
//   			GroupBy: []interface{}{
//   				&GroupDefinitionProperty{
//   					Key: jsii.String("key"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   			TimeRange: &DateTimeRangeProperty{
//   				EndTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   				StartTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		ReservationUtilization: &ReservationUtilizationQueryProperty{
//   			Filter: &ExpressionProperty{
//   				And: []interface{}{
//   					expressionProperty_,
//   				},
//   				CostCategories: &CostCategoryValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Dimensions: &DimensionValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Not: expressionProperty_,
//   				Tags: &TagValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Granularity: jsii.String("granularity"),
//   			GroupBy: []interface{}{
//   				&GroupDefinitionProperty{
//   					Key: jsii.String("key"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			TimeRange: &DateTimeRangeProperty{
//   				EndTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   				StartTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SavingsPlansCoverage: &SavingsPlansCoverageQueryProperty{
//   			Filter: &ExpressionProperty{
//   				And: []interface{}{
//   					expressionProperty_,
//   				},
//   				CostCategories: &CostCategoryValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Dimensions: &DimensionValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Not: expressionProperty_,
//   				Tags: &TagValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Granularity: jsii.String("granularity"),
//   			GroupBy: []interface{}{
//   				&GroupDefinitionProperty{
//   					Key: jsii.String("key"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   			TimeRange: &DateTimeRangeProperty{
//   				EndTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   				StartTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SavingsPlansUtilization: &SavingsPlansUtilizationQueryProperty{
//   			Filter: &ExpressionProperty{
//   				And: []interface{}{
//   					expressionProperty_,
//   				},
//   				CostCategories: &CostCategoryValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Dimensions: &DimensionValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Not: expressionProperty_,
//   				Tags: &TagValuesProperty{
//   					Key: jsii.String("key"),
//   					MatchOptions: []*string{
//   						jsii.String("matchOptions"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Granularity: jsii.String("granularity"),
//   			TimeRange: &DateTimeRangeProperty{
//   				EndTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   				StartTime: &DateTimeValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widgetconfig.html
//
type CfnDashboardPropsMixin_WidgetConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widgetconfig.html#cfn-bcm-dashboard-widgetconfig-displayconfig
	//
	DisplayConfig interface{} `field:"optional" json:"displayConfig" yaml:"displayConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-widgetconfig.html#cfn-bcm-dashboard-widgetconfig-queryparameters
	//
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}

