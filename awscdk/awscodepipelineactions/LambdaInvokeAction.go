package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline invoke Action that is provided by an AWS Lambda function.
//
// Example:
//   var fn Function
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
//   	ActionName: jsii.String("Lambda"),
//   	Lambda: fn,
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Lambda"),
//   	Actions: []IAction{
//   		lambdaAction,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html
//
type LambdaInvokeAction interface {
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
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Reference a CodePipeline variable defined by the Lambda function this action points to.
	//
	// Variables in Lambda invoke actions are defined by calling the PutJobSuccessResult CodePipeline API call
	// with the 'outputVariables' property filled.
	// See: https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_PutJobSuccessResult.html
	//
	Variable(variableName *string) *string
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_LambdaInvokeAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewLambdaInvokeAction(props *LambdaInvokeActionProps) LambdaInvokeAction {
	_init_.Initialize()

	if err := validateNewLambdaInvokeActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.LambdaInvokeAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLambdaInvokeAction_Override(l LambdaInvokeAction, props *LambdaInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.LambdaInvokeAction",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LambdaInvokeAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := l.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := l.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		l,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := l.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		l,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) Variable(variableName *string) *string {
	if err := l.validateVariableParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"variable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) VariableExpression(variableName *string) *string {
	if err := l.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

