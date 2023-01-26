package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   stepScalingActionProps := &stepScalingActionProps{
//   	autoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	adjustmentType: awscdk.Aws_autoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: cdk.duration.minutes(jsii.Number(30)),
//   	estimatedInstanceWarmup: cdk.*duration.minutes(jsii.Number(30)),
//   	metricAggregationType: awscdk.*Aws_autoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   }
//
type StepScalingActionProps struct {
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// How the adjustment numbers are interpreted.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

