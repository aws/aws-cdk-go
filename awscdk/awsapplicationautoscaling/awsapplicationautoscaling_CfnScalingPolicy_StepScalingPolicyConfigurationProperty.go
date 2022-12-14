package awsapplicationautoscaling


// `StepScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a step scaling policy configuration for Application Auto Scaling.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScalingPolicy.html) in the *Application Auto Scaling API Reference* . For more information about step scaling policies, see [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepScalingPolicyConfigurationProperty := &stepScalingPolicyConfigurationProperty{
//   	adjustmentType: jsii.String("adjustmentType"),
//   	cooldown: jsii.Number(123),
//   	metricAggregationType: jsii.String("metricAggregationType"),
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	stepAdjustments: []interface{}{
//   		&stepAdjustmentProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			metricIntervalLowerBound: jsii.Number(123),
//   			metricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnScalingPolicy_StepScalingPolicyConfigurationProperty struct {
	// Specifies whether the `ScalingAdjustment` value in the `StepAdjustment` property is an absolute number or a percentage of the current capacity.
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, to wait for a previous scaling activity to take effect.
	//
	// With scale-out policies, the intention is to continuously (but not excessively) scale out. After Application Auto Scaling successfully scales out using a step scaling policy, it starts to calculate the cooldown time. The scaling policy won't increase the desired capacity again unless either a larger scale out is triggered or the cooldown period ends. While the cooldown period is in effect, capacity added by the initiating scale-out activity is calculated as part of the desired capacity for the next scale-out activity. For example, when an alarm triggers a step scaling policy to increase the capacity by 2, the scaling activity completes successfully, and a cooldown period starts. If the alarm triggers again during the cooldown period but at a more aggressive step adjustment of 3, the previous increase of 2 is considered part of the current capacity. Therefore, only 1 is added to the capacity.
	//
	// With scale-in policies, the intention is to scale in conservatively to protect your application’s availability, so scale-in activities are blocked until the cooldown period has expired. However, if another alarm triggers a scale-out activity during the cooldown period after a scale-in activity, Application Auto Scaling scales out the target immediately. In this case, the cooldown period for the scale-in activity stops and doesn't complete.
	//
	// Application Auto Scaling provides a default value of 600 for Amazon ElastiCache replication groups and a default value of 300 for the following scalable targets:
	//
	// - AppStream 2.0 fleets
	// - Aurora DB clusters
	// - ECS services
	// - EMR clusters
	// - Neptune clusters
	// - SageMaker endpoint variants
	// - Spot Fleets
	// - Custom resources
	//
	// For all other scalable targets, the default value is 0:
	//
	// - Amazon Comprehend document classification and entity recognizer endpoints
	// - DynamoDB tables and global secondary indexes
	// - Amazon Keyspaces tables
	// - Lambda provisioned concurrency
	// - Amazon MSK broker storage.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	//
	// Valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Amazon ECS service by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the service has 4 tasks and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Application Auto Scaling scales out the service by 2 tasks.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// At least one step adjustment is required if you are adding a new step scaling policy configuration.
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
}

