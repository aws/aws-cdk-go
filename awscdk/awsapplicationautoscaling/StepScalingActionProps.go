package awsapplicationautoscaling

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
//   var scalableTarget scalableTarget
//
//   stepScalingActionProps := &StepScalingActionProps{
//   	ScalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	AdjustmentType: awscdk.Aws_applicationautoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	MetricAggregationType: awscdk.*Aws_applicationautoscaling.MetricAggregationType_AVERAGE,
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   	PolicyName: jsii.String("policyName"),
//   }
//
type StepScalingActionProps struct {
	// The scalable target.
	ScalingTarget IScalableTarget `field:"required" json:"scalingTarget" yaml:"scalingTarget"`
	// How the adjustment numbers are interpreted.
	// Default: ChangeInCapacity.
	//
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
	// Default: No cooldown period.
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
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
	// A name for the scaling policy.
	// Default: Automatically generated name.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

