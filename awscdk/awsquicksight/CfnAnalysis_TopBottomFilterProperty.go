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
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			SortDirection: jsii.String("sortDirection"),
//
//   			// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html
//
type CfnAnalysis_TopBottomFilterProperty struct {
	// The aggregation and sort configuration of the top bottom filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-aggregationsortconfigurations
	//
	AggregationSortConfigurations interface{} `field:"required" json:"aggregationSortConfigurations" yaml:"aggregationSortConfigurations"`
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The number of items to include in the top bottom filter results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-limit
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The parameter whose value should be used for the filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

