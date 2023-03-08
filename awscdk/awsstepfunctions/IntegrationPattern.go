package awsstepfunctions


// AWS Step Functions integrates with services directly in the Amazon States Language.
//
// You can control these AWS services using service integration patterns:.
//
// Example:
//   // Define a state machine with one Pass state
//   child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &StateMachineProps{
//   	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("PassState"))),
//   })
//
//   // Include the state machine in a Task state with callback pattern
//   task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &StepFunctionsStartExecutionProps{
//   	StateMachine: child,
//   	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
//   	Input: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"token": sfn.JsonPath_taskToken(),
//   		"foo": jsii.String("bar"),
//   	}),
//   	Name: jsii.String("MyExecutionName"),
//   })
//
//   // Define a second state machine with the Task state above
//   // Define a second state machine with the Task state above
//   sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &StateMachineProps{
//   	Definition: task,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
//
type IntegrationPattern string

const (
	// Step Functions will wait for an HTTP response and then progress to the next state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-default
	//
	IntegrationPattern_REQUEST_RESPONSE IntegrationPattern = "REQUEST_RESPONSE"
	// Step Functions can wait for a request to complete before progressing to the next state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
	//
	IntegrationPattern_RUN_JOB IntegrationPattern = "RUN_JOB"
	// Callback tasks provide a way to pause a workflow until a task token is returned.
	//
	// You must set a task token when using the callback pattern.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern_WAIT_FOR_TASK_TOKEN IntegrationPattern = "WAIT_FOR_TASK_TOKEN"
)

