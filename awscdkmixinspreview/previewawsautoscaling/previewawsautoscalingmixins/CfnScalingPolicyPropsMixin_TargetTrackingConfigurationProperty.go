package previewawsautoscalingmixins


// `TargetTrackingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html) resource that specifies a target tracking scaling policy configuration for Amazon EC2 Auto Scaling.
//
// For more information about scaling policies, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetTrackingConfigurationProperty := &TargetTrackingConfigurationProperty{
//   	CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Metrics: []interface{}{
//   			&TargetTrackingMetricDataQueryProperty{
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &TargetTrackingMetricStatProperty{
//   					Metric: &MetricProperty{
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Period: jsii.Number(123),
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				Period: jsii.Number(123),
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   		Namespace: jsii.String("namespace"),
//   		Period: jsii.Number(123),
//   		Statistic: jsii.String("statistic"),
//   		Unit: jsii.String("unit"),
//   	},
//   	DisableScaleIn: jsii.Boolean(false),
//   	PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	TargetValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html
//
type CfnScalingPolicyPropsMixin_TargetTrackingConfigurationProperty struct {
	// A customized metric.
	//
	// You must specify either a predefined metric or a customized metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-customizedmetricspecification
	//
	CustomizedMetricSpecification interface{} `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Indicates whether scaling in by the target tracking scaling policy is disabled.
	//
	// If scaling in is disabled, the target tracking scaling policy doesn't remove instances from the Auto Scaling group. Otherwise, the target tracking scaling policy can remove instances from the Auto Scaling group. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-disablescalein
	//
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A predefined metric.
	//
	// You must specify either a predefined metric or a customized metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-predefinedmetricspecification
	//
	PredefinedMetricSpecification interface{} `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// The target value for the metric.
	//
	// > Some metrics are based on a count instead of a percentage, such as the request count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy specifies one of these metrics, specify the target utilization as the optimal average request or message count per instance during any one-minute interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-targetvalue
	//
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

