package awsquicksight


// The sort configuration of a heat map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   heatMapSortConfigurationProperty := &HeatMapSortConfigurationProperty{
//   	HeatMapColumnItemsLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	HeatMapColumnSort: []interface{}{
//   		&FieldSortOptionsProperty{
//   			ColumnSort: &ColumnSortProperty{
//   				Direction: jsii.String("direction"),
//   				SortBy: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   			},
//   			FieldSort: &FieldSortProperty{
//   				Direction: jsii.String("direction"),
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   	},
//   	HeatMapRowItemsLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	HeatMapRowSort: []interface{}{
//   		&FieldSortOptionsProperty{
//   			ColumnSort: &ColumnSortProperty{
//   				Direction: jsii.String("direction"),
//   				SortBy: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   			},
//   			FieldSort: &FieldSortProperty{
//   				Direction: jsii.String("direction"),
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapsortconfiguration.html
//
type CfnDashboard_HeatMapSortConfigurationProperty struct {
	// The limit on the number of columns that are displayed in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapsortconfiguration.html#cfn-quicksight-dashboard-heatmapsortconfiguration-heatmapcolumnitemslimitconfiguration
	//
	HeatMapColumnItemsLimitConfiguration interface{} `field:"optional" json:"heatMapColumnItemsLimitConfiguration" yaml:"heatMapColumnItemsLimitConfiguration"`
	// The column sort configuration for heat map for columns that aren't a part of a field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapsortconfiguration.html#cfn-quicksight-dashboard-heatmapsortconfiguration-heatmapcolumnsort
	//
	HeatMapColumnSort interface{} `field:"optional" json:"heatMapColumnSort" yaml:"heatMapColumnSort"`
	// The limit on the number of rows that are displayed in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapsortconfiguration.html#cfn-quicksight-dashboard-heatmapsortconfiguration-heatmaprowitemslimitconfiguration
	//
	HeatMapRowItemsLimitConfiguration interface{} `field:"optional" json:"heatMapRowItemsLimitConfiguration" yaml:"heatMapRowItemsLimitConfiguration"`
	// The field sort configuration of the rows fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapsortconfiguration.html#cfn-quicksight-dashboard-heatmapsortconfiguration-heatmaprowsort
	//
	HeatMapRowSort interface{} `field:"optional" json:"heatMapRowSort" yaml:"heatMapRowSort"`
}

