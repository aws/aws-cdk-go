package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// A single Application Load Balancer's listener as the target for load balancing.
//
// Example:
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import patterns "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//
//
//   task := ecs.NewFargateTaskDefinition(this, jsii.String("Task"), &FargateTaskDefinitionProps{
//   	Cpu: jsii.Number(256),
//   	MemoryLimitMiB: jsii.Number(512),
//   })
//   task.AddContainer(jsii.String("nginx"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest")),
//   	PortMappings: []PortMapping{
//   		&PortMapping{
//   			ContainerPort: jsii.Number(80),
//   		},
//   	},
//   })
//
//   svc := patterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Vpc: Vpc,
//   	TaskDefinition: task,
//   	PublicLoadBalancer: jsii.Boolean(false),
//   })
//
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("Nlb"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	CrossZoneEnabled: jsii.Boolean(true),
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   listener := nlb.AddListener(jsii.String("listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(80),
//   })
//
//   listener.AddTargets(jsii.String("Targets"), &AddNetworkTargetsProps{
//   	Targets: []INetworkLoadBalancerTarget{
//   		targets.NewAlbListenerTarget(svc.Listener),
//   	},
//   	Port: jsii.Number(80),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("NlbEndpoint"), &CfnOutputProps{
//   	Value: fmt.Sprintf("http://%v", nlb.LoadBalancerDnsName),
//   })
//
type AlbListenerTarget interface {
	AlbArnTarget
	// Register this ALB target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	//
	// This adds dependency on albListener because creation of ALB listener and NLB can vary during runtime.
	// More Details on - https://github.com/aws/aws-cdk/issues/17208
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbListenerTarget
type jsiiProxy_AlbListenerTarget struct {
	jsiiProxy_AlbArnTarget
}

// Create a new ALB target.
//
// The associated target group will automatically have a dependency added
// against the ALB's listener.
func NewAlbListenerTarget(albListener awselasticloadbalancingv2.ApplicationListener) AlbListenerTarget {
	_init_.Initialize()

	if err := validateNewAlbListenerTargetParameters(albListener); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListenerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbListenerTarget",
		[]interface{}{albListener},
		&j,
	)

	return &j
}

// Create a new ALB target.
//
// The associated target group will automatically have a dependency added
// against the ALB's listener.
func NewAlbListenerTarget_Override(a AlbListenerTarget, albListener awselasticloadbalancingv2.ApplicationListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbListenerTarget",
		[]interface{}{albListener},
		a,
	)
}

func (a *jsiiProxy_AlbListenerTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

