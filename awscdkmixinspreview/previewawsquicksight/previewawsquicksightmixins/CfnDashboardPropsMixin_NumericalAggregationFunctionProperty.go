package previewawsquicksightmixins


// Aggregation for numerical values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   numericalAggregationFunctionProperty := &NumericalAggregationFunctionProperty{
//   	PercentileAggregation: &PercentileAggregationProperty{
//   		PercentileValue: jsii.Number(123),
//   	},
//   	SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericalaggregationfunction.html
//
type CfnDashboardPropsMixin_NumericalAggregationFunctionProperty struct {
	// An aggregation based on the percentile of values in a dimension or measure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericalaggregationfunction.html#cfn-quicksight-dashboard-numericalaggregationfunction-percentileaggregation
	//
	PercentileAggregation interface{} `field:"optional" json:"percentileAggregation" yaml:"percentileAggregation"`
	// Built-in aggregation functions for numerical values.
	//
	// - `SUM` : The sum of a dimension or measure.
	// - `AVERAGE` : The average of a dimension or measure.
	// - `MIN` : The minimum value of a dimension or measure.
	// - `MAX` : The maximum value of a dimension or measure.
	// - `COUNT` : The count of a dimension or measure.
	// - `DISTINCT_COUNT` : The count of distinct values in a dimension or measure.
	// - `VAR` : The variance of a dimension or measure.
	// - `VARP` : The partitioned variance of a dimension or measure.
	// - `STDEV` : The standard deviation of a dimension or measure.
	// - `STDEVP` : The partitioned standard deviation of a dimension or measure.
	// - `MEDIAN` : The median value of a dimension or measure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericalaggregationfunction.html#cfn-quicksight-dashboard-numericalaggregationfunction-simplenumericalaggregation
	//
	SimpleNumericalAggregation *string `field:"optional" json:"simpleNumericalAggregation" yaml:"simpleNumericalAggregation"`
}

