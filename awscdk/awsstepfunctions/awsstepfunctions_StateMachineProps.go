package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining a State Machine.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   inputArtifact := codepipeline.NewArtifact()
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.filePath(inputArtifact.atPath(jsii.String("assets/input.json"))),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StateMachineProps struct {
	// Definition for this state machine.
	Definition IChainable `field:"required" json:"definition" yaml:"definition"`
	// Defines what execution history events are logged and where they are logged.
	Logs *LogOptions `field:"optional" json:"logs" yaml:"logs"`
	// The execution role for the state machine service.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A name for the state machine.
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Type of the state machine.
	StateMachineType StateMachineType `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// Maximum run time for this state machine.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Specifies whether Amazon X-Ray tracing is enabled for this state machine.
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
}

