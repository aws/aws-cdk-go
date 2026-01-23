package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Represents a listener configuration for advanced load balancer settings.
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
type ListenerRuleConfiguration interface {
}

// The jsii proxy struct for ListenerRuleConfiguration
type jsiiProxy_ListenerRuleConfiguration struct {
	_ byte // padding
}

func NewListenerRuleConfiguration_Override(l ListenerRuleConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ListenerRuleConfiguration",
		nil, // no parameters
		l,
	)
}

// Use an Application Load Balancer listener rule.
func ListenerRuleConfiguration_ApplicationListenerRule(rule awselasticloadbalancingv2.ApplicationListenerRule) ListenerRuleConfiguration {
	_init_.Initialize()

	if err := validateListenerRuleConfiguration_ApplicationListenerRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns ListenerRuleConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ListenerRuleConfiguration",
		"applicationListenerRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Use a Network Load Balancer listener.
func ListenerRuleConfiguration_NetworkListener(listener awselasticloadbalancingv2.NetworkListener) ListenerRuleConfiguration {
	_init_.Initialize()

	if err := validateListenerRuleConfiguration_NetworkListenerParameters(listener); err != nil {
		panic(err)
	}
	var returns ListenerRuleConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ListenerRuleConfiguration",
		"networkListener",
		[]interface{}{listener},
		&returns,
	)

	return returns
}

