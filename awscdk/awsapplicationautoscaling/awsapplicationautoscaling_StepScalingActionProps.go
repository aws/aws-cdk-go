package awsapplicationautoscaling

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
//   var duration duration
//   var scalableTarget scalableTarget
//
//   stepScalingActionProps := &stepScalingActionProps{
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: awscdk.Aws_applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: duration,
//   	metricAggregationType: awscdk.*Aws_applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	policyName: jsii.String("policyName"),
//   }
//
// Experimental.
type StepScalingActionProps struct {
	// The scalable target.
	// Experimental.
	ScalingTarget IScalableTarget `field:"required" json:"scalingTarget" yaml:"scalingTarget"`
	// How the adjustment numbers are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	//
	// For scale out policies, multiple scale outs during the cooldown period are
	// squashed so that only the biggest scale out happens.
	//
	// For scale in policies, subsequent scale ins during the cooldown period are
	// ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	// Experimental.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

