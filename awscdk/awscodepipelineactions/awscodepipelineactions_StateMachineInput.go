package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Represents the input for the StateMachine.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StateMachineInput interface {
	// When InputType is set to Literal (default), the Input field is used directly as the input for the state machine execution.
	//
	// Otherwise, the state machine is invoked with an empty JSON object {}.
	//
	// When InputType is set to FilePath, this field is required.
	// An input artifact is also required when InputType is set to FilePath.
	Input() interface{}
	// The optional input Artifact of the Action.
	//
	// If InputType is set to FilePath, this artifact is required
	// and is used to source the input for the state machine execution.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-StepFunctions.html#action-reference-StepFunctions-example
	//
	InputArtifact() awscodepipeline.Artifact
	// Optional StateMachine InputType InputType can be Literal or FilePath.
	InputType() *string
}

// The jsii proxy struct for StateMachineInput
type jsiiProxy_StateMachineInput struct {
	_ byte // padding
}

func (j *jsiiProxy_StateMachineInput) Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineInput) InputArtifact() awscodepipeline.Artifact {
	var returns awscodepipeline.Artifact
	_jsii_.Get(
		j,
		"inputArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineInput) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}


// When the input type is FilePath, input artifact and filepath must be specified.
func StateMachineInput_FilePath(inputFile awscodepipeline.ArtifactPath) StateMachineInput {
	_init_.Initialize()

	var returns StateMachineInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StateMachineInput",
		"filePath",
		[]interface{}{inputFile},
		&returns,
	)

	return returns
}

// When the input type is Literal, input value is passed directly to the state machine input.
func StateMachineInput_Literal(object *map[string]interface{}) StateMachineInput {
	_init_.Initialize()

	var returns StateMachineInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StateMachineInput",
		"literal",
		[]interface{}{object},
		&returns,
	)

	return returns
}

