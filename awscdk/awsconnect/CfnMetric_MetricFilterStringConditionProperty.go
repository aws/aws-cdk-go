package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricFilterStringConditionProperty := &MetricFilterStringConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterstringcondition.html
//
type CfnMetric_MetricFilterStringConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterstringcondition.html#cfn-connect-metric-metricfilterstringcondition-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterstringcondition.html#cfn-connect-metric-metricfilterstringcondition-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

