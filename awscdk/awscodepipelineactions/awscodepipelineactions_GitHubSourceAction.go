package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Source that is provided by a GitHub repository.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Read the secret from Secrets Manager
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("GitHub_Source"),
//   	owner: jsii.String("awslabs"),
//   	repo: jsii.String("aws-cdk"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	output: sourceOutput,
//   	branch: jsii.String("develop"),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type GitHubSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *GitHubSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for GitHubSourceAction
type jsiiProxy_GitHubSourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_GitHubSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubSourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubSourceAction) Variables() *GitHubSourceVariables {
	var returns *GitHubSourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewGitHubSourceAction(props *GitHubSourceActionProps) GitHubSourceAction {
	_init_.Initialize()

	if err := validateNewGitHubSourceActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewGitHubSourceAction_Override(g GitHubSourceAction, props *GitHubSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitHubSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := g.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubSourceAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := g.validateBoundParameters(scope, stage, _options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		g,
		"bound",
		[]interface{}{scope, stage, _options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := g.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		g,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubSourceAction) VariableExpression(variableName *string) *string {
	if err := g.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

