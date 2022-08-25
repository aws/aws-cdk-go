package awsstepfunctions


// AWS Step Functions integrates with services directly in the Amazon States Language.
//
// You can control these AWS services using service integration patterns:.
//
// Example:
//   // Define a state machine with one Pass state
//   child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &stateMachineProps{
//   	definition: sfn.chain.start(sfn.NewPass(this, jsii.String("PassState"))),
//   })
//
//   // Include the state machine in a Task state with callback pattern
//   task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &stepFunctionsStartExecutionProps{
//   	stateMachine: child,
//   	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
//   	input: sfn.taskInput.fromObject(map[string]interface{}{
//   		"token": sfn.JsonPath.taskToken,
//   		"foo": jsii.String("bar"),
//   	}),
//   	name: jsii.String("MyExecutionName"),
//   })
//
//   // Define a second state machine with the Task state above
//   // Define a second state machine with the Task state above
//   sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &stateMachineProps{
//   	definition: task,
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

