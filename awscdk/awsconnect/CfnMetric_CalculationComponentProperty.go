package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculationComponentProperty := &CalculationComponentProperty{
//   	Alias: jsii.String("alias"),
//
//   	// the properties below are optional
//   	MetricFilters: []interface{}{
//   		&MetricFilterProperty{
//   			MetricFilterKey: jsii.String("metricFilterKey"),
//
//   			// the properties below are optional
//   			BooleanCondition: &MetricFilterBooleanConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   			},
//   			Negate: jsii.Boolean(false),
//   			NumberCondition: &MetricFilterNumberConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   			StringCondition: &MetricFilterStringConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	MetricId: jsii.String("metricId"),
//   	MetricName: jsii.String("metricName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-calculationcomponent.html
//
type CfnMetric_CalculationComponentProperty struct {
	// Metric calculation component alias for use within a calculation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-calculationcomponent.html#cfn-connect-metric-calculationcomponent-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-calculationcomponent.html#cfn-connect-metric-calculationcomponent-metricfilters
	//
	MetricFilters interface{} `field:"optional" json:"metricFilters" yaml:"metricFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-calculationcomponent.html#cfn-connect-metric-calculationcomponent-metricid
	//
	MetricId *string `field:"optional" json:"metricId" yaml:"metricId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-calculationcomponent.html#cfn-connect-metric-calculationcomponent-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
}

