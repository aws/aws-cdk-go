package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline invoke action that uses AWS InspectorScan for ECR images.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//   var repository iRepository
//
//
//   scanOutput := codepipeline.NewArtifact()
//   scanAction := codepipeline_actions.NewInspectorEcrImageScanAction(&InspectorEcrImageScanActionProps{
//   	ActionName: jsii.String("InspectorEcrImageScanAction"),
//   	Output: scanOutput,
//   	Repository: repository,
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Scan"),
//   	Actions: []iAction{
//   		scanAction,
//   	},
//   })
//
type InspectorEcrImageScanAction interface {
	InspectorScanActionBase
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	Props() *InspectorScanActionBaseProps
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *InspectorScanVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	RenderActionConfiguration() *map[string]interface{}
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for InspectorEcrImageScanAction
type jsiiProxy_InspectorEcrImageScanAction struct {
	jsiiProxy_InspectorScanActionBase
}

func (j *jsiiProxy_InspectorEcrImageScanAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorEcrImageScanAction) Props() *InspectorScanActionBaseProps {
	var returns *InspectorScanActionBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorEcrImageScanAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorEcrImageScanAction) Variables() *InspectorScanVariables {
	var returns *InspectorScanVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewInspectorEcrImageScanAction(props *InspectorEcrImageScanActionProps) InspectorEcrImageScanAction {
	_init_.Initialize()

	if err := validateNewInspectorEcrImageScanActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InspectorEcrImageScanAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.InspectorEcrImageScanAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewInspectorEcrImageScanAction_Override(i InspectorEcrImageScanAction, props *InspectorEcrImageScanActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.InspectorEcrImageScanAction",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_InspectorEcrImageScanAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := i.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorEcrImageScanAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := i.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		i,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorEcrImageScanAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := i.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorEcrImageScanAction) RenderActionConfiguration() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"renderActionConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorEcrImageScanAction) VariableExpression(variableName *string) *string {
	if err := i.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

