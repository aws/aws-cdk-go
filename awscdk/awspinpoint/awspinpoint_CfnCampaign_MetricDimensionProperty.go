package awspinpoint


// Specifies metric-based criteria for including or excluding endpoints from a segment.
//
// These criteria derive from custom metrics that you define for endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	value: jsii.Number(123),
//   }
//
type CfnCampaign_MetricDimensionProperty struct {
	// The operator to use when comparing metric values.
	//
	// Valid values are: `GREATER_THAN` , `LESS_THAN` , `GREATER_THAN_OR_EQUAL` , `LESS_THAN_OR_EQUAL` , and `EQUAL` .
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value to compare.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

