package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to invoke a pipeline.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   targetPipeline := codepipeline.Pipeline_FromPipelineArn(this, jsii.String("Pipeline"), jsii.String("arn:aws:codepipeline:us-east-1:123456789012:InvokePipelineAction")) // If targetPipeline is not created by cdk, import from arn.
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("stageName"),
//   	Actions: []iAction{
//   		cpactions.NewPipelineInvokeAction(&PipelineInvokeActionProps{
//   			ActionName: jsii.String("Invoke"),
//   			TargetPipeline: *TargetPipeline,
//   			Variables: []variable{
//   				&variable{
//   					Name: jsii.String("name1"),
//   					Value: jsii.String("value1"),
//   				},
//   			},
//   			SourceRevisions: []sourceRevision{
//   				&sourceRevision{
//   					ActionName: jsii.String("Source"),
//   					RevisionType: cpactions.RevisionType_S3_OBJECT_VERSION_ID,
//   					RevisionValue: jsii.String("testRevisionValue"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
type PipelineInvokeAction interface {
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
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for PipelineInvokeAction
type jsiiProxy_PipelineInvokeAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_PipelineInvokeAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineInvokeAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewPipelineInvokeAction(props *PipelineInvokeActionProps) PipelineInvokeAction {
	_init_.Initialize()

	if err := validateNewPipelineInvokeActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineInvokeAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.PipelineInvokeAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPipelineInvokeAction_Override(p PipelineInvokeAction, props *PipelineInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.PipelineInvokeAction",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_PipelineInvokeAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := p.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineInvokeAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := p.validateBoundParameters(scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineInvokeAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := p.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineInvokeAction) VariableExpression(variableName *string) *string {
	if err := p.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

