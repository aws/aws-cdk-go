package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartSortConfigurationProperty := &LineChartSortConfigurationProperty{
//   	CategoryItemsLimitConfiguration: &ItemsLimitConfigurationProperty{
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
//   	ColorItemsLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
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
type CfnDashboard_LineChartSortConfigurationProperty struct {
	// `CfnDashboard.LineChartSortConfigurationProperty.CategoryItemsLimitConfiguration`.
	CategoryItemsLimitConfiguration interface{} `field:"optional" json:"categoryItemsLimitConfiguration" yaml:"categoryItemsLimitConfiguration"`
	// `CfnDashboard.LineChartSortConfigurationProperty.CategorySort`.
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
	// `CfnDashboard.LineChartSortConfigurationProperty.ColorItemsLimitConfiguration`.
	ColorItemsLimitConfiguration interface{} `field:"optional" json:"colorItemsLimitConfiguration" yaml:"colorItemsLimitConfiguration"`
	// `CfnDashboard.LineChartSortConfigurationProperty.SmallMultiplesLimitConfiguration`.
	SmallMultiplesLimitConfiguration interface{} `field:"optional" json:"smallMultiplesLimitConfiguration" yaml:"smallMultiplesLimitConfiguration"`
	// `CfnDashboard.LineChartSortConfigurationProperty.SmallMultiplesSort`.
	SmallMultiplesSort interface{} `field:"optional" json:"smallMultiplesSort" yaml:"smallMultiplesSort"`
}

