package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Fargate service running on an ECS cluster fronted by an application load balancer.
//
// Example:
//   // One application load balancer with one listener and two target groups.
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("Service"), &applicationMultipleTargetGroupsFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageProps{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	targetGroups: []applicationTargetProps{
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(80),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(90),
//   			pathPattern: jsii.String("a/b/c"),
//   			priority: jsii.Number(10),
//   		},
//   	},
//   })
//
// Experimental.
type ApplicationMultipleTargetGroupsFargateService interface {
	ApplicationMultipleTargetGroupsServiceBase
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp() *bool
	// The cluster that hosts the service.
	// Experimental.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `internalDesiredCount` instead.
	DesiredCount() *float64
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service, if one is not provided.
	// Experimental.
	InternalDesiredCount() *float64
	// The default listener for the service (first added listener).
	// Experimental.
	Listener() awselasticloadbalancingv2.ApplicationListener
	// Experimental.
	Listeners() *[]awselasticloadbalancingv2.ApplicationListener
	// Experimental.
	SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener)
	// The default Application Load Balancer for the service (first added load balancer).
	// Experimental.
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	// Experimental.
	LogDriver() awsecs.LogDriver
	// Experimental.
	SetLogDriver(val awsecs.LogDriver)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Fargate service in this construct.
	// Experimental.
	Service() awsecs.FargateService
	// The default target group for the service.
	// Experimental.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// Experimental.
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	// Experimental.
	SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup)
	// The Fargate task definition in this construct.
	// Experimental.
	TaskDefinition() awsecs.FargateTaskDefinition
	// Experimental.
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	// Experimental.
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Experimental.
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	// Returns the default cluster.
	// Experimental.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
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
	// Experimental.
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
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

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
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

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
// Experimental.
func NewApplicationMultipleTargetGroupsFargateService(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) ApplicationMultipleTargetGroupsFargateService {
	_init_.Initialize()

	if err := validateNewApplicationMultipleTargetGroupsFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationMultipleTargetGroupsFargateService{}

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationMultipleTargetGroupsFargateService class.
// Experimental.
func NewApplicationMultipleTargetGroupsFargateService_Override(a ApplicationMultipleTargetGroupsFargateService, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService)SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener) {
	if err := j.validateSetListenersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService)SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService)SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) {
	if err := j.validateSetTargetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ApplicationMultipleTargetGroupsFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationMultipleTargetGroupsFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		"isConstruct",
		[]interface{}{x},
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

