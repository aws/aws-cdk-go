package awsautoscaling


// A structure that specifies a metric specification for the `MetricSpecifications` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html) property type.
//
// You must specify either a metric pair, or a load metric and a scaling metric individually. Specifying a metric pair instead of individual metrics provides a simpler way to configure metrics for a scaling policy. You choose the metric pair, and the policy automatically knows the correct sum and average statistics to use for the load metric and the scaling metric.
//
// Example
//
// - You create a predictive scaling policy and specify `ALBRequestCount` as the value for the metric pair and `1000.0` as the target value. For this type of metric, you must provide the metric dimension for the corresponding target group, so you also provide a resource label for the Application Load Balancer target group that is attached to your Auto Scaling group.
// - The number of requests the target group receives per minute provides the load metric, and the request count averaged between the members of the target group provides the scaling metric. In CloudWatch, this refers to the `RequestCount` and `RequestCountPerTarget` metrics, respectively.
// - For optimal use of predictive scaling, you adhere to the best practice of using a dynamic scaling policy to automatically scale between the minimum capacity and maximum capacity in response to real-time changes in resource utilization.
// - Amazon EC2 Auto Scaling consumes data points for the load metric over the last 14 days and creates an hourly load forecast for predictive scaling. (A minimum of 24 hours of data is required.)
// - After creating the load forecast, Amazon EC2 Auto Scaling determines when to reduce or increase the capacity of your Auto Scaling group in each hour of the forecast period so that the average number of requests received by each instance is as close to 1000 requests per minute as possible at all times.
//
// For information about using custom metrics with predictive scaling, see [Advanced predictive scaling policy configurations using custom metrics](https://docs.aws.amazon.com/autoscaling/ec2/userguide/predictive-scaling-customized-metric-specification.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingMetricSpecificationProperty := &predictiveScalingMetricSpecificationProperty{
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customizedCapacityMetricSpecification: &predictiveScalingCustomizedCapacityMetricProperty{
//   		metricDataQueries: []interface{}{
//   			&metricDataQueryProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				expression: jsii.String("expression"),
//   				label: jsii.String("label"),
//   				metricStat: &metricStatProperty{
//   					metric: &metricProperty{
//   						metricName: jsii.String("metricName"),
//   						namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&metricDimensionProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					unit: jsii.String("unit"),
//   				},
//   				returnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	customizedLoadMetricSpecification: &predictiveScalingCustomizedLoadMetricProperty{
//   		metricDataQueries: []interface{}{
//   			&metricDataQueryProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				expression: jsii.String("expression"),
//   				label: jsii.String("label"),
//   				metricStat: &metricStatProperty{
//   					metric: &metricProperty{
//   						metricName: jsii.String("metricName"),
//   						namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&metricDimensionProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					unit: jsii.String("unit"),
//   				},
//   				returnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	customizedScalingMetricSpecification: &predictiveScalingCustomizedScalingMetricProperty{
//   		metricDataQueries: []interface{}{
//   			&metricDataQueryProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				expression: jsii.String("expression"),
//   				label: jsii.String("label"),
//   				metricStat: &metricStatProperty{
//   					metric: &metricProperty{
//   						metricName: jsii.String("metricName"),
//   						namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&metricDimensionProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					unit: jsii.String("unit"),
//   				},
//   				returnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	predefinedLoadMetricSpecification: &predictiveScalingPredefinedLoadMetricProperty{
//   		predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   	predefinedMetricPairSpecification: &predictiveScalingPredefinedMetricPairProperty{
//   		predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   	predefinedScalingMetricSpecification: &predictiveScalingPredefinedScalingMetricProperty{
//   		predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   }
//
type CfnScalingPolicy_PredictiveScalingMetricSpecificationProperty struct {
	// Specifies the target utilization.
	//
	// > Some metrics are based on a count instead of a percentage, such as the request count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy specifies one of these metrics, specify the target utilization as the optimal average request or message count per instance during any one-minute interval.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// The customized capacity metric specification.
	CustomizedCapacityMetricSpecification interface{} `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// The customized load metric specification.
	CustomizedLoadMetricSpecification interface{} `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// The customized scaling metric specification.
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// The predefined load metric specification.
	PredefinedLoadMetricSpecification interface{} `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// The predefined metric pair specification from which Amazon EC2 Auto Scaling determines the appropriate scaling metric and load metric to use.
	PredefinedMetricPairSpecification interface{} `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// The predefined scaling metric specification.
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

