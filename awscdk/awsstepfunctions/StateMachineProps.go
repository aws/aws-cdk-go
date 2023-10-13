package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining a State Machine.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   inputArtifact := codepipeline.NewArtifact()
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &StateMachineProps{
//   	Definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&StepFunctionsInvokeActionProps{
//   	ActionName: jsii.String("Invoke"),
//   	StateMachine: simpleStateMachine,
//   	StateMachineInput: codepipeline_actions.StateMachineInput_FilePath(inputArtifact.AtPath(jsii.String("assets/input.json"))),
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("StepFunctions"),
//   	Actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StateMachineProps struct {
	// Comment that describes this state machine.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Definition for this state machine.
	// Deprecated: use definitionBody: DefinitionBody.fromChainable()
	Definition IChainable `field:"optional" json:"definition" yaml:"definition"`
	// Definition for this state machine.
	DefinitionBody DefinitionBody `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// substitutions for the definition body as a key-value map.
	DefinitionSubstitutions *map[string]*string `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// Defines what execution history events are logged and where they are logged.
	// Default: No logging.
	//
	Logs *LogOptions `field:"optional" json:"logs" yaml:"logs"`
	// The removal policy to apply to state machine.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The execution role for the state machine service.
	// Default: A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A name for the state machine.
	// Default: A name is automatically generated.
	//
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Type of the state machine.
	// Default: StateMachineType.STANDARD
	//
	StateMachineType StateMachineType `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// Maximum run time for this state machine.
	// Default: No timeout.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Specifies whether Amazon X-Ray tracing is enabled for this state machine.
	// Default: false.
	//
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
}

