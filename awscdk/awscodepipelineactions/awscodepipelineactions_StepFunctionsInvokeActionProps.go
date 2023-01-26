package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Construction properties of the `StepFunctionsInvokeAction StepFunction Invoke Action`.
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
type StepFunctionsInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The state machine to invoke.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Prefix (optional).
	//
	// By default, the action execution ID is used as the state machine execution name.
	// If a prefix is provided, it is prepended to the action execution ID with a hyphen and
	// together used as the state machine execution name.
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
	// The optional output Artifact of the Action.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// Represents the input to the StateMachine.
	//
	// This includes input artifact, input type and the statemachine input.
	StateMachineInput StateMachineInput `field:"optional" json:"stateMachineInput" yaml:"stateMachineInput"`
}

