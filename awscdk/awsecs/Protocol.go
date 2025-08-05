package awsecs


// Network protocol.
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
type Protocol string

const (
	// TCP.
	Protocol_TCP Protocol = "TCP"
	// UDP.
	Protocol_UDP Protocol = "UDP"
)

