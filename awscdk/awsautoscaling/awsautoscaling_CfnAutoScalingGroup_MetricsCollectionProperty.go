package awsautoscaling


// `MetricsCollection` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource that describes the group metrics that an Amazon EC2 Auto Scaling group sends to Amazon CloudWatch. These metrics describe the group rather than any of its instances.
//
// For more information, see [Monitor CloudWatch metrics for your Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsCollectionProperty := &metricsCollectionProperty{
//   	granularity: jsii.String("granularity"),
//
//   	// the properties below are optional
//   	metrics: []*string{
//   		jsii.String("metrics"),
//   	},
//   }
//
type CfnAutoScalingGroup_MetricsCollectionProperty struct {
	// The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch.
	//
	// The only valid value is `1Minute` .
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// Specifies which group-level metrics to start collecting. You can specify one or more of the following metrics:.
	//
	// - `GroupMinSize`
	// - `GroupMaxSize`
	// - `GroupDesiredCapacity`
	// - `GroupInServiceInstances`
	// - `GroupPendingInstances`
	// - `GroupStandbyInstances`
	// - `GroupTerminatingInstances`
	// - `GroupTotalInstances`
	//
	// The instance weighting feature supports the following additional metrics:
	//
	// - `GroupInServiceCapacity`
	// - `GroupPendingCapacity`
	// - `GroupStandbyCapacity`
	// - `GroupTerminatingCapacity`
	// - `GroupTotalCapacity`
	//
	// The warm pools feature supports the following additional metrics:
	//
	// - `WarmPoolDesiredCapacity`
	// - `WarmPoolWarmedCapacity`
	// - `WarmPoolPendingCapacity`
	// - `WarmPoolTerminatingCapacity`
	// - `WarmPoolTotalCapacity`
	// - `GroupAndWarmPoolDesiredCapacity`
	// - `GroupAndWarmPoolTotalCapacity`
	//
	// If you specify `Granularity` and don't specify any metrics, all metrics are enabled.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

