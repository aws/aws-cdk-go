package awsquicksight


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
type CfnDashboard_TopBottomFilterProperty struct {
	// `CfnDashboard.TopBottomFilterProperty.AggregationSortConfigurations`.
	AggregationSortConfigurations interface{} `field:"required" json:"aggregationSortConfigurations" yaml:"aggregationSortConfigurations"`
	// `CfnDashboard.TopBottomFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.TopBottomFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnDashboard.TopBottomFilterProperty.Limit`.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// `CfnDashboard.TopBottomFilterProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnDashboard.TopBottomFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

