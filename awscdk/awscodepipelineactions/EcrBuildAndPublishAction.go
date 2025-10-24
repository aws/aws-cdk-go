package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline build action that uses AWS EcrBuildAndPublish.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline Pipeline
//   var repository IRepository
//
//
//   sourceOutput := codepipeline.NewArtifact()
//   // your source repository
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
//   	ActionName: jsii.String("CodeStarConnectionsSourceAction"),
//   	Output: sourceOutput,
//   	ConnectionArn: jsii.String("your-connection-arn"),
//   	Owner: jsii.String("your-owner"),
//   	Repo: jsii.String("your-repo"),
//   })
//
//   buildAction := codepipeline_actions.NewEcrBuildAndPublishAction(&EcrBuildAndPublishActionProps{
//   	ActionName: jsii.String("EcrBuildAndPublishAction"),
//   	RepositoryName: repository.RepositoryName,
//   	RegistryType: codepipeline_actions.RegistryType_PRIVATE,
//   	DockerfileDirectoryPath: jsii.String("./my-dir"),
//   	 // The path indicates ./my-dir/Dockerfile in the source repository
//   	ImageTags: []*string{
//   		jsii.String("my-tag-1"),
//   		jsii.String("my-tag-2"),
//   	},
//   	Input: sourceOutput,
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []IAction{
//   		sourceAction,
//   	},
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Build"),
//   	Actions: []IAction{
//   		buildAction,
//   	},
//   })
//
type EcrBuildAndPublishAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *EcrBuildAndPublishVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for EcrBuildAndPublishAction
type jsiiProxy_EcrBuildAndPublishAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_EcrBuildAndPublishAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrBuildAndPublishAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrBuildAndPublishAction) Variables() *EcrBuildAndPublishVariables {
	var returns *EcrBuildAndPublishVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewEcrBuildAndPublishAction(props *EcrBuildAndPublishActionProps) EcrBuildAndPublishAction {
	_init_.Initialize()

	if err := validateNewEcrBuildAndPublishActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrBuildAndPublishAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcrBuildAndPublishAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcrBuildAndPublishAction_Override(e EcrBuildAndPublishAction, props *EcrBuildAndPublishActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcrBuildAndPublishAction",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcrBuildAndPublishAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrBuildAndPublishAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrBuildAndPublishAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := e.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		e,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrBuildAndPublishAction) VariableExpression(variableName *string) *string {
	if err := e.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

