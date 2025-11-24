package mixinsawsapplicationautoscaling


// One or more metric data queries to provide data points for a metric specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveScalingCustomizedScalingMetricProperty := &PredictiveScalingCustomizedScalingMetricProperty{
//   	MetricDataQueries: []interface{}{
//   		&PredictiveScalingMetricDataQueryProperty{
//   			Expression: jsii.String("expression"),
//   			Id: jsii.String("id"),
//   			Label: jsii.String("label"),
//   			MetricStat: &PredictiveScalingMetricStatProperty{
//   				Metric: &PredictiveScalingMetricProperty{
//   					Dimensions: []interface{}{
//   						&PredictiveScalingMetricDimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   				Stat: jsii.String("stat"),
//   				Unit: jsii.String("unit"),
//   			},
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedscalingmetric.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedScalingMetricProperty struct {
	// One or more metric data queries to provide data points for a metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedscalingmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingcustomizedscalingmetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
}

