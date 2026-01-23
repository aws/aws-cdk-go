package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for alternate target groups used in blue/green deployments with load balancers.
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

