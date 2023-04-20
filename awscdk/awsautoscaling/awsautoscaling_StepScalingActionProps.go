package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var duration duration
//
//   stepScalingActionProps := &StepScalingActionProps{
//   	AutoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	AdjustmentType: awscdk.Aws_autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   	Cooldown: duration,
//   	EstimatedInstanceWarmup: duration,
//   	MetricAggregationType: awscdk.*Aws_autoscaling.MetricAggregationType_AVERAGE,
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   }
//
// Experimental.
type StepScalingActionProps struct {
	// The auto scaling group.
	// Experimental.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// How the adjustment numbers are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Period after a scaling completes before another scaling activity can start.
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Experimental.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	// Experimental.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

