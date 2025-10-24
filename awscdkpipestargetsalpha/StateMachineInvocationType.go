package awscdkpipestargetsalpha


// InvocationType for invoking the State Machine.
//
// Example:
//   var sourceQueue Queue
//   var targetStateMachine IStateMachine
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
// See: https://docs.aws.amazon.com/eventbridge/latest/pipes-reference/API_PipeTargetStateMachineParameters.html
//
// Experimental.
type StateMachineInvocationType string

const (
	// Invoke StepFunction asynchronously (`StartExecution`).
	//
	// See https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html for more details.
	// Experimental.
	StateMachineInvocationType_FIRE_AND_FORGET StateMachineInvocationType = "FIRE_AND_FORGET"
	// Invoke StepFunction synchronously (`StartSyncExecution`) and wait for the execution to complete.
	//
	// See https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartSyncExecution.html for more details.
	// Experimental.
	StateMachineInvocationType_REQUEST_RESPONSE StateMachineInvocationType = "REQUEST_RESPONSE"
)

