package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Source that is provided by a GitHub repository.
//
// Example:
//   var sourceOutput artifact
//   var project pipelineProject
//
//
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("Github_Source"),
//   	output: sourceOutput,
//   	owner: jsii.String("my-owner"),
//   	repo: jsii.String("my-repo"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	variablesNamespace: jsii.String("MyNamespace"),
//   })
//
//   // later:
//
//   // later:
//   codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"COMMIT_URL": &buildEnvironmentVariable{
//   			"value": sourceAction.variables.commitUrl,
//   		},
//   	},
//   })
//
// Experimental.
type GitHubSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	// Experimental.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	// Experimental.
	Variables() *GitHubSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	// Experimental.
	Bound(scope awscdk.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Experimental.
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


// Experimental.
func NewGitHubSourceAction(props *GitHubSourceActionProps) GitHubSourceAction {
	_init_.Initialize()

	j := jsiiProxy_GitHubSourceAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubSourceAction_Override(g GitHubSourceAction, props *GitHubSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitHubSourceAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubSourceAction) Bound(scope awscdk.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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
	var returns *string

	_jsii_.Invoke(
		g,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

