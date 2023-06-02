package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for StartExecution.
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
type StepFunctionsStartExecutionProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The Step Functions state machine to start the execution on.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Pass the execution ID from the context object to the execution input.
	//
	// This allows the Step Functions UI to link child executions from parent executions, making it easier to trace execution flow across state machines.
	//
	// If you set this property to `true`, the `input` property must be an object (provided by `sfn.TaskInput.fromObject`) or omitted entirely.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-nested-workflows.html#nested-execution-startid
	//
	AssociateWithParent *bool `field:"optional" json:"associateWithParent" yaml:"associateWithParent"`
	// The JSON input for the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	Input awsstepfunctions.TaskInput `field:"optional" json:"input" yaml:"input"`
	// The name of the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

