package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets"
	"github.com/aws/constructs-go/constructs/v3"
)

// A scheduled EC2 task that will be initiated off of CloudWatch Events.
//
// Example:
//   // Instantiate an Amazon EC2 Task to run at a scheduled interval
//   var cluster cluster
//
//   ecsScheduledTask := ecsPatterns.NewScheduledEc2Task(this, jsii.String("ScheduledTask"), &ScheduledEc2TaskProps{
//   	Cluster: Cluster,
//   	ScheduledEc2TaskImageOptions: &ScheduledEc2TaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		MemoryLimitMiB: jsii.Number(256),
//   		Environment: map[string]*string{
//   			"name": jsii.String("TRIGGER"),
//   			"value": jsii.String("CloudWatch Events"),
//   		},
//   	},
//   	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
//   	Enabled: jsii.Boolean(true),
//   	RuleName: jsii.String("sample-scheduled-task-rule"),
//   })
//
// Experimental.
type ScheduledEc2Task interface {
	ScheduledTaskBase
	// The name of the cluster that hosts the service.
	// Experimental.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Experimental.
	DesiredTaskCount() *float64
	// The CloudWatch Events rule for the service.
	// Experimental.
	EventRule() awsevents.Rule
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	// Experimental.
	SubnetSelection() *awsec2.SubnetSelection
	// The ECS task in this construct.
	// Experimental.
	Task() awseventstargets.EcsTask
	// The EC2 task definition in this construct.
	// Experimental.
	TaskDefinition() awsecs.Ec2TaskDefinition
	// Adds task as a target of the scheduled event rule.
	// Experimental.
	AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask)
	// Create an ECS task using the task definition provided and add it to the scheduled event rule.
	// Experimental.
	AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask
	// Create an AWS Log Driver with the provided streamPrefix.
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

// The jsii proxy struct for ScheduledEc2Task
type jsiiProxy_ScheduledEc2Task struct {
	jsiiProxy_ScheduledTaskBase
}

func (j *jsiiProxy_ScheduledEc2Task) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) DesiredTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) EventRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) Task() awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScheduledEc2Task class.
// Experimental.
func NewScheduledEc2Task(scope constructs.Construct, id *string, props *ScheduledEc2TaskProps) ScheduledEc2Task {
	_init_.Initialize()

	if err := validateNewScheduledEc2TaskParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScheduledEc2Task{}

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ScheduledEc2Task",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScheduledEc2Task class.
// Experimental.
func NewScheduledEc2Task_Override(s ScheduledEc2Task, scope constructs.Construct, id *string, props *ScheduledEc2TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs_patterns.ScheduledEc2Task",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ScheduledEc2Task_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScheduledEc2Task_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs_patterns.ScheduledEc2Task",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledEc2Task) AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask) {
	if err := s.validateAddTaskAsTargetParameters(ecsTaskTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTaskAsTarget",
		[]interface{}{ecsTaskTarget},
	)
}

func (s *jsiiProxy_ScheduledEc2Task) AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask {
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

func (s *jsiiProxy_ScheduledEc2Task) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
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

func (s *jsiiProxy_ScheduledEc2Task) GetDefaultCluster(scope awscdk.Construct, vpc awsec2.IVpc) awsecs.Cluster {
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

func (s *jsiiProxy_ScheduledEc2Task) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScheduledEc2Task) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ScheduledEc2Task) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledEc2Task) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScheduledEc2Task) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ScheduledEc2Task) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledEc2Task) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

