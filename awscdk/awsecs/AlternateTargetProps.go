package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for AlternateTarget configuration.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var lambdaHook function
//   var blueTargetGroup applicationTargetGroup
//   var greenTargetGroup applicationTargetGroup
//   var prodListenerRule applicationListenerRule
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentStrategy: ecs.DeploymentStrategy_BLUE_GREEN,
//   })
//
//   service.AddLifecycleHook(ecs.NewDeploymentLifecycleLambdaTarget(lambdaHook, jsii.String("PreScaleHook"), &DeploymentLifecycleLambdaTargetProps{
//   	LifecycleStages: []deploymentLifecycleStage{
//   		ecs.*deploymentLifecycleStage_PRE_SCALE_UP,
//   	},
//   }))
//
//   target := service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("nginx"),
//   	ContainerPort: jsii.Number(80),
//   	Protocol: ecs.Protocol_TCP,
//   }, ecs.NewAlternateTarget(jsii.String("AlternateTarget"), &AlternateTargetProps{
//   	AlternateTargetGroup: greenTargetGroup,
//   	ProductionListener: ecs.ListenerRuleConfiguration_ApplicationListenerRule(prodListenerRule),
//   }))
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
	AlternateTargetGroup awselasticloadbalancingv2.ITargetGroup `field:"required" json:"alternateTargetGroup" yaml:"alternateTargetGroup"`
	// The production listener rule ARN (ALB) or listener ARN (NLB).
	ProductionListener ListenerRuleConfiguration `field:"required" json:"productionListener" yaml:"productionListener"`
}

