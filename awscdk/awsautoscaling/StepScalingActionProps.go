package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   stepScalingActionProps := &StepScalingActionProps{
//   	AutoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	AdjustmentType: awscdk.Aws_autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	EstimatedInstanceWarmup: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MetricAggregationType: awscdk.*Aws_autoscaling.MetricAggregationType_AVERAGE,
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   }
//
type StepScalingActionProps struct {
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// How the adjustment numbers are interpreted.
	// Default: ChangeInCapacity.
	//
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Period after a scaling completes before another scaling activity can start.
	// Default: The default cooldown configured on the AutoScalingGroup.
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Default: Same as the cooldown.
	//
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	// Default: Average.
	//
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Default: No minimum scaling effect.
	//
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

