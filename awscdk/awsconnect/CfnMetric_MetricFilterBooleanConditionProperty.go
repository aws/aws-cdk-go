package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricFilterBooleanConditionProperty := &MetricFilterBooleanConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterbooleancondition.html
//
type CfnMetric_MetricFilterBooleanConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterbooleancondition.html#cfn-connect-metric-metricfilterbooleancondition-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
}

