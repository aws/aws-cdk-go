package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Properties for AlternateTarget configuration.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//   var lambdaHook Function
//   var blueTargetGroup ApplicationTargetGroup
//   var greenTargetGroup ApplicationTargetGroup
//   var prodListenerRule ApplicationListenerRule
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentStrategy: ecs.DeploymentStrategy_BLUE_GREEN,
//   })
//
//   service.AddLifecycleHook(ecs.NewDeploymentLifecycleLambdaTarget(lambdaHook, jsii.String("PreScaleHook"), &DeploymentLifecycleLambdaTargetProps{
//   	LifecycleStages: []DeploymentLifecycleStage{
//   		ecs.DeploymentLifecycleStage_PRE_SCALE_UP,
//   	},
//   }))
//
//   target := service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("nginx"),
//   	ContainerPort: jsii.Number(80),
//   	Protocol: ecs.Protocol_TCP,
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

