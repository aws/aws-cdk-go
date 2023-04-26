package awsquicksight


// A `TopBottomFilter` filters values that are at the top or the bottom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topBottomFilterProperty := &TopBottomFilterProperty{
//   	AggregationSortConfigurations: []interface{}{
//   		&AggregationSortConfigurationProperty{
//   			AggregationFunction: &AggregationFunctionProperty{
//   				CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   				DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   				NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   					PercentileAggregation: &PercentileAggregationProperty{
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   				},
//   			},
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			SortDirection: jsii.String("sortDirection"),
//   		},
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//
//   	// the properties below are optional
//   	Limit: jsii.Number(123),
//   	ParameterName: jsii.String("parameterName"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
type CfnAnalysis_TopBottomFilterProperty struct {
	// The aggregation and sort configuration of the top bottom filter.
	AggregationSortConfigurations interface{} `field:"required" json:"aggregationSortConfigurations" yaml:"aggregationSortConfigurations"`
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The number of items to include in the top bottom filter results.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The parameter whose value should be used for the filter value.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

