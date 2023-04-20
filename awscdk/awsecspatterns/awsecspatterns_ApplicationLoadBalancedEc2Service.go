package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v3"
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
//   	},
//   	DesiredCount: jsii.Number(2),
//   })
//
// Experimental.
type ApplicationLoadBalancedEc2Service interface {
	ApplicationLoadBalancedServiceBase
	// Certificate Manager certificate to associate with the load balancer.
	// Experimental.
	Certificate() awscertificatemanager.ICertificate
	// The cluster that hosts the service.
	// Experimental.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `internalDesiredCount` instead.
	DesiredCount() *float64
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service if one is not provided.
	// Experimental.
	InternalDesiredCount() *float64
	// The listener for the service.
	// Experimental.
	Listener() awselasticloadbalancingv2.ApplicationListener
	// The Application Load Balancer for the service.
	// Experimental.
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The redirect listener for the service if redirectHTTP is enabled.
	// Experimental.
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	// The EC2 service in this construct.
	// Experimental.
	Service() awsecs.Ec2Service
	// The target group for the service.
	// Experimental.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// The EC2 Task Definition in this construct.
	// Experimental.
	TaskDefinition() awsecs.Ec2TaskDefinition
	// Adds service as a target of the target group.
	// Experimental.
	AddServiceAsTarget(service awsecs.BaseService)
	// Experimental.
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	// Experimental.
	GetDefaultCluster(scope awscdk.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
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

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
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

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
// Experimental.
func NewApplicationLoadBalancedEc2Service(scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) ApplicationLoadBalancedEc2Service {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancedEc2ServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancedEc2Service{}

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationLoadBalancedEc2Service class.
// Experimental.
func NewApplicationLoadBalancedEc2Service_Override(a ApplicationLoadBalancedEc2Service, scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ApplicationLoadBalancedEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancedEc2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
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

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) GetDefaultCluster(scope awscdk.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

