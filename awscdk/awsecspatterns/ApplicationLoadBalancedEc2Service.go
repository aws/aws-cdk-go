package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// An EC2 service running on an ECS cluster fronted by an application load balancer.
//
// Example:
//   var cluster cluster
//
//   loadBalancedEcsService := ecsPatterns.NewApplicationLoadBalancedEc2Service(this, jsii.String("Service"), &ApplicationLoadBalancedEc2ServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
//   		Environment: map[string]*string{
//   			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
//   			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
//   		},
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entry"),
//   			jsii.String("point"),
//   		},
//   	},
//   	DesiredCount: jsii.Number(2),
//   })
//
type ApplicationLoadBalancedEc2Service interface {
	ApplicationLoadBalancedServiceBase
	// Certificate Manager certificate to associate with the load balancer.
	Certificate() awscertificatemanager.ICertificate
	// The cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service if one is not provided.
	InternalDesiredCount() *float64
	// The listener for the service.
	Listener() awselasticloadbalancingv2.ApplicationListener
	// The Application Load Balancer for the service.
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	// The tree node.
	Node() constructs.Node
	// The redirect listener for the service if redirectHTTP is enabled.
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	// The EC2 service in this construct.
	Service() awsecs.Ec2Service
	// The target group for the service.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// The EC2 Task Definition in this construct.
	TaskDefinition() awsecs.Ec2TaskDefinition
	// Adds service as a target of the target group.
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancedEc2Service
type jsiiProxy_ApplicationLoadBalancedEc2Service struct {
	jsiiProxy_ApplicationLoadBalancedServiceBase
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) RedirectListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"redirectListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationLoadBalancedEc2Service class.
func NewApplicationLoadBalancedEc2Service(scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) ApplicationLoadBalancedEc2Service {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancedEc2ServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancedEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationLoadBalancedEc2Service class.
func NewApplicationLoadBalancedEc2Service_Override(a ApplicationLoadBalancedEc2Service, scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		a,
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
func ApplicationLoadBalancedEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancedEc2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) AddServiceAsTarget(service awsecs.BaseService) {
	if err := a.validateAddServiceAsTargetParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
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

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

