package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// StepFunctionInvokeAction that is provided by an AWS CodePipeline.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &StateMachineProps{
//   	Definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&StepFunctionsInvokeActionProps{
//   	ActionName: jsii.String("Invoke"),
//   	StateMachine: simpleStateMachine,
//   	StateMachineInput: codepipeline_actions.StateMachineInput_Literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("StepFunctions"),
//   	Actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StepFunctionInvokeAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for StepFunctionInvokeAction
type jsiiProxy_StepFunctionInvokeAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_StepFunctionInvokeAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionInvokeAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewStepFunctionInvokeAction(props *StepFunctionsInvokeActionProps) StepFunctionInvokeAction {
	_init_.Initialize()

	if err := validateNewStepFunctionInvokeActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionInvokeAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StepFunctionInvokeAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStepFunctionInvokeAction_Override(s StepFunctionInvokeAction, props *StepFunctionsInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StepFunctionInvokeAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionInvokeAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBoundParameters(_scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := s.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) VariableExpression(variableName *string) *string {
	if err := s.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

