package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   barChartSortConfigurationProperty := &BarChartSortConfigurationProperty{
//   	CategoryItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	CategorySort: []interface{}{
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
//   	ColorItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	ColorSort: []interface{}{
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
//   	SmallMultiplesLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	SmallMultiplesSort: []interface{}{
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
type CfnDashboard_BarChartSortConfigurationProperty struct {
	// `CfnDashboard.BarChartSortConfigurationProperty.CategoryItemsLimit`.
	CategoryItemsLimit interface{} `field:"optional" json:"categoryItemsLimit" yaml:"categoryItemsLimit"`
	// `CfnDashboard.BarChartSortConfigurationProperty.CategorySort`.
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
	// `CfnDashboard.BarChartSortConfigurationProperty.ColorItemsLimit`.
	ColorItemsLimit interface{} `field:"optional" json:"colorItemsLimit" yaml:"colorItemsLimit"`
	// `CfnDashboard.BarChartSortConfigurationProperty.ColorSort`.
	ColorSort interface{} `field:"optional" json:"colorSort" yaml:"colorSort"`
	// `CfnDashboard.BarChartSortConfigurationProperty.SmallMultiplesLimitConfiguration`.
	SmallMultiplesLimitConfiguration interface{} `field:"optional" json:"smallMultiplesLimitConfiguration" yaml:"smallMultiplesLimitConfiguration"`
	// `CfnDashboard.BarChartSortConfigurationProperty.SmallMultiplesSort`.
	SmallMultiplesSort interface{} `field:"optional" json:"smallMultiplesSort" yaml:"smallMultiplesSort"`
}

