package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

// Properties for enabling tracking of an arbitrary metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var metric metric
//
//   metricTargetTrackingProps := &metricTargetTrackingProps{
//   	metric: metric,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	cooldown: duration,
//   	disableScaleIn: jsii.Boolean(false),
//   	estimatedInstanceWarmup: duration,
//   }
//
// Experimental.
type MetricTargetTrackingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	// Experimental.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Experimental.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// Metric to track.
	//
	// The metric must represent a utilization, so that if it's higher than the
	// target value, your ASG should scale out, and if it's lower it should
	// scale in.
	// Experimental.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// Value to keep the metric around.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

