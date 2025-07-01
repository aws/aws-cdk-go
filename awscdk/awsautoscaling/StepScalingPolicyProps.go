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
//   	DatapointsToAlarm: jsii.Number(123),
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
	// Default: ChangeInCapacity.
	//
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	// Default: Default cooldown period on your AutoScalingGroup.
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The number of data points out of the evaluation periods that must be breaching to trigger a scaling action.
	//
	// Creates an "M out of N" alarm, where this property is the M and the value set for
	// `evaluationPeriods` is the N value.
	//
	// Only has meaning if `evaluationPeriods != 1`. Must be less than or equal to
	// `evaluationPeriods`.
	// Default: - Same as `evaluationPeriods`.
	//
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Default: Same as the cooldown.
	//
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	//
	// If `datapointsToAlarm` is not set, then all data points in the evaluation period
	// must meet the criteria to trigger a scaling action.
	// Default: 1.
	//
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Default: - The statistic from the metric if applicable (MIN, MAX, AVERAGE), otherwise AVERAGE.
	//
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Default: No minimum scaling effect.
	//
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

