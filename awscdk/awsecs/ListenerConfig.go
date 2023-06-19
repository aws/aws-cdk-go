package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Base class for configuring listener when registering targets.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   service.RegisterLoadBalancerTargets(&EcsTarget{
//   	ContainerName: jsii.String("web"),
//   	ContainerPort: jsii.Number(80),
//   	NewTargetGroupId: jsii.String("ECS"),
//   	Listener: ecs.ListenerConfig_ApplicationListener(listener, &AddApplicationTargetsProps{
//   		Protocol: elbv2.ApplicationProtocol_HTTPS,
//   	}),
//   })
//
// Experimental.
type ListenerConfig interface {
	// Create and attach a target group to listener.
	// Experimental.
	AddTargets(id *string, target *LoadBalancerTargetOptions, service BaseService)
}

// The jsii proxy struct for ListenerConfig
type jsiiProxy_ListenerConfig struct {
	_ byte // padding
}

// Experimental.
func NewListenerConfig_Override(l ListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.ListenerConfig",
		nil, // no parameters
		l,
	)
}

// Create a config for adding target group to ALB listener.
// Experimental.
func ListenerConfig_ApplicationListener(listener awselasticloadbalancingv2.ApplicationListener, props *awselasticloadbalancingv2.AddApplicationTargetsProps) ListenerConfig {
	_init_.Initialize()

	if err := validateListenerConfig_ApplicationListenerParameters(listener, props); err != nil {
		panic(err)
	}
	var returns ListenerConfig

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ListenerConfig",
		"applicationListener",
		[]interface{}{listener, props},
		&returns,
	)

	return returns
}

// Create a config for adding target group to NLB listener.
// Experimental.
func ListenerConfig_NetworkListener(listener awselasticloadbalancingv2.NetworkListener, props *awselasticloadbalancingv2.AddNetworkTargetsProps) ListenerConfig {
	_init_.Initialize()

	if err := validateListenerConfig_NetworkListenerParameters(listener, props); err != nil {
		panic(err)
	}
	var returns ListenerConfig

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ListenerConfig",
		"networkListener",
		[]interface{}{listener, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListenerConfig) AddTargets(id *string, target *LoadBalancerTargetOptions, service BaseService) {
	if err := l.validateAddTargetsParameters(id, target, service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addTargets",
		[]interface{}{id, target, service},
	)
}

