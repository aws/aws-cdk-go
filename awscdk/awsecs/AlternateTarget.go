package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for alternate target groups used in blue/green deployments with load balancers.
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
type AlternateTarget interface {
	IAlternateTarget
	// Bind this configuration to a service.
	Bind(scope constructs.IConstruct) *AlternateTargetConfig
}

// The jsii proxy struct for AlternateTarget
type jsiiProxy_AlternateTarget struct {
	jsiiProxy_IAlternateTarget
}

func NewAlternateTarget(id *string, props *AlternateTargetProps) AlternateTarget {
	_init_.Initialize()

	if err := validateNewAlternateTargetParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlternateTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AlternateTarget",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewAlternateTarget_Override(a AlternateTarget, id *string, props *AlternateTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AlternateTarget",
		[]interface{}{id, props},
		a,
	)
}

func (a *jsiiProxy_AlternateTarget) Bind(scope constructs.IConstruct) *AlternateTargetConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AlternateTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

