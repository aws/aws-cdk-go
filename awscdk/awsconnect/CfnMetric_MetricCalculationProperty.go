package awsconnect


// The calculation configuration for the metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricCalculationProperty := &MetricCalculationProperty{
//   	Calculation: jsii.String("calculation"),
//   	CalculationComponents: []interface{}{
//   		&CalculationComponentProperty{
//   			Alias: jsii.String("alias"),
//
//   			// the properties below are optional
//   			MetricFilters: []interface{}{
//   				&MetricFilterProperty{
//   					MetricFilterKey: jsii.String("metricFilterKey"),
//
//   					// the properties below are optional
//   					BooleanCondition: &MetricFilterBooleanConditionProperty{
//   						Comparison: jsii.String("comparison"),
//   					},
//   					Negate: jsii.Boolean(false),
//   					NumberCondition: &MetricFilterNumberConditionProperty{
//   						Comparison: jsii.String("comparison"),
//   						Values: []interface{}{
//   							jsii.Number(123),
//   						},
//   					},
//   					StringCondition: &MetricFilterStringConditionProperty{
//   						Comparison: jsii.String("comparison"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			MetricId: jsii.String("metricId"),
//   			MetricName: jsii.String("metricName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metriccalculation.html
//
type CfnMetric_MetricCalculationProperty struct {
	// The calculation formula.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metriccalculation.html#cfn-connect-metric-metriccalculation-calculation
	//
	Calculation *string `field:"required" json:"calculation" yaml:"calculation"`
	// The calculation components for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-metriccalculation.html#cfn-connect-metric-metriccalculation-calculationcomponents
	//
	CalculationComponents interface{} `field:"required" json:"calculationComponents" yaml:"calculationComponents"`
}

