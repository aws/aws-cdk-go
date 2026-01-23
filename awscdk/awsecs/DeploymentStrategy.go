package awsecs


// The deployment stratergy to use for ECS controller.
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
type DeploymentStrategy string

const (
	// Rolling update deployment.
	DeploymentStrategy_ROLLING DeploymentStrategy = "ROLLING"
	// Blue/green deployment.
	DeploymentStrategy_BLUE_GREEN DeploymentStrategy = "BLUE_GREEN"
	// Linear deployment with progressive traffic shifting.
	DeploymentStrategy_LINEAR DeploymentStrategy = "LINEAR"
	// Canary deployment with fixed traffic percentage testing.
	DeploymentStrategy_CANARY DeploymentStrategy = "CANARY"
)

