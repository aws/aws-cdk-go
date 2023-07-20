package awsautoscaling


// Contains capacity metric information for the `CustomizedCapacityMetricSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingCustomizedCapacityMetricProperty := &PredictiveScalingCustomizedCapacityMetricProperty{
//   	MetricDataQueries: []interface{}{
//   		&MetricDataQueryProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			Expression: jsii.String("expression"),
//   			Label: jsii.String("label"),
//   			MetricStat: &MetricStatProperty{
//   				Metric: &MetricProperty{
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//
//   					// the properties below are optional
//   					Dimensions: []interface{}{
//   						&MetricDimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				Unit: jsii.String("unit"),
//   			},
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.html
//
type CfnScalingPolicy_PredictiveScalingCustomizedCapacityMetricProperty struct {
	// One or more metric data queries to provide the data points for a capacity metric.
	//
	// Use multiple metric data queries only if you are performing a math expression on returned data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.html#cfn-autoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

