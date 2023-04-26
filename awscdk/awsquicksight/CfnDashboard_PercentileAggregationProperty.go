package awsquicksight


// An aggregation based on the percentile of values in a dimension or measure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   percentileAggregationProperty := &PercentileAggregationProperty{
//   	PercentileValue: jsii.Number(123),
//   }
//
type CfnDashboard_PercentileAggregationProperty struct {
	// The percentile value.
	//
	// This value can be any numeric constant 0–100. A percentile value of 50 computes the median value of the measure.
	PercentileValue *float64 `field:"optional" json:"percentileValue" yaml:"percentileValue"`
}

