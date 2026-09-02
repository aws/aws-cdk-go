package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricFilterProperty := &MetricFilterProperty{
//   	MetricFilterKey: jsii.String("metricFilterKey"),
//
//   	// the properties below are optional
//   	BooleanCondition: &MetricFilterBooleanConditionProperty{
//   		Comparison: jsii.String("comparison"),
//   	},
//   	Negate: jsii.Boolean(false),
//   	NumberCondition: &MetricFilterNumberConditionProperty{
//   		Comparison: jsii.String("comparison"),
//   		Values: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	StringCondition: &MetricFilterStringConditionProperty{
//   		Comparison: jsii.String("comparison"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html
//
type CfnMetric_MetricFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html#cfn-connect-metric-metricfilter-metricfilterkey
	//
	MetricFilterKey *string `field:"required" json:"metricFilterKey" yaml:"metricFilterKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html#cfn-connect-metric-metricfilter-booleancondition
	//
	BooleanCondition interface{} `field:"optional" json:"booleanCondition" yaml:"booleanCondition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html#cfn-connect-metric-metricfilter-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html#cfn-connect-metric-metricfilter-numbercondition
	//
	NumberCondition interface{} `field:"optional" json:"numberCondition" yaml:"numberCondition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metricfilter.html#cfn-connect-metric-metricfilter-stringcondition
	//
	StringCondition interface{} `field:"optional" json:"stringCondition" yaml:"stringCondition"`
}

