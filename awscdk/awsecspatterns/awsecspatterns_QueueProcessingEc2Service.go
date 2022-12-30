package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Class to create a queue processing EC2 service.
//
// Example:
//   var cluster cluster
//
//   queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &queueProcessingEc2ServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	image: ecs.containerImage.fromRegistry(jsii.String("test")),
//   	command: []*string{
//   		jsii.String("-c"),
//   		jsii.String("4"),
//   		jsii.String("amazon.com"),
//   	},
//   	enableLogging: jsii.Boolean(false),
//   	desiredTaskCount: jsii.Number(2),
//   	environment: map[string]*string{
//   		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
//   		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
//   	},
//   	maxScalingCapacity: jsii.Number(5),
//   	containerName: jsii.String("test"),
//   })
//
// Experimental.
type QueueProcessingEc2Service interface {
	QueueProcessingServiceBase
	// The cluster where your service will be deployed.
	// Experimental.
	Cluster() awsecs.ICluster
	// The dead letter queue for the primary SQS queue.
	// Experimental.
	DeadLetterQueue() awssqs.IQueue
	// The minimum number of tasks to run.
	// Deprecated: - Use `minCapacity` instead.
	DesiredCount() *float64
	// Environment variables that will include the queue name.
	// Experimental.
	Environment() *map[string]*string
	// The AwsLogDriver to use for logging if logging is enabled.
	// Experimental.
	LogDriver() awsecs.LogDriver
	// The maximum number of instances for autoscaling to scale up to.
	// Experimental.
	MaxCapacity() *float64
	// The minimum number of instances for autoscaling to scale down to.
	// Experimental.
	MinCapacity() *float64
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The scaling interval for autoscaling based off an SQS Queue size.
	// Experimental.
	ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval
	// The secret environment variables.
	// Experimental.
	Secrets() *map[string]awsecs.Secret
	// The EC2 service in this construct.
	// Experimental.
	Service() awsecs.Ec2Service
	// The SQS queue that the service will process from.
	// Experimental.
	SqsQueue() awssqs.IQueue
	// The EC2 task definition in this construct.
	// Experimental.
	TaskDefinition() awsecs.Ec2TaskDefinition
	// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
	// Experimental.
	ConfigureAutoscalingForService(service awsecs.BaseService)
	// Returns the default cluster.
	// Experimental.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Grant SQS permissions to an ECS service.
	// Experimental.
	GrantPermissionsToService(service awsecs.BaseService)
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

func (j *jsiiProxy_QueueProcessingEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
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

func (j *jsiiProxy_QueueProcessingEc2Service) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
// Experimental.
func NewQueueProcessingEc2Service(scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) QueueProcessingEc2Service {
	_init_.Initialize()

	if err := validateNewQueueProcessingEc2ServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueueProcessingEc2Service{}

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the QueueProcessingEc2Service class.
// Experimental.
func NewQueueProcessingEc2Service_Override(q QueueProcessingEc2Service, scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		q,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func QueueProcessingEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQueueProcessingEc2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs_patterns.QueueProcessingEc2Service",
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

func (q *jsiiProxy_QueueProcessingEc2Service) OnPrepare() {
	_jsii_.InvokeVoid(
		q,
		"onPrepare",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueProcessingEc2Service) OnSynthesize(session constructs.ISynthesisSession) {
	if err := q.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (q *jsiiProxy_QueueProcessingEc2Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueProcessingEc2Service) Prepare() {
	_jsii_.InvokeVoid(
		q,
		"prepare",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueProcessingEc2Service) Synthesize(session awscdk.ISynthesisSession) {
	if err := q.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"synthesize",
		[]interface{}{session},
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

func (q *jsiiProxy_QueueProcessingEc2Service) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

