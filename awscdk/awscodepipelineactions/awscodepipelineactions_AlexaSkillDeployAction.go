package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Deploys the skill to Alexa.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Read the secrets from ParameterStore
//   clientId := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientId"))
//   clientSecret := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientSecret"))
//   refreshToken := awscdk.SecretValue.secretsManager(jsii.String("AlexaRefreshToken"))
//
//   // Add deploy action
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewAlexaSkillDeployAction(&alexaSkillDeployActionProps{
//   	actionName: jsii.String("DeploySkill"),
//   	runOrder: jsii.Number(1),
//   	input: sourceOutput,
//   	clientId: clientId.toString(),
//   	clientSecret: clientSecret,
//   	refreshToken: refreshToken,
//   	skillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
//   })
//
type AlexaSkillDeployAction interface {
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
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for AlexaSkillDeployAction
type jsiiProxy_AlexaSkillDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_AlexaSkillDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlexaSkillDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewAlexaSkillDeployAction(props *AlexaSkillDeployActionProps) AlexaSkillDeployAction {
	_init_.Initialize()

	j := jsiiProxy_AlexaSkillDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.AlexaSkillDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAlexaSkillDeployAction_Override(a AlexaSkillDeployAction, props *AlexaSkillDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.AlexaSkillDeployAction",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AlexaSkillDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlexaSkillDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		a,
		"bound",
		[]interface{}{_scope, _stage, _options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlexaSkillDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlexaSkillDeployAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

