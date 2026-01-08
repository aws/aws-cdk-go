package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstepfunctions"
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
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage awsstepfunctions.QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// The Step Functions state machine to start the execution on.
	StateMachine interfacesawsstepfunctions.IStateMachineRef `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Pass the execution ID from the context object to the execution input.
	//
	// This allows the Step Functions UI to link child executions from parent executions, making it easier to trace execution flow across state machines.
	//
	// If you set this property to `true`, the `input` property must be an object (provided by `sfn.TaskInput.fromObject`) or omitted entirely.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-nested-workflows.html#nested-execution-startid
	//
	// Default: - false.
	//
	AssociateWithParent *bool `field:"optional" json:"associateWithParent" yaml:"associateWithParent"`
	// The JSON input for the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Default: - The state input (JSON path '$').
	//
	Input awsstepfunctions.TaskInput `field:"optional" json:"input" yaml:"input"`
	// The name of the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Default: - None.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

