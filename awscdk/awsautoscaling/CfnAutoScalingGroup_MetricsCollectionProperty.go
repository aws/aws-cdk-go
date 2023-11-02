package awsautoscaling


// `MetricsCollection` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource that describes the group metrics that an Amazon EC2 Auto Scaling group sends to Amazon CloudWatch. These metrics describe the group rather than any of its instances.
//
// For more information, see [Monitor CloudWatch metrics for your Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the `AWS::AutoScaling::AutoScalingGroup` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsCollectionProperty := &MetricsCollectionProperty{
//   	Granularity: jsii.String("granularity"),
//
//   	// the properties below are optional
//   	Metrics: []*string{
//   		jsii.String("metrics"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-metricscollection.html
//
type CfnAutoScalingGroup_MetricsCollectionProperty struct {
	// The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch.
	//
	// The only valid value is `1Minute` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-metricscollection.html#cfn-autoscaling-autoscalinggroup-metricscollection-granularity
	//
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// Identifies the metrics to enable.
	//
	// You can specify one or more of the following metrics:
	//
	// - `GroupMinSize`
	// - `GroupMaxSize`
	// - `GroupDesiredCapacity`
	// - `GroupInServiceInstances`
	// - `GroupPendingInstances`
	// - `GroupStandbyInstances`
	// - `GroupTerminatingInstances`
	// - `GroupTotalInstances`
	// - `GroupInServiceCapacity`
	// - `GroupPendingCapacity`
	// - `GroupStandbyCapacity`
	// - `GroupTerminatingCapacity`
	// - `GroupTotalCapacity`
	// - `WarmPoolDesiredCapacity`
	// - `WarmPoolWarmedCapacity`
	// - `WarmPoolPendingCapacity`
	// - `WarmPoolTerminatingCapacity`
	// - `WarmPoolTotalCapacity`
	// - `GroupAndWarmPoolDesiredCapacity`
	// - `GroupAndWarmPoolTotalCapacity`
	//
	// If you specify `Granularity` and don't specify any metrics, all metrics are enabled.
	//
	// For more information, see [Auto Scaling group metrics](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-cloudwatch-monitoring.html#as-group-metrics) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-metricscollection.html#cfn-autoscaling-autoscalinggroup-metricscollection-metrics
	//
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

