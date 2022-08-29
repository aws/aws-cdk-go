package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import patterns "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//
//   task := ecs.NewFargateTaskDefinition(this, jsii.String("Task"), &fargateTaskDefinitionProps{
//   	cpu: jsii.Number(256),
//   	memoryLimitMiB: jsii.Number(512),
//   })
//   task.addContainer(jsii.String("nginx"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest")),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   })
//
//   svc := patterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	vpc: vpc,
//   	taskDefinition: task,
//   	publicLoadBalancer: jsii.Boolean(false),
//   })
//
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("Nlb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   	crossZoneEnabled: jsii.Boolean(true),
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   listener := nlb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//
//   listener.addTargets(jsii.String("Targets"), &addNetworkTargetsProps{
//   	targets: []iNetworkLoadBalancerTarget{
//   		targets.NewAlbTarget(svc.loadBalancer, jsii.Number(80)),
//   	},
//   	port: jsii.Number(80),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("NlbEndpoint"), &cfnOutputProps{
//   	value: fmt.Sprintf("http://%v", nlb.loadBalancerDnsName),
//   })
//
type AlbTarget interface {
	AlbArnTarget
	// Register this alb target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbTarget
type jsiiProxy_AlbTarget struct {
	jsiiProxy_AlbArnTarget
}

func NewAlbTarget(alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) AlbTarget {
	_init_.Initialize()

	j := jsiiProxy_AlbTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		&j,
	)

	return &j
}

func NewAlbTarget_Override(a AlbTarget, alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		a,
	)
}

func (a *jsiiProxy_AlbTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

