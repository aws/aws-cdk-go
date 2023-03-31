package awsautoscaling


// `TargetTrackingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html) resource that specifies a target tracking scaling policy configuration for Amazon EC2 Auto Scaling.
//
// For more information about scaling policies, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingConfigurationProperty := &targetTrackingConfigurationProperty{
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customizedMetricSpecification: &customizedMetricSpecificationProperty{
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		unit: jsii.String("unit"),
//   	},
//   	disableScaleIn: jsii.Boolean(false),
//   	predefinedMetricSpecification: &predefinedMetricSpecificationProperty{
//   		predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   }
//
type CfnScalingPolicy_TargetTrackingConfigurationProperty struct {
	// The target value for the metric.
	//
	// > Some metrics are based on a count instead of a percentage, such as the request count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy specifies one of these metrics, specify the target utilization as the optimal average request or message count per instance during any one-minute interval.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You must specify either a predefined metric or a customized metric.
	CustomizedMetricSpecification interface{} `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Indicates whether scaling in by the target tracking scaling policy is disabled.
	//
	// If scaling in is disabled, the target tracking scaling policy doesn't remove instances from the Auto Scaling group. Otherwise, the target tracking scaling policy can remove instances from the Auto Scaling group. The default is `false` .
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A predefined metric.
	//
	// You must specify either a predefined metric or a customized metric.
	PredefinedMetricSpecification interface{} `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
}

