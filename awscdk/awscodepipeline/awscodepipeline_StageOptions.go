package awscodepipeline


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
// Experimental.
type StageOptions struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The reason for disabling transition to this stage.
	//
	// Only applicable
	// if `transitionToEnabled` is set to `false`.
	// Experimental.
	TransitionDisabledReason *string `field:"optional" json:"transitionDisabledReason" yaml:"transitionDisabledReason"`
	// Whether to enable transition to this stage.
	// Experimental.
	TransitionToEnabled *bool `field:"optional" json:"transitionToEnabled" yaml:"transitionToEnabled"`
	// Experimental.
	Placement *StagePlacement `field:"optional" json:"placement" yaml:"placement"`
}

