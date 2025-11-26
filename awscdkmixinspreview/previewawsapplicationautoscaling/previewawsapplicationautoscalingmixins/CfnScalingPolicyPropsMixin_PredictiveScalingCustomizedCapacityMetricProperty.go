package previewawsapplicationautoscalingmixins


// Represents a CloudWatch metric of your choosing for a predictive scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveScalingCustomizedCapacityMetricProperty := &PredictiveScalingCustomizedCapacityMetricProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedCapacityMetricProperty struct {
	// One or more metric data queries to provide data points for a metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
}

