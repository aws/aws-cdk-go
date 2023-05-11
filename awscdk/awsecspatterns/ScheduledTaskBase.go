package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for ScheduledEc2Task and ScheduledFargateTask tasks.
type ScheduledTaskBase interface {
	constructs.Construct
	// The name of the cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	DesiredTaskCount() *float64
	// The CloudWatch Events rule for the service.
	EventRule() awsevents.Rule
	// The tree node.
	Node() constructs.Node
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated.
	PropagateTags() awsecs.PropagatedTagSource
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	SubnetSelection() *awsec2.SubnetSelection
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags() *[]*awseventstargets.Tag
	// Adds task as a target of the scheduled event rule.
	AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask)
	// Create an ECS task using the task definition provided and add it to the scheduled event rule.
	AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask
	// Create an AWS Log Driver with the provided streamPrefix.
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ScheduledTaskBase
type jsiiProxy_ScheduledTaskBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ScheduledTaskBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) DesiredTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) EventRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) PropagateTags() awsecs.PropagatedTagSource {
	var returns awsecs.PropagatedTagSource
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) Tags() *[]*awseventstargets.Tag {
	var returns *[]*awseventstargets.Tag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScheduledTaskBase class.
func NewScheduledTaskBase_Override(s ScheduledTaskBase, scope constructs.Construct, id *string, props *ScheduledTaskBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledTaskBase",
		[]interface{}{scope, id, props},
		s,
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
func ScheduledTaskBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScheduledTaskBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledTaskBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledTaskBase) AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask) {
	if err := s.validateAddTaskAsTargetParameters(ecsTaskTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTaskAsTarget",
		[]interface{}{ecsTaskTarget},
	)
}

func (s *jsiiProxy_ScheduledTaskBase) AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask {
	if err := s.validateAddTaskDefinitionToEventTargetParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awseventstargets.EcsTask

	_jsii_.Invoke(
		s,
		"addTaskDefinitionToEventTarget",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledTaskBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	if err := s.validateCreateAWSLogDriverParameters(prefix); err != nil {
		panic(err)
	}
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		s,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledTaskBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := s.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		s,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledTaskBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

