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
//   predictiveScalingMetricSpecificationProperty := &PredictiveScalingMetricSpecificationProperty{
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				Expression: jsii.String("expression"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				Expression: jsii.String("expression"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				Expression: jsii.String("expression"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html
//
type CfnScalingPolicy_PredictiveScalingMetricSpecificationProperty struct {
	// Specifies the target utilization.
	//
	// > Some metrics are based on a count instead of a percentage, such as the request count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy specifies one of these metrics, specify the target utilization as the optimal average request or message count per instance during any one-minute interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// The customized capacity metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-customizedcapacitymetricspecification
	//
	CustomizedCapacityMetricSpecification interface{} `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// The customized load metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-customizedloadmetricspecification
	//
	CustomizedLoadMetricSpecification interface{} `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// The customized scaling metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-customizedscalingmetricspecification
	//
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// The predefined load metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedloadmetricspecification
	//
	PredefinedLoadMetricSpecification interface{} `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// The predefined metric pair specification from which Amazon EC2 Auto Scaling determines the appropriate scaling metric and load metric to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedmetricpairspecification
	//
	PredefinedMetricPairSpecification interface{} `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// The predefined scaling metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-autoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedscalingmetricspecification
	//
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

