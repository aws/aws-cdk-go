package awsautoscaling


// Contains scaling metric information for the `CustomizedScalingMetricSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingCustomizedScalingMetricProperty := &predictiveScalingCustomizedScalingMetricProperty{
//   	metricDataQueries: []interface{}{
//   		&metricDataQueryProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   			label: jsii.String("label"),
//   			metricStat: &metricStatProperty{
//   				metric: &metricProperty{
//   					metricName: jsii.String("metricName"),
//   					namespace: jsii.String("namespace"),
//
//   					// the properties below are optional
//   					dimensions: []interface{}{
//   						&metricDimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				unit: jsii.String("unit"),
//   			},
//   			returnData: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnScalingPolicy_PredictiveScalingCustomizedScalingMetricProperty struct {
	// One or more metric data queries to provide the data points for a scaling metric.
	//
	// Use multiple metric data queries only if you are performing a math expression on returned data.
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

