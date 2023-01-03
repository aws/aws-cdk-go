package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for ApplicationMultipleTargetGroupsEc2Service and ApplicationMultipleTargetGroupsFargateService classes.
type ApplicationMultipleTargetGroupsServiceBase interface {
	constructs.Construct
	// The cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service, if one is not provided.
	InternalDesiredCount() *float64
	// The default listener for the service (first added listener).
	// Deprecated: - Use `listeners` instead.
	Listener() awselasticloadbalancingv2.ApplicationListener
	// The listeners of the service.
	Listeners() *[]awselasticloadbalancingv2.ApplicationListener
	// The default Application Load Balancer for the service (first added load balancer).
	// Deprecated: - Use `loadBalancers` instead.
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	// The load balancers of the service.
	LoadBalancers() *[]awselasticloadbalancingv2.ApplicationLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	// The tree node.
	Node() constructs.Node
	// The target groups of the service.
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApplicationMultipleTargetGroupsServiceBase
type jsiiProxy_ApplicationMultipleTargetGroupsServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Listeners() *[]awselasticloadbalancingv2.ApplicationListener {
	var returns *[]awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) LoadBalancers() *[]awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns *[]awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns *[]awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationMultipleTargetGroupsServiceBase class.
func NewApplicationMultipleTargetGroupsServiceBase_Override(a ApplicationMultipleTargetGroupsServiceBase, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsServiceBase",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase)SetLogDriver(val awsecs.LogDriver) {
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
func ApplicationMultipleTargetGroupsServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationMultipleTargetGroupsServiceBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) {
	if err := a.validateAddPortMappingForTargetsParameters(container, targets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	if err := a.validateCreateAWSLogDriverParameters(prefix); err != nil {
		panic(err)
	}
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) FindListener(name *string) awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener

	_jsii_.Invoke(
		a,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := a.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup {
	if err := a.validateRegisterECSTargetsParameters(service, container, targets); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

