package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline invoke action that uses AWS InspectorScan for source code.
//
// Example:
//   var pipeline pipeline
//
//
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
//   	ActionName: jsii.String("CodeStarConnectionsSourceAction"),
//   	Output: sourceOutput,
//   	ConnectionArn: jsii.String("your-connection-arn"),
//   	Owner: jsii.String("your-owner"),
//   	Repo: jsii.String("your-repo"),
//   })
//
//   scanOutput := codepipeline.NewArtifact()
//   scanAction := codepipeline_actions.NewInspectorSourceCodeScanAction(&InspectorSourceCodeScanActionProps{
//   	ActionName: jsii.String("InspectorSourceCodeScanAction"),
//   	Input: sourceOutput,
//   	Output: scanOutput,
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		sourceAction,
//   	},
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Scan"),
//   	Actions: []*iAction{
//   		scanAction,
//   	},
//   })
//
type InspectorSourceCodeScanAction interface {
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

// The jsii proxy struct for InspectorSourceCodeScanAction
type jsiiProxy_InspectorSourceCodeScanAction struct {
	jsiiProxy_InspectorScanActionBase
}

func (j *jsiiProxy_InspectorSourceCodeScanAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorSourceCodeScanAction) Props() *InspectorScanActionBaseProps {
	var returns *InspectorScanActionBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorSourceCodeScanAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorSourceCodeScanAction) Variables() *InspectorScanVariables {
	var returns *InspectorScanVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewInspectorSourceCodeScanAction(props *InspectorSourceCodeScanActionProps) InspectorSourceCodeScanAction {
	_init_.Initialize()

	if err := validateNewInspectorSourceCodeScanActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InspectorSourceCodeScanAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.InspectorSourceCodeScanAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewInspectorSourceCodeScanAction_Override(i InspectorSourceCodeScanAction, props *InspectorSourceCodeScanActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.InspectorSourceCodeScanAction",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_InspectorSourceCodeScanAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (i *jsiiProxy_InspectorSourceCodeScanAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (i *jsiiProxy_InspectorSourceCodeScanAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (i *jsiiProxy_InspectorSourceCodeScanAction) RenderActionConfiguration() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"renderActionConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorSourceCodeScanAction) VariableExpression(variableName *string) *string {
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

