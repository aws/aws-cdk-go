package awsapplicationautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingCustomizedLoadMetricProperty := &PredictiveScalingCustomizedLoadMetricProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedloadmetric.html
//
type CfnScalingPolicy_PredictiveScalingCustomizedLoadMetricProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedloadmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingcustomizedloadmetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

