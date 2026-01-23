package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Properties for AlternateTarget configuration.
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
type AlternateTargetProps struct {
	// The IAM role for the configuration.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The test listener configuration.
	// Default: - none.
	//
	TestListener ListenerRuleConfiguration `field:"optional" json:"testListener" yaml:"testListener"`
	// The alternate target group.
	AlternateTargetGroup interfacesawselasticloadbalancingv2.ITargetGroupRef `field:"required" json:"alternateTargetGroup" yaml:"alternateTargetGroup"`
	// The production listener rule ARN (ALB) or listener ARN (NLB).
	ProductionListener ListenerRuleConfiguration `field:"required" json:"productionListener" yaml:"productionListener"`
}

