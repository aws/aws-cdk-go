package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an AWS Step Functions State Machine.
//
// Example:
//   var sourceQueue queue
//   var targetStateMachine iStateMachine
//
//
//   pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
//   	InvocationType: targets.StateMachineInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type SfnStateMachine interface {
	awscdkpipesalpha.ITarget
	// The ARN of the target resource.
	// Experimental.
	TargetArn() *string
	// Bind this target to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig
	// Grant the pipe role to push to the target.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for SfnStateMachine
type jsiiProxy_SfnStateMachine struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_SfnStateMachine) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSfnStateMachine(stateMachine awsstepfunctions.IStateMachine, parameters *SfnStateMachineParameters) SfnStateMachine {
	_init_.Initialize()

	if err := validateNewSfnStateMachineParameters(stateMachine, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_SfnStateMachine{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SfnStateMachine",
		[]interface{}{stateMachine, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewSfnStateMachine_Override(s SfnStateMachine, stateMachine awsstepfunctions.IStateMachine, parameters *SfnStateMachineParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SfnStateMachine",
		[]interface{}{stateMachine, parameters},
		s,
	)
}

func (s *jsiiProxy_SfnStateMachine) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SfnStateMachine) GrantPush(grantee awsiam.IRole) {
	if err := s.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee},
	)
}

