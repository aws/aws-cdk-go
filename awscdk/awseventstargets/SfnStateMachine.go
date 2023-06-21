package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Use a StepFunctions state machine as a target for Amazon EventBridge rules.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
//   })
//   stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &StateMachineProps{
//   	Definition: sfn.NewWait(this, jsii.String("Hello"), &WaitProps{
//   		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(10))),
//   	}),
//   })
//
//   rule.AddTarget(targets.NewSfnStateMachine(stateMachine, &SfnStateMachineProps{
//   	Input: events.RuleTargetInput_FromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	DeadLetterQueue: dlq,
//   	Role: role,
//   }))
//
type SfnStateMachine interface {
	awsevents.IRuleTarget
	Machine() awsstepfunctions.IStateMachine
	// Returns a properties that are used in an Rule to trigger this State Machine.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
	//
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SfnStateMachine
type jsiiProxy_SfnStateMachine struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SfnStateMachine) Machine() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}


func NewSfnStateMachine(machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) SfnStateMachine {
	_init_.Initialize()

	if err := validateNewSfnStateMachineParameters(machine, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SfnStateMachine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		&j,
	)

	return &j
}

func NewSfnStateMachine_Override(s SfnStateMachine, machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		s,
	)
}

func (s *jsiiProxy_SfnStateMachine) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := s.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

