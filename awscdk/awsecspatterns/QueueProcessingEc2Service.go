package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to create a queue processing EC2 service.
//
// Example:
//   var cluster Cluster
//
//   queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &QueueProcessingEc2ServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
//   	Command: []*string{
//   		jsii.String("-c"),
//   		jsii.String("4"),
//   		jsii.String("amazon.com"),
//   	},
//   	EnableLogging: jsii.Boolean(false),
//   	DesiredTaskCount: jsii.Number(2),
//   	Environment: map[string]*string{
//   		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
//   		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
//   	},
//   	MaxScalingCapacity: jsii.Number(5),
//   	ContainerName: jsii.String("test"),
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
type QueueProcessingEc2Service interface {
	QueueProcessingServiceBase
	// The cluster where your service will be deployed.
	Cluster() awsecs.ICluster
	// The dead letter queue for the primary SQS queue.
	DeadLetterQueue() awssqs.IQueue
	// Environment variables that will include the queue name.
	Environment() *map[string]*string
	// The AwsLogDriver to use for logging if logging is enabled.
	LogDriver() awsecs.LogDriver
	// The maximum number of instances for autoscaling to scale up to.
	MaxCapacity() *float64
	// The minimum number of instances for autoscaling to scale down to.
	MinCapacity() *float64
	// The tree node.
	Node() constructs.Node
	// The scaling interval for autoscaling based off an SQS Queue size.
	ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval
	// The secret environment variables.
	Secrets() *map[string]awsecs.Secret
	// The EC2 service in this construct.
	Service() awsecs.Ec2Service
	// The SQS queue that the service will process from.
	SqsQueue() awssqs.IQueue
	// The EC2 task definition in this construct.
	TaskDefinition() awsecs.Ec2TaskDefinition
	// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
	ConfigureAutoscalingForService(service awsecs.BaseService)
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Grant SQS permissions to an ECS service.
	GrantPermissionsToService(service awsecs.BaseService)
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

// The jsii proxy struct for QueueProcessingEc2Service
type jsiiProxy_QueueProcessingEc2Service struct {
	jsiiProxy_QueueProcessingServiceBase
}

func (j *jsiiProxy_QueueProcessingEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval {
	var returns *[]*awsapplicationautoscaling.ScalingInterval
	_jsii_.Get(
		j,
		"scalingSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Secrets() *map[string]awsecs.Secret {
	var returns *map[string]awsecs.Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) SqsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the QueueProcessingEc2Service class.
func NewQueueProcessingEc2Service(scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) QueueProcessingEc2Service {
	_init_.Initialize()

	if err := validateNewQueueProcessingEc2ServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueueProcessingEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the QueueProcessingEc2Service class.
func NewQueueProcessingEc2Service_Override(q QueueProcessingEc2Service, scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		q,
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
func QueueProcessingEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQueueProcessingEc2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueProcessingEc2Service) ConfigureAutoscalingForService(service awsecs.BaseService) {
	if err := q.validateConfigureAutoscalingForServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"configureAutoscalingForService",
		[]interface{}{service},
	)
}

func (q *jsiiProxy_QueueProcessingEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := q.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		q,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueProcessingEc2Service) GrantPermissionsToService(service awsecs.BaseService) {
	if err := q.validateGrantPermissionsToServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"grantPermissionsToService",
		[]interface{}{service},
	)
}

func (q *jsiiProxy_QueueProcessingEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueProcessingEc2Service) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		q,
		"with",
		args,
		&returns,
	)

	return returns
}

