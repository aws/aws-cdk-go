package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnMetricPropsMixin_MetricFilterStringConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterstringcondition.html#cfn-connect-metric-metricfilterstringcondition-comparison
	//
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilterstringcondition.html#cfn-connect-metric-metricfilterstringcondition-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

