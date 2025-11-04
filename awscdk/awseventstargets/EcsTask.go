package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Start a task on an ECS cluster.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster ICluster
//   var taskDefinition TaskDefinition
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	TaskCount: jsii.Number(1),
//
//   	// Overrides the cpu and memory values in the task definition
//   	Cpu: jsii.String("512"),
//   	Memory: jsii.String("512"),
//   }))
//
type EcsTask interface {
	awsevents.IRuleTarget
	// The security groups associated with the task.
	//
	// Only applicable with awsvpc network mode.
	// Default: - A new security group is created.
	//
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Allows using tasks as target of EventBridge events.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EcsTask
type jsiiProxy_EcsTask struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_EcsTask) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}


func NewEcsTask(props *EcsTaskProps) EcsTask {
	_init_.Initialize()

	if err := validateNewEcsTaskParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcsTask_Override(e EcsTask, props *EcsTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcsTask) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	if err := e.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

