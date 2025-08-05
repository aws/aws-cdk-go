package awsecs


// Deployment lifecycle stages where hooks can be executed.
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
type DeploymentLifecycleStage string

const (
	// Execute during service reconciliation.
	DeploymentLifecycleStage_RECONCILE_SERVICE DeploymentLifecycleStage = "RECONCILE_SERVICE"
	// Execute before scaling up tasks.
	DeploymentLifecycleStage_PRE_SCALE_UP DeploymentLifecycleStage = "PRE_SCALE_UP"
	// Execute after scaling up tasks.
	DeploymentLifecycleStage_POST_SCALE_UP DeploymentLifecycleStage = "POST_SCALE_UP"
	// Execute during test traffic shift.
	DeploymentLifecycleStage_TEST_TRAFFIC_SHIFT DeploymentLifecycleStage = "TEST_TRAFFIC_SHIFT"
	// Execute after test traffic shift.
	DeploymentLifecycleStage_POST_TEST_TRAFFIC_SHIFT DeploymentLifecycleStage = "POST_TEST_TRAFFIC_SHIFT"
	// Execute during production traffic shift.
	DeploymentLifecycleStage_PRODUCTION_TRAFFIC_SHIFT DeploymentLifecycleStage = "PRODUCTION_TRAFFIC_SHIFT"
	// Execute after production traffic shift.
	DeploymentLifecycleStage_POST_PRODUCTION_TRAFFIC_SHIFT DeploymentLifecycleStage = "POST_PRODUCTION_TRAFFIC_SHIFT"
)

