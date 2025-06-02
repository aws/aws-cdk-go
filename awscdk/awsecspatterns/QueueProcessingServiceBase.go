package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for QueueProcessingEc2Service and QueueProcessingFargateService services.
type QueueProcessingServiceBase interface {
	constructs.Construct
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
	// The SQS queue that the service will process from.
	SqsQueue() awssqs.IQueue
	// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
	ConfigureAutoscalingForService(service awsecs.BaseService)
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Grant SQS permissions to an ECS service.
	GrantPermissionsToService(service awsecs.BaseService)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for QueueProcessingServiceBase
type jsiiProxy_QueueProcessingServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_QueueProcessingServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval {
	var returns *[]*awsapplicationautoscaling.ScalingInterval
	_jsii_.Get(
		j,
		"scalingSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Secrets() *map[string]awsecs.Secret {
	var returns *map[string]awsecs.Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) SqsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


// Constructs a new instance of the QueueProcessingServiceBase class.
func NewQueueProcessingServiceBase_Override(q QueueProcessingServiceBase, scope constructs.Construct, id *string, props *QueueProcessingServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingServiceBase",
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
func QueueProcessingServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQueueProcessingServiceBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueProcessingServiceBase) ConfigureAutoscalingForService(service awsecs.BaseService) {
	if err := q.validateConfigureAutoscalingForServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"configureAutoscalingForService",
		[]interface{}{service},
	)
}

func (q *jsiiProxy_QueueProcessingServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (q *jsiiProxy_QueueProcessingServiceBase) GrantPermissionsToService(service awsecs.BaseService) {
	if err := q.validateGrantPermissionsToServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"grantPermissionsToService",
		[]interface{}{service},
	)
}

func (q *jsiiProxy_QueueProcessingServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

