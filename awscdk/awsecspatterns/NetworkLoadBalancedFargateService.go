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
//   var vpc vpc
//   var securityGroup securityGroup
//
//   queueProcessingFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("Service"), &NetworkLoadBalancedFargateServiceProps{
//   	Vpc: Vpc,
//   	MemoryLimitMiB: jsii.Number(512),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   })
//
type NetworkLoadBalancedFargateService interface {
	NetworkLoadBalancedServiceBase
	AssignPublicIp() *bool
	// The cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service, if one is not provided.
	InternalDesiredCount() *float64
	// The listener for the service.
	Listener() awselasticloadbalancingv2.NetworkListener
	// The Network Load Balancer for the service.
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	// The tree node.
	Node() constructs.Node
	// The Fargate service in this construct.
	Service() awsecs.FargateService
	// The target group for the service.
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	// The Fargate task definition in this construct.
	TaskDefinition() awsecs.FargateTaskDefinition
	// Adds service as a target of the target group.
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancedFargateService
type jsiiProxy_NetworkLoadBalancedFargateService struct {
	jsiiProxy_NetworkLoadBalancedServiceBase
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkLoadBalancedFargateService class.
func NewNetworkLoadBalancedFargateService(scope constructs.Construct, id *string, props *NetworkLoadBalancedFargateServiceProps) NetworkLoadBalancedFargateService {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancedFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancedFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkLoadBalancedFargateService class.
func NewNetworkLoadBalancedFargateService_Override(n NetworkLoadBalancedFargateService, scope constructs.Construct, id *string, props *NetworkLoadBalancedFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		n,
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
func NetworkLoadBalancedFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkLoadBalancedFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) AddServiceAsTarget(service awsecs.BaseService) {
	if err := n.validateAddServiceAsTargetParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
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

func (n *jsiiProxy_NetworkLoadBalancedFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (n *jsiiProxy_NetworkLoadBalancedFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

