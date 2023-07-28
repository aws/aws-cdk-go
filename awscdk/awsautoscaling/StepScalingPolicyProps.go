package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var metric metric
//
//   stepScalingPolicyProps := &StepScalingPolicyProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	Metric: metric,
//   	ScalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			Change: jsii.Number(123),
//
//   			// the properties below are optional
//   			Lower: jsii.Number(123),
//   			Upper: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdjustmentType: awscdk.Aws_autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	EstimatedInstanceWarmup: cdk.Duration_*Minutes(jsii.Number(30)),
//   	EvaluationPeriods: jsii.Number(123),
//   	MetricAggregationType: awscdk.*Aws_autoscaling.MetricAggregationType_AVERAGE,
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   }
//
type StepScalingPolicyProps struct {
	// Metric to scale on.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	//
	// Must be between 2 and 40 steps.
	ScalingSteps *[]*ScalingInterval `field:"required" json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

