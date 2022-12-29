package awscodepipeline


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
type StageOptions struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The reason for disabling transition to this stage.
	//
	// Only applicable
	// if `transitionToEnabled` is set to `false`.
	TransitionDisabledReason *string `field:"optional" json:"transitionDisabledReason" yaml:"transitionDisabledReason"`
	// Whether to enable transition to this stage.
	TransitionToEnabled *bool `field:"optional" json:"transitionToEnabled" yaml:"transitionToEnabled"`
	Placement *StagePlacement `field:"optional" json:"placement" yaml:"placement"`
}

