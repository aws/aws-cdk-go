package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline Action to deploy an ECS Service.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var service fargateService
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   buildOutput := codepipeline.NewArtifact()
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		codepipeline_actions.NewEcsDeployAction(&ecsDeployActionProps{
//   			actionName: jsii.String("DeployAction"),
//   			service: service,
//   			// if your file is called imagedefinitions.json,
//   			// use the `input` property,
//   			// and leave out the `imageFile` property
//   			input: buildOutput,
//   			// if your file name is _not_ imagedefinitions.json,
//   			// use the `imageFile` property,
//   			// and leave out the `input` property
//   			imageFile: buildOutput.atPath(jsii.String("imageDef.json")),
//   			deploymentTimeout: awscdk.Duration.minutes(jsii.Number(60)),
//   		}),
//   	},
//   })
//
type EcsDeployAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for EcsDeployAction
type jsiiProxy_EcsDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_EcsDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewEcsDeployAction(props *EcsDeployActionProps) EcsDeployAction {
	_init_.Initialize()

	if err := validateNewEcsDeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcsDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcsDeployAction_Override(e EcsDeployAction, props *EcsDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcsDeployAction",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcsDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (e *jsiiProxy_EcsDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBoundParameters(_scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (e *jsiiProxy_EcsDeployAction) VariableExpression(variableName *string) *string {
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

