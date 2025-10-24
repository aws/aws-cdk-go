package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Fargate service running on an ECS cluster fronted by a network load balancer.
//
// Example:
//   // Two network load balancers, each with their own listener and target group.
//   var cluster Cluster
//
//   loadBalancedFargateService := ecsPatterns.NewNetworkMultipleTargetGroupsFargateService(this, jsii.String("Service"), &NetworkMultipleTargetGroupsFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(512),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	LoadBalancers: []NetworkLoadBalancerProps{
//   		&NetworkLoadBalancerProps{
//   			Name: jsii.String("lb1"),
//   			Listeners: []NetworkListenerProps{
//   				&NetworkListenerProps{
//   					Name: jsii.String("listener1"),
//   				},
//   			},
//   		},
//   		&NetworkLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			Listeners: []NetworkListenerProps{
//   				&NetworkListenerProps{
//   					Name: jsii.String("listener2"),
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []NetworkTargetProps{
//   		&NetworkTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener1"),
//   		},
//   		&NetworkTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   	MinHealthyPercent: jsii.Number(100),
//   	MaxHealthyPercent: jsii.Number(200),
//   })
//
type NetworkMultipleTargetGroupsFargateService interface {
	NetworkMultipleTargetGroupsServiceBase
	// Determines whether the service will be assigned a public IP address.
	AssignPublicIp() *bool
	// The cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service, if one is not provided.
	InternalDesiredCount() *float64
	// The listener for the service.
	// Deprecated: - Use `listeners` instead.
	Listener() awselasticloadbalancingv2.NetworkListener
	// The listeners of the service.
	Listeners() *[]awselasticloadbalancingv2.NetworkListener
	// The Network Load Balancer for the service.
	// Deprecated: - Use `loadBalancers` instead.
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	// The load balancers of the service.
	LoadBalancers() *[]awselasticloadbalancingv2.NetworkLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	// The tree node.
	Node() constructs.Node
	// The Fargate service in this construct.
	Service() awsecs.FargateService
	// The default target group for the service.
	// Deprecated: - Use `targetGroups` instead.
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	// The target groups of the service.
	TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup
	// The Fargate task definition in this construct.
	TaskDefinition() awsecs.FargateTaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.NetworkListener
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkMultipleTargetGroupsFargateService
type jsiiProxy_NetworkMultipleTargetGroupsFargateService struct {
	jsiiProxy_NetworkMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Listeners() *[]awselasticloadbalancingv2.NetworkListener {
	var returns *[]awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) LoadBalancers() *[]awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns *[]awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup {
	var returns *[]awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkMultipleTargetGroupsFargateService class.
func NewNetworkMultipleTargetGroupsFargateService(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) NetworkMultipleTargetGroupsFargateService {
	_init_.Initialize()

	if err := validateNewNetworkMultipleTargetGroupsFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkMultipleTargetGroupsFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkMultipleTargetGroupsFargateService class.
func NewNetworkMultipleTargetGroupsFargateService_Override(n NetworkMultipleTargetGroupsFargateService, scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService)SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetworkMultipleTargetGroupsFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkMultipleTargetGroupsFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkMultipleTargetGroupsFargateService_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) {
	if err := n.validateAddPortMappingForTargetsParameters(container, targets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	if err := n.validateCreateAWSLogDriverParameters(prefix); err != nil {
		panic(err)
	}
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) FindListener(name *string) awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener

	_jsii_.Invoke(
		n,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := n.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup {
	if err := n.validateRegisterECSTargetsParameters(service, container, targets); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.NetworkTargetGroup

	_jsii_.Invoke(
		n,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

