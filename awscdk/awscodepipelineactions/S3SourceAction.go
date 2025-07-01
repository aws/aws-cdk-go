package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Source that is provided by a specific Amazon S3 object.
//
// Will trigger the pipeline as soon as the S3 object changes, but only if there is
// a CloudTrail Trail in the account that captures the S3 event.
//
// Example:
//   var sourceBucket bucket
//
//   // later:
//   var project pipelineProject
//   key := "some/key.zip"
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
//   	ActionName: jsii.String("S3Source"),
//   	BucketKey: key,
//   	Bucket: sourceBucket,
//   	Output: sourceOutput,
//   	VariablesNamespace: jsii.String("MyNamespace"),
//   })
//   codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	EnvironmentVariables: map[string]buildEnvironmentVariable{
//   		"VERSION_ID": &buildEnvironmentVariable{
//   			"value": sourceAction.variables.versionId,
//   		},
//   	},
//   })
//
type S3SourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *S3SourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for S3SourceAction
type jsiiProxy_S3SourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_S3SourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3SourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3SourceAction) Variables() *S3SourceVariables {
	var returns *S3SourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewS3SourceAction(props *S3SourceActionProps) S3SourceAction {
	_init_.Initialize()

	if err := validateNewS3SourceActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3SourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3SourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3SourceAction_Override(s S3SourceAction, props *S3SourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3SourceAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3SourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (s *jsiiProxy_S3SourceAction) Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBoundParameters(_scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3SourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (s *jsiiProxy_S3SourceAction) VariableExpression(variableName *string) *string {
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

