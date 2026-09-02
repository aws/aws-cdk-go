package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricFilterNumberConditionProperty := &MetricFilterNumberConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilternumbercondition.html
//
type CfnMetric_MetricFilterNumberConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilternumbercondition.html#cfn-connect-metric-metricfilternumbercondition-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilternumbercondition.html#cfn-connect-metric-metricfilternumbercondition-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

