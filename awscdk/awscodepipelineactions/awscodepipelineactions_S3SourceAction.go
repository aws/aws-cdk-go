package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Source that is provided by a specific Amazon S3 object.
//
// Will trigger the pipeline as soon as the S3 object changes, but only if there is
// a CloudTrail Trail in the account that captures the S3 event.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.AddS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		Bucket: sourceBucket,
//   		ObjectPrefix: key,
//   	},
//   }, &AddEventSelectorOptions{
//   	ReadWriteType: cloudtrail.ReadWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
//   	ActionName: jsii.String("S3Source"),
//   	BucketKey: key,
//   	Bucket: sourceBucket,
//   	Output: sourceOutput,
//   	Trigger: codepipeline_actions.S3Trigger_EVENTS,
//   })
//
// Experimental.
type S3SourceAction interface {
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
	Variables() *S3SourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	// Experimental.
	Bound(_scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Experimental.
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


// Experimental.
func NewS3SourceAction(props *S3SourceActionProps) S3SourceAction {
	_init_.Initialize()

	if err := validateNewS3SourceActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3SourceAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.S3SourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3SourceAction_Override(s S3SourceAction, props *S3SourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.S3SourceAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3SourceAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (s *jsiiProxy_S3SourceAction) Bound(_scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

