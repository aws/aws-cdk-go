package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Deploys the skill to Alexa.
//
// Example:
//   // Read the secrets from ParameterStore
//   clientId := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientId"))
//   clientSecret := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientSecret"))
//   refreshToken := awscdk.SecretValue_SecretsManager(jsii.String("AlexaRefreshToken"))
//
//   // Add deploy action
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewAlexaSkillDeployAction(&AlexaSkillDeployActionProps{
//   	ActionName: jsii.String("DeploySkill"),
//   	RunOrder: jsii.Number(1),
//   	Input: sourceOutput,
//   	ClientId: clientId.ToString(),
//   	ClientSecret: clientSecret,
//   	RefreshToken: refreshToken,
//   	SkillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
//   })
//
// Experimental.
type AlexaSkillDeployAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	// Experimental.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	// Experimental.
	Bound(_scope awscdk.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Experimental.
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


// Experimental.
func NewAlexaSkillDeployAction(props *AlexaSkillDeployActionProps) AlexaSkillDeployAction {
	_init_.Initialize()

	if err := validateNewAlexaSkillDeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlexaSkillDeployAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.AlexaSkillDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlexaSkillDeployAction_Override(a AlexaSkillDeployAction, props *AlexaSkillDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.AlexaSkillDeployAction",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AlexaSkillDeployAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := a.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlexaSkillDeployAction) Bound(_scope awscdk.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := a.validateBoundParameters(_scope, _stage, _options); err != nil {
		panic(err)
	}
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
	if err := a.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
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
	if err := a.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

