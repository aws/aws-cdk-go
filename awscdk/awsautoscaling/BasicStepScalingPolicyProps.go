package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   workerUtilizationMetric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("MyService"),
//   	MetricName: jsii.String("WorkerUtilization"),
//   })
//
//   autoScalingGroup.scaleOnMetric(jsii.String("ScaleToCPU"), &BasicStepScalingPolicyProps{
//   	Metric: workerUtilizationMetric,
//   	ScalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			Upper: jsii.Number(10),
//   			Change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			Lower: jsii.Number(50),
//   			Change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			Lower: jsii.Number(70),
//   			Change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PERCENT_CHANGE_IN_CAPACITY to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	AdjustmentType: autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   })
//
// Experimental.
type BasicStepScalingPolicyProps struct {
	// Metric to scale on.
	// Experimental.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	// Experimental.
	ScalingSteps *[]*ScalingInterval `field:"required" json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Experimental.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Experimental.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

