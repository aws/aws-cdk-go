package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import patterns "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//
//   task := ecs.NewFargateTaskDefinition(this, jsii.String("Task"), &FargateTaskDefinitionProps{
//   	Cpu: jsii.Number(256),
//   	MemoryLimitMiB: jsii.Number(512),
//   })
//   task.AddContainer(jsii.String("nginx"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest")),
//   	PortMappings: []portMapping{
//   		&portMapping{
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
//   	Targets: []iNetworkLoadBalancerTarget{
//   		targets.NewAlbTarget(svc.loadBalancer, jsii.Number(80)),
//   	},
//   	Port: jsii.Number(80),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("NlbEndpoint"), &CfnOutputProps{
//   	Value: fmt.Sprintf("http://%v", nlb.LoadBalancerDnsName),
//   })
//
// Experimental.
type AlbTarget interface {
	AlbArnTarget
	// Register this alb target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Experimental.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbTarget
type jsiiProxy_AlbTarget struct {
	jsiiProxy_AlbArnTarget
}

// Experimental.
func NewAlbTarget(alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) AlbTarget {
	_init_.Initialize()

	if err := validateNewAlbTargetParameters(alb, port); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		&j,
	)

	return &j
}

// Experimental.
func NewAlbTarget_Override(a AlbTarget, alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		a,
	)
}

func (a *jsiiProxy_AlbTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
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

