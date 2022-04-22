package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import elasticloadbalancingv2_targets "github.com/aws/aws-cdk-go/awscdk/aws_elasticloadbalancingv2_targets"
//   albArnTarget := elasticloadbalancingv2_targets.NewAlbArnTarget(jsii.String("albArn"), jsii.Number(123))
//
type AlbArnTarget interface {
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// Register this alb target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbArnTarget
type jsiiProxy_AlbArnTarget struct {
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// Create a new alb target.
func NewAlbArnTarget(albArn *string, port *float64) AlbArnTarget {
	_init_.Initialize()

	j := jsiiProxy_AlbArnTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbArnTarget",
		[]interface{}{albArn, port},
		&j,
	)

	return &j
}

// Create a new alb target.
func NewAlbArnTarget_Override(a AlbArnTarget, albArn *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbArnTarget",
		[]interface{}{albArn, port},
		a,
	)
}

func (a *jsiiProxy_AlbArnTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   import targets "github.com/aws/aws-cdk-go/awscdk"import ecs "github.com/aws/aws-cdk-go/awscdk"import patterns "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
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
//   NewCfnOutput(this, jsii.String("NlbEndpoint"), &cfnOutputProps{
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

// An EC2 instance that is the target for load balancing.
//
// If you register a target of this type, you are responsible for making
// sure the load balancer's security group can connect to the instance.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import elasticloadbalancingv2_targets "github.com/aws/aws-cdk-go/awscdk/aws_elasticloadbalancingv2_targets"
//   instanceIdTarget := elasticloadbalancingv2_targets.NewInstanceIdTarget(jsii.String("instanceId"), jsii.Number(123))
//
type InstanceIdTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for InstanceIdTarget
type jsiiProxy_InstanceIdTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// Create a new Instance target.
func NewInstanceIdTarget(instanceId *string, port *float64) InstanceIdTarget {
	_init_.Initialize()

	j := jsiiProxy_InstanceIdTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceIdTarget",
		[]interface{}{instanceId, port},
		&j,
	)

	return &j
}

// Create a new Instance target.
func NewInstanceIdTarget_Override(i InstanceIdTarget, instanceId *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceIdTarget",
		[]interface{}{instanceId, port},
		i,
	)
}

func (i *jsiiProxy_InstanceIdTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceIdTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import elasticloadbalancingv2_targets "github.com/aws/aws-cdk-go/awscdk/aws_elasticloadbalancingv2_targets"
//
//   var instance instance
//   instanceTarget := elasticloadbalancingv2_targets.NewInstanceTarget(instance, jsii.Number(123))
//
type InstanceTarget interface {
	InstanceIdTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for InstanceTarget
type jsiiProxy_InstanceTarget struct {
	jsiiProxy_InstanceIdTarget
}

// Create a new Instance target.
func NewInstanceTarget(instance awsec2.Instance, port *float64) InstanceTarget {
	_init_.Initialize()

	j := jsiiProxy_InstanceTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceTarget",
		[]interface{}{instance, port},
		&j,
	)

	return &j
}

// Create a new Instance target.
func NewInstanceTarget_Override(i InstanceTarget, instance awsec2.Instance, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceTarget",
		[]interface{}{instance, port},
		i,
	)
}

func (i *jsiiProxy_InstanceTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// An IP address that is a target for load balancing.
//
// Specify IP addresses from the subnets of the virtual private cloud (VPC) for
// the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and
// 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify
// publicly routable IP addresses.
//
// If you register a target of this type, you are responsible for making
// sure the load balancer's security group can send packets to the IP address.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import elasticloadbalancingv2_targets "github.com/aws/aws-cdk-go/awscdk/aws_elasticloadbalancingv2_targets"
//   ipTarget := elasticloadbalancingv2_targets.NewIpTarget(jsii.String("ipAddress"), jsii.Number(123), jsii.String("availabilityZone"))
//
type IpTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for IpTarget
type jsiiProxy_IpTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// Create a new IPAddress target.
//
// The availabilityZone parameter determines whether the target receives
// traffic from the load balancer nodes in the specified Availability Zone
// or from all enabled Availability Zones for the load balancer.
//
// This parameter is not supported if the target type of the target group
// is instance. If the IP address is in a subnet of the VPC for the target
// group, the Availability Zone is automatically detected and this
// parameter is optional. If the IP address is outside the VPC, this
// parameter is required.
//
// With an Application Load Balancer, if the IP address is outside the VPC
// for the target group, the only supported value is all.
//
// Default is automatic.
func NewIpTarget(ipAddress *string, port *float64, availabilityZone *string) IpTarget {
	_init_.Initialize()

	j := jsiiProxy_IpTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.IpTarget",
		[]interface{}{ipAddress, port, availabilityZone},
		&j,
	)

	return &j
}

// Create a new IPAddress target.
//
// The availabilityZone parameter determines whether the target receives
// traffic from the load balancer nodes in the specified Availability Zone
// or from all enabled Availability Zones for the load balancer.
//
// This parameter is not supported if the target type of the target group
// is instance. If the IP address is in a subnet of the VPC for the target
// group, the Availability Zone is automatically detected and this
// parameter is optional. If the IP address is outside the VPC, this
// parameter is required.
//
// With an Application Load Balancer, if the IP address is outside the VPC
// for the target group, the only supported value is all.
//
// Default is automatic.
func NewIpTarget_Override(i IpTarget, ipAddress *string, port *float64, availabilityZone *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.IpTarget",
		[]interface{}{ipAddress, port, availabilityZone},
		i,
	)
}

func (i *jsiiProxy_IpTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var lambdaFunction function
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("Targets"), &addApplicationTargetsProps{
//   	targets: []iApplicationLoadBalancerTarget{
//   		targets.NewLambdaTarget(lambdaFunction),
//   	},
//
//   	// For Lambda Targets, you need to explicitly enable health checks if you
//   	// want them.
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(true),
//   	},
//   })
//
type LambdaTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for LambdaTarget
type jsiiProxy_LambdaTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
}

// Create a new Lambda target.
func NewLambdaTarget(fn awslambda.IFunction) LambdaTarget {
	_init_.Initialize()

	j := jsiiProxy_LambdaTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		&j,
	)

	return &j
}

// Create a new Lambda target.
func NewLambdaTarget_Override(l LambdaTarget, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		l,
	)
}

func (l *jsiiProxy_LambdaTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

