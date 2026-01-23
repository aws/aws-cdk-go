package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for traffic shift during progressive deployments.
//
// Example:
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//   var blueTargetGroup ApplicationTargetGroup
//   var greenTargetGroup ApplicationTargetGroup
//   var prodListenerRule ApplicationListenerRule
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentStrategy: ecs.DeploymentStrategy_LINEAR,
//   	LinearConfiguration: &TrafficShiftConfig{
//   		StepPercent: jsii.Number(10),
//   		StepBakeTime: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   })
//
//   target := service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("web"),
//   	ContainerPort: jsii.Number(80),
//   	AlternateTarget: ecs.NewAlternateTarget(jsii.String("AlternateTarget"), &AlternateTargetProps{
//   		AlternateTargetGroup: greenTargetGroup,
//   		ProductionListener: ecs.ListenerRuleConfiguration_ApplicationListenerRule(prodListenerRule),
//   	}),
//   })
//
//   target.AttachToApplicationTargetGroup(blueTargetGroup)
//
type TrafficShiftConfig struct {
	// The duration to wait between traffic shifting steps.
	//
	// Valid values are 0 to 1440 minutes (24 hours).
	// Default: - Duration.minutes(6) for linear, Duration.minutes(10) for canary
	//
	StepBakeTime awscdk.Duration `field:"optional" json:"stepBakeTime" yaml:"stepBakeTime"`
	// The percentage of production traffic to shift in each step.
	//
	// - For linear deployment: multiples of 0.1 from 3.0 to 100.0
	// - For canary deployment: multiples of 0.1 from 0.1 to 100.0
	// Default: - 10.0 for linear, 5.0 for canary
	//
	StepPercent *float64 `field:"optional" json:"stepPercent" yaml:"stepPercent"`
}

