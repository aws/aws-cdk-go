package awsapplicationautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric metric
//   var scalableTarget scalableTarget
//
//   stepScalingPolicyProps := &stepScalingPolicyProps{
//   	metric: metric,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			change: jsii.Number(123),
//
//   			// the properties below are optional
//   			lower: jsii.Number(123),
//   			upper: jsii.Number(123),
//   		},
//   	},
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: awscdk.Aws_applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: cdk.duration.minutes(jsii.Number(30)),
//   	datapointsToAlarm: jsii.Number(123),
//   	evaluationPeriods: jsii.Number(123),
//   	metricAggregationType: awscdk.*Aws_applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   }
//
type StepScalingPolicyProps struct {
	// Metric to scale on.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	ScalingSteps *[]*ScalingInterval `field:"required" json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	//
	// Subsequent scale outs during the cooldown period are squashed so that only
	// the biggest scale out happens.
	//
	// Subsequent scale ins during the cooldown period are ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The number of data points out of the evaluation periods that must be breaching to trigger a scaling action.
	//
	// Creates an "M out of N" alarm, where this property is the M and the value set for
	// `evaluationPeriods` is the N value.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	//
	// If `datapointsToAlarm` is not set, then all data points in the evaluation period
	// must meet the criteria to trigger a scaling action.
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
	// The scaling target.
	ScalingTarget IScalableTarget `field:"required" json:"scalingTarget" yaml:"scalingTarget"`
}

