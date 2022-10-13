package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// The properties for enabling target tracking scaling based on a custom CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric metric
//
//   trackCustomMetricProps := &trackCustomMetricProps{
//   	metric: metric,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	scaleInCooldown: cdk.duration.minutes(jsii.Number(30)),
//   	scaleOutCooldown: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type TrackCustomMetricProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The custom CloudWatch metric to track.
	//
	// The metric must represent utilization; that is, you will always get the following behavior:
	//
	// - metric > targetValue => scale out
	// - metric < targetValue => scale in.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The target value for the custom CloudWatch metric.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

