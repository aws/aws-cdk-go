package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A Step Functions Task to call StartExecution on another state machine.
//
// It supports three service integration patterns: FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var input interface{}
//   var stateMachine stateMachine
//
//   startExecution := awscdk.Aws_stepfunctions_tasks.NewStartExecution(stateMachine, &StartExecutionProps{
//   	Input: map[string]interface{}{
//   		"inputKey": input,
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	Name: jsii.String("name"),
//   })
//
// Deprecated: - use 'StepFunctionsStartExecution'.
type StartExecution interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: - use 'StepFunctionsStartExecution'.
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for StartExecution
type jsiiProxy_StartExecution struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: - use 'StepFunctionsStartExecution'.
func NewStartExecution(stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) StartExecution {
	_init_.Initialize()

	if err := validateNewStartExecutionParameters(stateMachine, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StartExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StartExecution",
		[]interface{}{stateMachine, props},
		&j,
	)

	return &j
}

// Deprecated: - use 'StepFunctionsStartExecution'.
func NewStartExecution_Override(s StartExecution, stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StartExecution",
		[]interface{}{stateMachine, props},
		s,
	)
}

func (s *jsiiProxy_StartExecution) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := s.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

