package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Fargate service running on an ECS cluster fronted by an application load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &ApplicationMultipleTargetGroupsFargateServiceProps{
//   	Cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   		Vpc: *Vpc,
//   	}),
//   	MemoryLimitMiB: jsii.Number(256),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	EnableExecuteCommand: jsii.Boolean(true),
//   	LoadBalancers: []ApplicationLoadBalancerProps{
//   		&ApplicationLoadBalancerProps{
//   			Name: jsii.String("lb"),
//   			IdleTimeout: awscdk.Duration_Seconds(jsii.Number(400)),
//   			DomainName: jsii.String("api.example.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("example.com"),
//   			}),
//   			Listeners: []ApplicationListenerProps{
//   				&ApplicationListenerProps{
//   					Name: jsii.String("listener"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   		&ApplicationLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			IdleTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
//   			DomainName: jsii.String("frontend.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("frontend.com"),
//   			}),
//   			Listeners: []ApplicationListenerProps{
//   				&ApplicationListenerProps{
//   					Name: jsii.String("listener2"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_*FromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []ApplicationTargetProps{
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(443),
//   			Listener: jsii.String("listener2"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
type ApplicationMultipleTargetGroupsFargateService interface {
	ApplicationMultipleTargetGroupsServiceBase
	// Determines whether the service will be assigned a public IP address.
	AssignPublicIp() *bool
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
	// The Fargate service in this construct.
	Service() awsecs.FargateService
	// The default target group for the service.
	// Deprecated: - Use `targetGroups` instead.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// The target groups of the service.
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	// The Fargate task definition in this construct.
	TaskDefinition() awsecs.FargateTaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for ApplicationMultipleTargetGroupsFargateService
type jsiiProxy_ApplicationMultipleTargetGroupsFargateService struct {
	jsiiProxy_ApplicationMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Listeners() *[]awselasticloadbalancingv2.ApplicationListener {
	var returns *[]awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LoadBalancers() *[]awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns *[]awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns *[]awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationMultipleTargetGroupsFargateService class.
func NewApplicationMultipleTargetGroupsFargateService(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) ApplicationMultipleTargetGroupsFargateService {
	_init_.Initialize()

	if err := validateNewApplicationMultipleTargetGroupsFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationMultipleTargetGroupsFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationMultipleTargetGroupsFargateService class.
func NewApplicationMultipleTargetGroupsFargateService_Override(a ApplicationMultipleTargetGroupsFargateService, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService)SetLogDriver(val awsecs.LogDriver) {
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
func ApplicationMultipleTargetGroupsFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationMultipleTargetGroupsFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationMultipleTargetGroupsFargateService_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) {
	if err := a.validateAddPortMappingForTargetsParameters(container, targets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) FindListener(name *string) awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener

	_jsii_.Invoke(
		a,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup {
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

