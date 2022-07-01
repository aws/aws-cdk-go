package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Low-level class for generic CodePipeline Actions.
//
// If you're implementing your own IAction,
// prefer to use the Action class from the codepipeline module.
type Action interface {
	awscodepipeline.Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for Action
type jsiiProxy_Action struct {
	internal.Type__awscodepipelineAction
}

func (j *jsiiProxy_Action) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewAction_Override(a Action, actionProperties *awscodepipeline.ActionProperties) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Action",
		[]interface{}{actionProperties},
		a,
	)
}

func (a *jsiiProxy_Action) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		a,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

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

// Construction properties of the {@link AlexaSkillDeployAction Alexa deploy Action}.
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
type AlexaSkillDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The client id of the developer console token.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret of the developer console token.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The source artifact containing the voice model and skill manifest.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The refresh token of the developer console token.
	RefreshToken awscdk.SecretValue `field:"required" json:"refreshToken" yaml:"refreshToken"`
	// The Alexa skill id.
	SkillId *string `field:"required" json:"skillId" yaml:"skillId"`
	// An optional artifact containing overrides for the skill manifest.
	ParameterOverridesArtifact awscodepipeline.Artifact `field:"optional" json:"parameterOverridesArtifact" yaml:"parameterOverridesArtifact"`
}

type BaseJenkinsProvider interface {
	constructs.Construct
	IJenkinsProvider
	// The tree node.
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BaseJenkinsProvider
type jsiiProxy_BaseJenkinsProvider struct {
	internal.Type__constructsConstruct
	jsiiProxy_IJenkinsProvider
}

func (j *jsiiProxy_BaseJenkinsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseJenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewBaseJenkinsProvider_Override(b BaseJenkinsProvider, scope constructs.Construct, id *string, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.BaseJenkinsProvider",
		[]interface{}{scope, id, version},
		b,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func BaseJenkinsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.BaseJenkinsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseJenkinsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Used for HTTP cache-control header, which influences downstream caches.
//
// Use the provided static factory methods to construct instances of this class.
// Used in the {@link S3DeployActionProps.cacheControl} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheControl := awscdk.Aws_codepipeline_actions.cacheControl.fromString(jsii.String("s"))
//
// See: https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
//
type CacheControl interface {
	// the actual text value of the created directive.
	Value() *string
	SetValue(val *string)
}

// The jsii proxy struct for CacheControl
type jsiiProxy_CacheControl struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheControl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_CacheControl) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Allows you to create an arbitrary cache control directive, in case our support is missing a method for a particular directive.
func CacheControl_FromString(s *string) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// The 'max-age' cache control directive.
func CacheControl_MaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"maxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// The 'must-revalidate' cache control directive.
func CacheControl_MustRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"mustRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'no-cache' cache control directive.
func CacheControl_NoCache() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"noCache",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'no-transform' cache control directive.
func CacheControl_NoTransform() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"noTransform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'proxy-revalidate' cache control directive.
func CacheControl_ProxyRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"proxyRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'private' cache control directive.
func CacheControl_SetPrivate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"setPrivate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'public' cache control directive.
func CacheControl_SetPublic() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"setPublic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 's-max-age' cache control directive.
func CacheControl_SMaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"sMaxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// CodePipeline action to prepare a change set.
//
// Creates the change set if it doesn't exist based on the stack name and template that you submit.
// If the change set exists, AWS CloudFormation deletes it, and then creates a new one.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationCreateReplaceChangeSetAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationCreateReplaceChangeSetAction
type jsiiProxy_CloudFormationCreateReplaceChangeSetAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationCreateReplaceChangeSetAction(props *CloudFormationCreateReplaceChangeSetActionProps) CloudFormationCreateReplaceChangeSetAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationCreateReplaceChangeSetAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationCreateReplaceChangeSetAction_Override(c CloudFormationCreateReplaceChangeSetAction, props *CloudFormationCreateReplaceChangeSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationCreateReplaceChangeSetAction.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationCreateReplaceChangeSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to grant full permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have full (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// Name of the change set to create or update.
	ChangeSetName *string `field:"required" json:"changeSetName" yaml:"changeSetName"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Input artifact with the ChangeSet's CloudFormation template.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	DeploymentRole awsiam.IRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// The list of additional input Artifacts for this Action.
	//
	// This is especially useful when used in conjunction with the `parameterOverrides` property.
	// For example, if you have:
	//
	//    parameterOverrides: {
	//      'Param1': action1.outputArtifact.bucketName,
	//      'Param2': action2.outputArtifact.objectKey,
	//    }
	//
	// , if the output Artifacts of `action1` and `action2` were not used to
	// set either the `templateConfiguration` or the `templatePath` properties,
	// you need to make sure to include them in the `extraInputs` -
	// otherwise, you'll get an "unrecognized Artifact" error during your Pipeline's execution.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Additional template parameters.
	//
	// Template parameters specified here take precedence over template parameters
	// found in the artifact specified by the `templateConfiguration` property.
	//
	// We recommend that you use the template configuration file to specify
	// most of your parameter values. Use parameter overrides to specify only
	// dynamic parameter values (values that are unknown until you run the
	// pipeline).
	//
	// All parameter names must be present in the stack template.
	//
	// Note: the entire object cannot be more than 1kB.
	ParameterOverrides *map[string]interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	TemplateConfiguration awscodepipeline.ArtifactPath `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

// CodePipeline action to deploy a stack.
//
// Creates the stack if the specified stack doesn't exist. If the stack exists,
// AWS CloudFormation updates the stack. Use this action to update existing
// stacks.
//
// AWS CodePipeline won't replace the stack, and will fail deployment if the
// stack is in a failed state. Use `ReplaceOnFailure` for an action that
// will delete and recreate the stack to try and recover from failed states.
//
// Use this action to automatically replace failed stacks without recovering or
// troubleshooting them. You would typically choose this mode for testing.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // in stack for account 123456789012...
//   var otherAccountStack stack
//
//   actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &roleProps{
//   	assumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
//   	// the role has to have a physical name set
//   	roleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   })
//
//   // in the pipeline stack...
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   	actionName: jsii.String("CloudFormationCreateUpdate"),
//   	stackName: jsii.String("MyStackName"),
//   	adminPermissions: jsii.Boolean(true),
//   	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
//   	role: actionRole,
//   })
//
type CloudFormationCreateUpdateStackAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationCreateUpdateStackAction
type jsiiProxy_CloudFormationCreateUpdateStackAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationCreateUpdateStackAction(props *CloudFormationCreateUpdateStackActionProps) CloudFormationCreateUpdateStackAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationCreateUpdateStackAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationCreateUpdateStackAction_Override(c CloudFormationCreateUpdateStackAction, props *CloudFormationCreateUpdateStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationCreateUpdateStackAction.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // in stack for account 123456789012...
//   var otherAccountStack stack
//
//   actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &roleProps{
//   	assumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
//   	// the role has to have a physical name set
//   	roleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   })
//
//   // in the pipeline stack...
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   	actionName: jsii.String("CloudFormationCreateUpdate"),
//   	stackName: jsii.String("MyStackName"),
//   	adminPermissions: jsii.Boolean(true),
//   	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
//   	role: actionRole,
//   })
//
type CloudFormationCreateUpdateStackActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to grant full permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have full (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Input artifact with the CloudFormation template to deploy.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	DeploymentRole awsiam.IRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// The list of additional input Artifacts for this Action.
	//
	// This is especially useful when used in conjunction with the `parameterOverrides` property.
	// For example, if you have:
	//
	//    parameterOverrides: {
	//      'Param1': action1.outputArtifact.bucketName,
	//      'Param2': action2.outputArtifact.objectKey,
	//    }
	//
	// , if the output Artifacts of `action1` and `action2` were not used to
	// set either the `templateConfiguration` or the `templatePath` properties,
	// you need to make sure to include them in the `extraInputs` -
	// otherwise, you'll get an "unrecognized Artifact" error during your Pipeline's execution.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Additional template parameters.
	//
	// Template parameters specified here take precedence over template parameters
	// found in the artifact specified by the `templateConfiguration` property.
	//
	// We recommend that you use the template configuration file to specify
	// most of your parameter values. Use parameter overrides to specify only
	// dynamic parameter values (values that are unknown until you run the
	// pipeline).
	//
	// All parameter names must be present in the stack template.
	//
	// Note: the entire object cannot be more than 1kB.
	ParameterOverrides *map[string]interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Replace the stack if it's in a failed state.
	//
	// If this is set to true and the stack is in a failed state (one of
	// ROLLBACK_COMPLETE, ROLLBACK_FAILED, CREATE_FAILED, DELETE_FAILED, or
	// UPDATE_ROLLBACK_FAILED), AWS CloudFormation deletes the stack and then
	// creates a new stack.
	//
	// If this is not set to true and the stack is in a failed state,
	// the deployment fails.
	ReplaceOnFailure *bool `field:"optional" json:"replaceOnFailure" yaml:"replaceOnFailure"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	TemplateConfiguration awscodepipeline.ArtifactPath `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

// CodePipeline action to delete a stack.
//
// Deletes a stack. If you specify a stack that doesn't exist, the action completes successfully
// without deleting a stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var parameterOverrides interface{}
//   var role role
//
//   cloudFormationDeleteStackAction := awscdk.Aws_codepipeline_actions.NewCloudFormationDeleteStackAction(&cloudFormationDeleteStackActionProps{
//   	actionName: jsii.String("actionName"),
//   	adminPermissions: jsii.Boolean(false),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	cfnCapabilities: []cfnCapabilities{
//   		cdk.*cfnCapabilities_NONE,
//   	},
//   	deploymentRole: role,
//   	extraInputs: []*artifact{
//   		artifact,
//   	},
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	parameterOverrides: map[string]interface{}{
//   		"parameterOverridesKey": parameterOverrides,
//   	},
//   	region: jsii.String("region"),
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	templateConfiguration: artifactPath,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   })
//
type CloudFormationDeleteStackAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationDeleteStackAction
type jsiiProxy_CloudFormationDeleteStackAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationDeleteStackAction(props *CloudFormationDeleteStackActionProps) CloudFormationDeleteStackAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationDeleteStackAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationDeleteStackAction_Override(c CloudFormationDeleteStackAction, props *CloudFormationDeleteStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationDeleteStackAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var parameterOverrides interface{}
//   var role role
//
//   cloudFormationDeleteStackActionProps := &cloudFormationDeleteStackActionProps{
//   	actionName: jsii.String("actionName"),
//   	adminPermissions: jsii.Boolean(false),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	cfnCapabilities: []cfnCapabilities{
//   		cdk.*cfnCapabilities_NONE,
//   	},
//   	deploymentRole: role,
//   	extraInputs: []*artifact{
//   		artifact,
//   	},
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	parameterOverrides: map[string]interface{}{
//   		"parameterOverridesKey": parameterOverrides,
//   	},
//   	region: jsii.String("region"),
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	templateConfiguration: artifactPath,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type CloudFormationDeleteStackActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to grant full permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have full (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	DeploymentRole awsiam.IRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// The list of additional input Artifacts for this Action.
	//
	// This is especially useful when used in conjunction with the `parameterOverrides` property.
	// For example, if you have:
	//
	//    parameterOverrides: {
	//      'Param1': action1.outputArtifact.bucketName,
	//      'Param2': action2.outputArtifact.objectKey,
	//    }
	//
	// , if the output Artifacts of `action1` and `action2` were not used to
	// set either the `templateConfiguration` or the `templatePath` properties,
	// you need to make sure to include them in the `extraInputs` -
	// otherwise, you'll get an "unrecognized Artifact" error during your Pipeline's execution.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Additional template parameters.
	//
	// Template parameters specified here take precedence over template parameters
	// found in the artifact specified by the `templateConfiguration` property.
	//
	// We recommend that you use the template configuration file to specify
	// most of your parameter values. Use parameter overrides to specify only
	// dynamic parameter values (values that are unknown until you run the
	// pipeline).
	//
	// All parameter names must be present in the stack template.
	//
	// Note: the entire object cannot be more than 1kB.
	ParameterOverrides *map[string]interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	TemplateConfiguration awscodepipeline.ArtifactPath `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

// CodePipeline action to create/update Stack Instances of a StackSet.
//
// After the initial creation of a stack set, you can add new stack instances by
// using CloudFormationStackInstances. Template parameter values can be
// overridden at the stack instance level during create or update stack set
// instance operations.
//
// Each stack set has one template and set of template parameters. When you
// update the template or template parameters, you update them for the entire
// set. Then all instance statuses are set to OUTDATED until the changes are
// deployed to that instance.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackInstancesAction interface {
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
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationDeployStackInstancesAction
type jsiiProxy_CloudFormationDeployStackInstancesAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationDeployStackInstancesAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeployStackInstancesAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationDeployStackInstancesAction(props *CloudFormationDeployStackInstancesActionProps) CloudFormationDeployStackInstancesAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationDeployStackInstancesAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackInstancesAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationDeployStackInstancesAction_Override(c CloudFormationDeployStackInstancesAction, props *CloudFormationDeployStackInstancesActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackInstancesAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationDeployStackInstancesAction.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackInstancesActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
	// Specify where to create or update Stack Instances.
	//
	// You can specify either AWS Accounts Ids or AWS Organizations Organizational Units.
	StackInstances StackInstances `field:"required" json:"stackInstances" yaml:"stackInstances"`
	// The name of the StackSet we are adding instances to.
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// Parameter values that only apply to the current Stack Instances.
	//
	// These parameters are shared between all instances added by this action.
	ParameterOverrides StackSetParameters `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
}

// CodePipeline action to deploy a stackset.
//
// CodePipeline offers the ability to perform AWS CloudFormation StackSets
// operations as part of your CI/CD process. You use a stack set to create
// stacks in AWS accounts across AWS Regions by using a single AWS
// CloudFormation template. All the resources included in each stack are defined
// by the stack set’s AWS CloudFormation template. When you create the stack
// set, you specify the template to use, as well as any parameters and
// capabilities that the template requires.
//
// For more information about concepts for AWS CloudFormation StackSets, see
// [StackSets
// concepts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html)
// in the AWS CloudFormation User Guide.
//
// If you use this action to make an update that includes adding stack
// instances, the new instances are deployed first and the update is completed
// last. The new instances first receive the old version, and then the update is
// applied to all instances.
//
// As a best practice, you should construct your pipeline so that the stack set
// is created and initially deploys to a subset or a single instance. After you
// test your deployment and view the generated stack set, then add the
// CloudFormationStackInstances action so that the remaining instances are
// created and updated.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackSetAction interface {
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
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationDeployStackSetAction
type jsiiProxy_CloudFormationDeployStackSetAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationDeployStackSetAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeployStackSetAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationDeployStackSetAction(props *CloudFormationDeployStackSetActionProps) CloudFormationDeployStackSetAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationDeployStackSetAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackSetAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationDeployStackSetAction_Override(c CloudFormationDeployStackSetAction, props *CloudFormationDeployStackSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackSetAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationDeployStackSetAction.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
	// The name to associate with the stack set.
	//
	// This name must be unique in the Region where it is created.
	//
	// The name may only contain alphanumeric and hyphen characters. It must begin with an alphabetic character and be 128 characters or fewer.
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// The location of the template that defines the resources in the stack set.
	//
	// This must point to a template with a maximum size of 460,800 bytes.
	//
	// Enter the path to the source artifact name and template file.
	Template StackSetTemplate `field:"required" json:"template" yaml:"template"`
	// Indicates that the template can create and update resources, depending on the types of resources in the template.
	//
	// You must use this property if you have IAM resources in your stack template or you create a stack directly from a template containing macros.
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// Determines how IAM roles are created and managed.
	//
	// The choices are:
	//
	// - Self Managed: you create IAM roles with the required permissions
	//    in the administration account and all target accounts.
	// - Service Managed: only available if the account and target accounts
	//    are part of an AWS Organization. The necessary roles will be created
	//    for you.
	//
	// If you want to deploy to all accounts that are a member of AWS
	// Organizations Organizational Units (OUs), you must select Service Managed
	// permissions.
	//
	// Note: This parameter can only be changed when no stack instances exist in
	// the stack set.
	DeploymentModel StackSetDeploymentModel `field:"optional" json:"deploymentModel" yaml:"deploymentModel"`
	// A description of the stack set.
	//
	// You can use this to describe the stack set’s purpose or other relevant information.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The template parameters for your stack set.
	//
	// These parameters are shared between all instances of the stack set.
	Parameters StackSetParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// Specify where to create or update Stack Instances.
	//
	// You can specify either AWS Accounts Ids or AWS Organizations Organizational Units.
	StackInstances StackInstances `field:"optional" json:"stackInstances" yaml:"stackInstances"`
}

// CodePipeline action to execute a prepared change set.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationExecuteChangeSetAction interface {
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
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationExecuteChangeSetAction
type jsiiProxy_CloudFormationExecuteChangeSetAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationExecuteChangeSetAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationExecuteChangeSetAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationExecuteChangeSetAction(props *CloudFormationExecuteChangeSetActionProps) CloudFormationExecuteChangeSetAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationExecuteChangeSetAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationExecuteChangeSetAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationExecuteChangeSetAction_Override(c CloudFormationExecuteChangeSetAction, props *CloudFormationExecuteChangeSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationExecuteChangeSetAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Properties for the CloudFormationExecuteChangeSetAction.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationExecuteChangeSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Name of the change set to execute.
	ChangeSetName *string `field:"required" json:"changeSetName" yaml:"changeSetName"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

// CodePipeline build action that uses AWS CodeBuild.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &pipelineProjectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &buildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.account, distribution.distributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	resources: []*string{
//   		distributionArn,
//   	},
//   	actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		&stageProps{
//   			stageName: jsii.String("Deploy"),
//   			actions: []iAction{
//   				codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   					actionName: jsii.String("S3Deploy"),
//   					bucket: deployBucket,
//   					input: deployInput,
//   					runOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   					actionName: jsii.String("InvalidateCache"),
//   					project: invalidateBuildProject,
//   					input: deployInput,
//   					runOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
type CodeBuildAction interface {
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
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Reference a CodePipeline variable defined by the CodeBuild project this action points to.
	//
	// Variables in CodeBuild actions are defined using the 'exported-variables' subsection of the 'env'
	// section of the buildspec.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-syntax
	//
	Variable(variableName *string) *string
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeBuildAction
type jsiiProxy_CodeBuildAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeBuildAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeBuildAction(props *CodeBuildActionProps) CodeBuildAction {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeBuildAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeBuildAction_Override(c CodeBuildAction, props *CodeBuildActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeBuildAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) Variable(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link CodeBuildAction CodeBuild build CodePipeline action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &pipelineProjectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &buildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.account, distribution.distributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	resources: []*string{
//   		distributionArn,
//   	},
//   	actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		&stageProps{
//   			stageName: jsii.String("Deploy"),
//   			actions: []iAction{
//   				codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   					actionName: jsii.String("S3Deploy"),
//   					bucket: deployBucket,
//   					input: deployInput,
//   					runOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   					actionName: jsii.String("InvalidateCache"),
//   					project: invalidateBuildProject,
//   					input: deployInput,
//   					runOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
type CodeBuildActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The source to use as input for this action.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The action's Project.
	Project awscodebuild.IProject `field:"required" json:"project" yaml:"project"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	CheckSecretsInPlainTextEnvVariables *bool `field:"optional" json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Combine the build artifacts for a batch builds.
	//
	// Enabling this will combine the build artifacts into the same location for batch builds.
	// If `executeBatchBuild` is not set to `true`, this property is ignored.
	CombineBatchBuildArtifacts *bool `field:"optional" json:"combineBatchBuildArtifacts" yaml:"combineBatchBuildArtifacts"`
	// The environment variables to pass to the CodeBuild project when this action executes.
	//
	// If a variable with the same name was set both on the project level, and here,
	// this value will take precedence.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Trigger a batch build.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	ExecuteBatchBuild *bool `field:"optional" json:"executeBatchBuild" yaml:"executeBatchBuild"`
	// The list of additional input Artifacts for this action.
	//
	// The directories the additional inputs will be available at are available
	// during the project's build in the CODEBUILD_SRC_DIR_<artifact-name> environment variables.
	// The project's build always starts in the directory with the primary input artifact checked out,
	// the one pointed to by the {@link input} property.
	// For more information,
	// see https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html .
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The list of output Artifacts for this action.
	//
	// **Note**: if you specify more than one output Artifact here,
	// you cannot use the primary 'artifacts' section of the buildspec;
	// you have to use the 'secondary-artifacts' section instead.
	// See https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	// for details.
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	// The type of the action that determines its CodePipeline Category - Build, or Test.
	Type CodeBuildActionType `field:"optional" json:"type" yaml:"type"`
}

// The type of the CodeBuild action that determines its CodePipeline Category - Build, or Test.
//
// The default is Build.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var project pipelineProject
//
//   sourceOutput := codepipeline.NewArtifact()
//   testAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("IntegrationTest"),
//   	project: project,
//   	input: sourceOutput,
//   	type: codepipeline_actions.codeBuildActionType_TEST,
//   })
//
type CodeBuildActionType string

const (
	// The action will have the Build Category.
	//
	// This is the default.
	CodeBuildActionType_BUILD CodeBuildActionType = "BUILD"
	// The action will have the Test Category.
	CodeBuildActionType_TEST CodeBuildActionType = "TEST"
)

// CodePipeline Source that is provided by an AWS CodeCommit repository.
//
// If the CodeCommit repository is in a different account, you must use
// `CodeCommitTrigger.EVENTS` to trigger the pipeline.
//
// (That is because the Pipeline structure normally only has a `RepositoryName`
// field, and that is not enough for the pipeline to locate the repository's
// source account. However, if the pipeline is triggered via an EventBridge
// event, the event itself has the full repository ARN in there, allowing the
// pipeline to locate the repository).
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CodeCommitSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *CodeCommitSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeCommitSourceAction
type jsiiProxy_CodeCommitSourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeCommitSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitSourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitSourceAction) Variables() *CodeCommitSourceVariables {
	var returns *CodeCommitSourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewCodeCommitSourceAction(props *CodeCommitSourceActionProps) CodeCommitSourceAction {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeCommitSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeCommitSourceAction_Override(c CodeCommitSourceAction, props *CodeCommitSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeCommitSourceAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeCommitSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeCommitSourceAction) Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeCommitSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeCommitSourceAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link CodeCommitSourceAction CodeCommit source CodePipeline Action}.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CodeCommitSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The CodeCommit repository.
	Repository awscodecommit.IRepository `field:"required" json:"repository" yaml:"repository"`
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	Trigger CodeCommitTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

// The CodePipeline variables emitted by the CodeCommit source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeCommitSourceVariables := &codeCommitSourceVariables{
//   	authorDate: jsii.String("authorDate"),
//   	branchName: jsii.String("branchName"),
//   	commitId: jsii.String("commitId"),
//   	commitMessage: jsii.String("commitMessage"),
//   	committerDate: jsii.String("committerDate"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type CodeCommitSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	AuthorDate *string `field:"required" json:"authorDate" yaml:"authorDate"`
	// The name of the branch this action tracks.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The message of the currently last commit on the tracked branch.
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// The date the currently last commit on the tracked branch was committed, in ISO-8601 format.
	CommitterDate *string `field:"required" json:"committerDate" yaml:"committerDate"`
	// The name of the repository this action points to.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

// How should the CodeCommit Action detect changes.
//
// This is the type of the {@link CodeCommitSourceAction.trigger} property.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CodeCommitTrigger string

const (
	// The Action will never detect changes - the Pipeline it's part of will only begin a run when explicitly started.
	CodeCommitTrigger_NONE CodeCommitTrigger = "NONE"
	// CodePipeline will poll the repository to detect changes.
	CodeCommitTrigger_POLL CodeCommitTrigger = "POLL"
	// CodePipeline will use CloudWatch Events to be notified of changes.
	//
	// This is the default method of detecting changes.
	CodeCommitTrigger_EVENTS CodeCommitTrigger = "EVENTS"
)

// Configuration for replacing a placeholder string in the ECS task definition template file with an image URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   codeDeployEcsContainerImageInput := &codeDeployEcsContainerImageInput{
//   	input: artifact,
//
//   	// the properties below are optional
//   	taskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   }
//
type CodeDeployEcsContainerImageInput struct {
	// The artifact that contains an `imageDetails.json` file with the image URI.
	//
	// The artifact's `imageDetails.json` file must be a JSON file containing an
	// `ImageURI` property.  For example:
	// `{ "ImageURI": "ACCOUNTID.dkr.ecr.us-west-2.amazonaws.com/dk-image-repo@sha256:example3" }`
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The placeholder string in the ECS task definition template file that will be replaced with the image URI.
	//
	// The placeholder string must be surrounded by angle brackets in the template file.
	// For example, if the task definition template file contains a placeholder like
	// `"image": "<PLACEHOLDER>"`, then the `taskDefinitionPlaceholder` value should
	// be `PLACEHOLDER`.
	TaskDefinitionPlaceholder *string `field:"optional" json:"taskDefinitionPlaceholder" yaml:"taskDefinitionPlaceholder"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var ecsDeploymentGroup iEcsDeploymentGroup
//   var role role
//
//   codeDeployEcsDeployAction := awscdk.Aws_codepipeline_actions.NewCodeDeployEcsDeployAction(&codeDeployEcsDeployActionProps{
//   	actionName: jsii.String("actionName"),
//   	deploymentGroup: ecsDeploymentGroup,
//
//   	// the properties below are optional
//   	appSpecTemplateFile: artifactPath,
//   	appSpecTemplateInput: artifact,
//   	containerImageInputs: []codeDeployEcsContainerImageInput{
//   		&codeDeployEcsContainerImageInput{
//   			input: artifact,
//
//   			// the properties below are optional
//   			taskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   		},
//   	},
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	taskDefinitionTemplateFile: artifactPath,
//   	taskDefinitionTemplateInput: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   })
//
type CodeDeployEcsDeployAction interface {
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

// The jsii proxy struct for CodeDeployEcsDeployAction
type jsiiProxy_CodeDeployEcsDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeDeployEcsDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeDeployEcsDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeDeployEcsDeployAction(props *CodeDeployEcsDeployActionProps) CodeDeployEcsDeployAction {
	_init_.Initialize()

	j := jsiiProxy_CodeDeployEcsDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeDeployEcsDeployAction_Override(c CodeDeployEcsDeployAction, props *CodeDeployEcsDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link CodeDeployEcsDeployAction CodeDeploy ECS deploy CodePipeline Action}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var ecsDeploymentGroup iEcsDeploymentGroup
//   var role role
//
//   codeDeployEcsDeployActionProps := &codeDeployEcsDeployActionProps{
//   	actionName: jsii.String("actionName"),
//   	deploymentGroup: ecsDeploymentGroup,
//
//   	// the properties below are optional
//   	appSpecTemplateFile: artifactPath,
//   	appSpecTemplateInput: artifact,
//   	containerImageInputs: []codeDeployEcsContainerImageInput{
//   		&codeDeployEcsContainerImageInput{
//   			input: artifact,
//
//   			// the properties below are optional
//   			taskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   		},
//   	},
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	taskDefinitionTemplateFile: artifactPath,
//   	taskDefinitionTemplateInput: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type CodeDeployEcsDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The CodeDeploy ECS Deployment Group to deploy to.
	DeploymentGroup awscodedeploy.IEcsDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The name of the CodeDeploy AppSpec file.
	//
	// During deployment, a new task definition will be registered
	// with ECS, and the new task definition ID will be inserted into
	// the CodeDeploy AppSpec file.  The AppSpec file contents will be
	// provided to CodeDeploy for the deployment.
	//
	// Use this property if you want to use a different name for this file than the default 'appspec.yaml'.
	// If you use this property, you don't need to specify the `appSpecTemplateInput` property.
	AppSpecTemplateFile awscodepipeline.ArtifactPath `field:"optional" json:"appSpecTemplateFile" yaml:"appSpecTemplateFile"`
	// The artifact containing the CodeDeploy AppSpec file.
	//
	// During deployment, a new task definition will be registered
	// with ECS, and the new task definition ID will be inserted into
	// the CodeDeploy AppSpec file.  The AppSpec file contents will be
	// provided to CodeDeploy for the deployment.
	//
	// If you use this property, it's assumed the file is called 'appspec.yaml'.
	// If your AppSpec file uses a different filename, leave this property empty,
	// and use the `appSpecTemplateFile` property instead.
	AppSpecTemplateInput awscodepipeline.Artifact `field:"optional" json:"appSpecTemplateInput" yaml:"appSpecTemplateInput"`
	// Configuration for dynamically updated images in the task definition.
	//
	// Provide pairs of an image details input artifact and a placeholder string
	// that will be used to dynamically update the ECS task definition template
	// file prior to deployment. A maximum of 4 images can be given.
	ContainerImageInputs *[]*CodeDeployEcsContainerImageInput `field:"optional" json:"containerImageInputs" yaml:"containerImageInputs"`
	// The name of the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// Use this property if you want to use a different name for this file than the default 'taskdef.json'.
	// If you use this property, you don't need to specify the `taskDefinitionTemplateInput` property.
	TaskDefinitionTemplateFile awscodepipeline.ArtifactPath `field:"optional" json:"taskDefinitionTemplateFile" yaml:"taskDefinitionTemplateFile"`
	// The artifact containing the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// If you use this property, it's assumed the file is called 'taskdef.json'.
	// If your task definition template uses a different filename, leave this property empty,
	// and use the `taskDefinitionTemplateFile` property instead.
	TaskDefinitionTemplateInput awscodepipeline.Artifact `field:"optional" json:"taskDefinitionTemplateInput" yaml:"taskDefinitionTemplateInput"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var deploymentGroup serverDeploymentGroup
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
//   	pipelineName: jsii.String("MyPipeline"),
//   })
//
//   // add the source and build Stages to the Pipeline...
//   buildOutput := codepipeline.NewArtifact()
//   deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&codeDeployServerDeployActionProps{
//   	actionName: jsii.String("CodeDeploy"),
//   	input: buildOutput,
//   	deploymentGroup: deploymentGroup,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type CodeDeployServerDeployAction interface {
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

// The jsii proxy struct for CodeDeployServerDeployAction
type jsiiProxy_CodeDeployServerDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeDeployServerDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeDeployServerDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeDeployServerDeployAction(props *CodeDeployServerDeployActionProps) CodeDeployServerDeployAction {
	_init_.Initialize()

	j := jsiiProxy_CodeDeployServerDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployServerDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeDeployServerDeployAction_Override(c CodeDeployServerDeployAction, props *CodeDeployServerDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployServerDeployAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeDeployServerDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployServerDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployServerDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployServerDeployAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link CodeDeployServerDeployAction CodeDeploy server deploy CodePipeline Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var deploymentGroup serverDeploymentGroup
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
//   	pipelineName: jsii.String("MyPipeline"),
//   })
//
//   // add the source and build Stages to the Pipeline...
//   buildOutput := codepipeline.NewArtifact()
//   deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&codeDeployServerDeployActionProps{
//   	actionName: jsii.String("CodeDeploy"),
//   	input: buildOutput,
//   	deploymentGroup: deploymentGroup,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type CodeDeployServerDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The CodeDeploy server Deployment Group to deploy to.
	DeploymentGroup awscodedeploy.IServerDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The source to use as input for deployment.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
}

// A CodePipeline source action for the CodeStar Connections source, which allows connecting to GitHub and BitBucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&codeStarConnectionsSourceActionProps{
//   	actionName: jsii.String("BitBucket_Source"),
//   	owner: jsii.String("aws"),
//   	repo: jsii.String("aws-cdk"),
//   	output: sourceOutput,
//   	connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
//   })
//
type CodeStarConnectionsSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *CodeStarSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeStarConnectionsSourceAction
type jsiiProxy_CodeStarConnectionsSourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) Variables() *CodeStarSourceVariables {
	var returns *CodeStarSourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewCodeStarConnectionsSourceAction(props *CodeStarConnectionsSourceActionProps) CodeStarConnectionsSourceAction {
	_init_.Initialize()

	j := jsiiProxy_CodeStarConnectionsSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeStarConnectionsSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeStarConnectionsSourceAction_Override(c CodeStarConnectionsSourceAction, props *CodeStarConnectionsSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeStarConnectionsSourceAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties for {@link CodeStarConnectionsSourceAction}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&codeStarConnectionsSourceActionProps{
//   	actionName: jsii.String("BitBucket_Source"),
//   	owner: jsii.String("aws"),
//   	repo: jsii.String("aws-cdk"),
//   	output: sourceOutput,
//   	connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
//   })
//
type CodeStarConnectionsSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The output artifact that this action produces.
	//
	// Can be used as input for further pipeline actions.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The owning user or organization of the repository.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "aws"
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "aws-cdk"
	//
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to build.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	TriggerOnPush *bool `field:"optional" json:"triggerOnPush" yaml:"triggerOnPush"`
}

// The CodePipeline variables emitted by CodeStar source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeStarSourceVariables := &codeStarSourceVariables{
//   	authorDate: jsii.String("authorDate"),
//   	branchName: jsii.String("branchName"),
//   	commitId: jsii.String("commitId"),
//   	commitMessage: jsii.String("commitMessage"),
//   	connectionArn: jsii.String("connectionArn"),
//   	fullRepositoryName: jsii.String("fullRepositoryName"),
//   }
//
type CodeStarSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	AuthorDate *string `field:"required" json:"authorDate" yaml:"authorDate"`
	// The name of the branch this action tracks.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The message of the currently last commit on the tracked branch.
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// The connection ARN this source uses.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The name of the repository this action points to.
	FullRepositoryName *string `field:"required" json:"fullRepositoryName" yaml:"fullRepositoryName"`
}

// Options in common between both StackSet actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonCloudFormationStackSetOptions := &commonCloudFormationStackSetOptions{
//   	failureTolerancePercentage: jsii.Number(123),
//   	maxAccountConcurrencyPercentage: jsii.Number(123),
//   	stackSetRegion: jsii.String("stackSetRegion"),
//   }
//
type CommonCloudFormationStackSetOptions struct {
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
}

// The ECR Repository source CodePipeline Action.
//
// Will trigger the pipeline as soon as the target tag in the repository
// changes, but only if there is a CloudTrail Trail in the account that
// captures the ECR event.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecrRepository repository
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
//   	actionName: jsii.String("ECR"),
//   	repository: ecrRepository,
//   	imageTag: jsii.String("some-tag"),
//   	 // optional, default: 'latest'
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type EcrSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *EcrSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for EcrSourceAction
type jsiiProxy_EcrSourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_EcrSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrSourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrSourceAction) Variables() *EcrSourceVariables {
	var returns *EcrSourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewEcrSourceAction(props *EcrSourceActionProps) EcrSourceAction {
	_init_.Initialize()

	j := jsiiProxy_EcrSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcrSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcrSourceAction_Override(e EcrSourceAction, props *EcrSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcrSourceAction",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcrSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrSourceAction) Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bound",
		[]interface{}{_scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		e,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrSourceAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of {@link EcrSourceAction}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecrRepository repository
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
//   	actionName: jsii.String("ECR"),
//   	repository: ecrRepository,
//   	imageTag: jsii.String("some-tag"),
//   	 // optional, default: 'latest'
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type EcrSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The repository that will be watched for changes.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The image tag that will be checked for changes.
	//
	// It is not possible to trigger on changes to more than one tag.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

// The CodePipeline variables emitted by the ECR source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrSourceVariables := &ecrSourceVariables{
//   	imageDigest: jsii.String("imageDigest"),
//   	imageTag: jsii.String("imageTag"),
//   	imageUri: jsii.String("imageUri"),
//   	registryId: jsii.String("registryId"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type EcrSourceVariables struct {
	// The digest of the current image, in the form '<digest type>:<digest value>'.
	ImageDigest *string `field:"required" json:"imageDigest" yaml:"imageDigest"`
	// The Docker tag of the current image.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// The full ECR Docker URI of the current image.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The identifier of the registry.
	//
	// In ECR, this is usually the ID of the AWS account owning it.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// The physical name of the repository that this action tracks.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

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
	var returns *string

	_jsii_.Invoke(
		e,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of {@link EcsDeployAction}.
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
type EcsDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The ECS Service to deploy.
	Service awsecs.IBaseService `field:"required" json:"service" yaml:"service"`
	// Timeout for the ECS deployment in minutes.
	//
	// Value must be between 1-60.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-ECS.html
	//
	DeploymentTimeout awscdk.Duration `field:"optional" json:"deploymentTimeout" yaml:"deploymentTimeout"`
	// The name of the JSON image definitions file to use for deployments.
	//
	// The JSON file is a list of objects,
	// each with 2 keys: `name` is the name of the container in the Task Definition,
	// and `imageUri` is the Docker image URI you want to update your service with.
	// Use this property if you want to use a different name for this file than the default 'imagedefinitions.json'.
	// If you use this property, you don't need to specify the `input` property.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions
	//
	ImageFile awscodepipeline.ArtifactPath `field:"optional" json:"imageFile" yaml:"imageFile"`
	// The input artifact that contains the JSON image definitions file to use for deployments.
	//
	// The JSON file is a list of objects,
	// each with 2 keys: `name` is the name of the container in the Task Definition,
	// and `imageUri` is the Docker image URI you want to update your service with.
	// If you use this property, it's assumed the file is called 'imagedefinitions.json'.
	// If your build uses a different file, leave this property empty,
	// and use the `imageFile` property instead.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions
	//
	Input awscodepipeline.Artifact `field:"optional" json:"input" yaml:"input"`
}

// Source that is provided by a GitHub repository.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Read the secret from Secrets Manager
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("GitHub_Source"),
//   	owner: jsii.String("awslabs"),
//   	repo: jsii.String("aws-cdk"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	output: sourceOutput,
//   	branch: jsii.String("develop"),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type GitHubSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *GitHubSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
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


func NewGitHubSourceAction(props *GitHubSourceActionProps) GitHubSourceAction {
	_init_.Initialize()

	j := jsiiProxy_GitHubSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewGitHubSourceAction_Override(g GitHubSourceAction, props *GitHubSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitHubSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubSourceAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

// Construction properties of the {@link GitHubSourceAction GitHub source action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Read the secret from Secrets Manager
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("GitHub_Source"),
//   	owner: jsii.String("awslabs"),
//   	repo: jsii.String("aws-cdk"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	output: sourceOutput,
//   	branch: jsii.String("develop"),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type GitHubSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	//    const oauth = cdk.SecretValue.secretsManager('my-github-token');
	//    new GitHubSource(this, 'GitHubAction', { oauthToken: oauth, ... });
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/appendix-github-oauth.html#GitHub-create-personal-token-CLI
	//
	OauthToken awscdk.SecretValue `field:"required" json:"oauthToken" yaml:"oauthToken"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The GitHub account/user that owns the repo.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repo, without the username.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to use.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action
	// With "POLL", CodePipeline periodically checks the source for changes
	// With "None", the action is not triggered through changes in the source
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	Trigger GitHubTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

// The CodePipeline variables emitted by GitHub source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubSourceVariables := &gitHubSourceVariables{
//   	authorDate: jsii.String("authorDate"),
//   	branchName: jsii.String("branchName"),
//   	commitId: jsii.String("commitId"),
//   	commitMessage: jsii.String("commitMessage"),
//   	committerDate: jsii.String("committerDate"),
//   	commitUrl: jsii.String("commitUrl"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type GitHubSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	AuthorDate *string `field:"required" json:"authorDate" yaml:"authorDate"`
	// The name of the branch this action tracks.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The message of the currently last commit on the tracked branch.
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// The date the currently last commit on the tracked branch was committed, in ISO-8601 format.
	CommitterDate *string `field:"required" json:"committerDate" yaml:"committerDate"`
	// The GitHub API URL of the currently last commit on the tracked branch.
	CommitUrl *string `field:"required" json:"commitUrl" yaml:"commitUrl"`
	// The name of the repository this action points to.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

// If and how the GitHub source action should be triggered.
type GitHubTrigger string

const (
	GitHubTrigger_NONE GitHubTrigger = "NONE"
	GitHubTrigger_POLL GitHubTrigger = "POLL"
	GitHubTrigger_WEBHOOK GitHubTrigger = "WEBHOOK"
)

// A Jenkins provider.
//
// If you want to create a new Jenkins provider managed alongside your CDK code,
// instantiate the {@link JenkinsProvider} class directly.
//
// If you want to reference an already registered provider,
// use the {@link JenkinsProvider#fromJenkinsProviderAttributes} method.
type IJenkinsProvider interface {
	constructs.IConstruct
	ProviderName() *string
	ServerUrl() *string
	Version() *string
}

// The jsii proxy for IJenkinsProvider
type jsiiProxy_IJenkinsProvider struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Jenkins build CodePipeline Action.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&jenkinsActionProps{
//   	actionName: jsii.String("JenkinsBuild"),
//   	jenkinsProvider: jenkinsProvider,
//   	projectName: jsii.String("MyProject"),
//   	type: codepipeline_actions.jenkinsActionType_BUILD,
//   })
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-four-stage-pipeline.html
//
type JenkinsAction interface {
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

// The jsii proxy struct for JenkinsAction
type jsiiProxy_JenkinsAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_JenkinsAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewJenkinsAction(props *JenkinsActionProps) JenkinsAction {
	_init_.Initialize()

	j := jsiiProxy_JenkinsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewJenkinsAction_Override(j JenkinsAction, props *JenkinsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsAction",
		[]interface{}{props},
		j,
	)
}

func (j *jsiiProxy_JenkinsAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		j,
		"bound",
		[]interface{}{_scope, _stage, _options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of {@link JenkinsAction}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&jenkinsActionProps{
//   	actionName: jsii.String("JenkinsBuild"),
//   	jenkinsProvider: jenkinsProvider,
//   	projectName: jsii.String("MyProject"),
//   	type: codepipeline_actions.jenkinsActionType_BUILD,
//   })
//
type JenkinsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Jenkins Provider for this Action.
	JenkinsProvider IJenkinsProvider `field:"required" json:"jenkinsProvider" yaml:"jenkinsProvider"`
	// The name of the project (sometimes also called job, or task) on your Jenkins installation that will be invoked by this Action.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "MyJob"
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The type of the Action - Build, or Test.
	Type JenkinsActionType `field:"required" json:"type" yaml:"type"`
	// The source to use as input for this build.
	Inputs *[]awscodepipeline.Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
}

// The type of the Jenkins Action that determines its CodePipeline Category - Build, or Test.
//
// Note that a Jenkins provider, even if it has the same name,
// must be separately registered for each type.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&jenkinsActionProps{
//   	actionName: jsii.String("JenkinsBuild"),
//   	jenkinsProvider: jenkinsProvider,
//   	projectName: jsii.String("MyProject"),
//   	type: codepipeline_actions.jenkinsActionType_BUILD,
//   })
//
type JenkinsActionType string

const (
	// The Action will have the Build Category.
	JenkinsActionType_BUILD JenkinsActionType = "BUILD"
	// The Action will have the Test Category.
	JenkinsActionType_TEST JenkinsActionType = "TEST"
)

// A class representing Jenkins providers.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
// See: #import.
//
type JenkinsProvider interface {
	BaseJenkinsProvider
	// The tree node.
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for JenkinsProvider
type jsiiProxy_JenkinsProvider struct {
	jsiiProxy_BaseJenkinsProvider
}

func (j *jsiiProxy_JenkinsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewJenkinsProvider(scope constructs.Construct, id *string, props *JenkinsProviderProps) JenkinsProvider {
	_init_.Initialize()

	j := jsiiProxy_JenkinsProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewJenkinsProvider_Override(j JenkinsProvider, scope constructs.Construct, id *string, props *JenkinsProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		[]interface{}{scope, id, props},
		j,
	)
}

// Import a Jenkins provider registered either outside the CDK, or in a different CDK Stack.
//
// Returns: a new Construct representing a reference to an existing Jenkins provider.
func JenkinsProvider_FromJenkinsProviderAttributes(scope constructs.Construct, id *string, attrs *JenkinsProviderAttributes) IJenkinsProvider {
	_init_.Initialize()

	var returns IJenkinsProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		"fromJenkinsProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func JenkinsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for importing an existing Jenkins provider.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   jenkinsProvider := codepipeline_actions.jenkinsProvider.fromJenkinsProviderAttributes(this, jsii.String("JenkinsProvider"), &jenkinsProviderAttributes{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
type JenkinsProviderAttributes struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "MyJenkinsProvider"
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "http://myjenkins.com:8080"
	//
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// The version of your provider.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
type JenkinsProviderProps struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "MyJenkinsProvider"
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "http://myjenkins.com:8080"
	//
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Whether to immediately register a Jenkins Provider for the build category.
	//
	// The Provider will always be registered if you create a {@link JenkinsAction}.
	ForBuild *bool `field:"optional" json:"forBuild" yaml:"forBuild"`
	// Whether to immediately register a Jenkins Provider for the test category.
	//
	// The Provider will always be registered if you create a {@link JenkinsTestAction}.
	ForTest *bool `field:"optional" json:"forTest" yaml:"forTest"`
	// The version of your provider.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// CodePipeline invoke Action that is provided by an AWS Lambda function.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var fn function
//
//   sourceOutput := codepipeline.NewArtifact()
//   buildOutput := codepipeline.NewArtifact()
//   lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
//   	actionName: jsii.String("Lambda"),
//   	inputs: []artifact{
//   		sourceOutput,
//   		buildOutput,
//   	},
//   	outputs: []*artifact{
//   		codepipeline.NewArtifact(jsii.String("Out1")),
//   		codepipeline.NewArtifact(jsii.String("Out2")),
//   	},
//   	lambda: fn,
//   })
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html
//
type LambdaInvokeAction interface {
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
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Reference a CodePipeline variable defined by the Lambda function this action points to.
	//
	// Variables in Lambda invoke actions are defined by calling the PutJobSuccessResult CodePipeline API call
	// with the 'outputVariables' property filled.
	// See: https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_PutJobSuccessResult.html
	//
	Variable(variableName *string) *string
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_LambdaInvokeAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewLambdaInvokeAction(props *LambdaInvokeActionProps) LambdaInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.LambdaInvokeAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLambdaInvokeAction_Override(l LambdaInvokeAction, props *LambdaInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.LambdaInvokeAction",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LambdaInvokeAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		l,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		l,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) Variable(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"variable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link LambdaInvokeAction Lambda invoke CodePipeline Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var fn function
//
//   sourceOutput := codepipeline.NewArtifact()
//   buildOutput := codepipeline.NewArtifact()
//   lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
//   	actionName: jsii.String("Lambda"),
//   	inputs: []artifact{
//   		sourceOutput,
//   		buildOutput,
//   	},
//   	outputs: []*artifact{
//   		codepipeline.NewArtifact(jsii.String("Out1")),
//   		codepipeline.NewArtifact(jsii.String("Out2")),
//   	},
//   	lambda: fn,
//   })
//
type LambdaInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The lambda function to invoke.
	Lambda awslambda.IFunction `field:"required" json:"lambda" yaml:"lambda"`
	// The optional input Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 inputs.
	// The inputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.inputArtifacts` path.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	Inputs *[]awscodepipeline.Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	// The optional names of the output Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 outputs.
	// The outputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.outputArtifacts` path.
	// It is the responsibility of the Lambda to upload ZIP files with the Artifact contents to the provided locations.
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	// A set of key-value pairs that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParameters` or `userParametersString` can be specified.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	UserParameters *map[string]interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
	// The string representation of the user parameters that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParametersString` or `userParameters` can be specified.
	UserParametersString *string `field:"optional" json:"userParametersString" yaml:"userParametersString"`
}

// Manual approval action.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
//   	actionName: jsii.String("Approve"),
//   })
//   approveStage.addAction(manualApprovalAction)
//
//   role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
//   	service: jsii.String("iam"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.grantManualApproval(role)
//
type ManualApprovalAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	NotificationTopic() awssns.ITopic
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// grant the provided principal the permissions to approve or reject this manual approval action.
	//
	// For more info see:
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/approvals-iam-permissions.html
	GrantManualApproval(grantable awsiam.IGrantable)
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for ManualApprovalAction
type jsiiProxy_ManualApprovalAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_ManualApprovalAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalAction) NotificationTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"notificationTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewManualApprovalAction(props *ManualApprovalActionProps) ManualApprovalAction {
	_init_.Initialize()

	j := jsiiProxy_ManualApprovalAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ManualApprovalAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewManualApprovalAction_Override(m ManualApprovalAction, props *ManualApprovalActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ManualApprovalAction",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_ManualApprovalAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		m,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) GrantManualApproval(grantable awsiam.IGrantable) {
	_jsii_.InvokeVoid(
		m,
		"grantManualApproval",
		[]interface{}{grantable},
	)
}

func (m *jsiiProxy_ManualApprovalAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		m,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link ManualApprovalAction}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
//   	actionName: jsii.String("Approve"),
//   })
//   approveStage.addAction(manualApprovalAction)
//
//   role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
//   	service: jsii.String("iam"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.grantManualApproval(role)
//
type ManualApprovalActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Any additional information that you want to include in the notification email message.
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// URL you want to provide to the reviewer as part of the approval request.
	ExternalEntityLink *string `field:"optional" json:"externalEntityLink" yaml:"externalEntityLink"`
	// Optional SNS topic to send notifications to when an approval is pending.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// A list of email addresses to subscribe to notifications when this Action is pending approval.
	//
	// If this has been provided, but not `notificationTopic`,
	// a new Topic will be created.
	NotifyEmails *[]*string `field:"optional" json:"notifyEmails" yaml:"notifyEmails"`
}

// Properties for configuring service-managed (Organizations) permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationsDeploymentProps := &organizationsDeploymentProps{
//   	autoDeployment: awscdk.Aws_codepipeline_actions.stackSetOrganizationsAutoDeployment_ENABLED,
//   }
//
type OrganizationsDeploymentProps struct {
	// Automatically deploy to new accounts added to Organizational Units.
	//
	// Whether AWS CloudFormation StackSets automatically deploys to AWS
	// Organizations accounts that are added to a target organization or
	// organizational unit (OU).
	AutoDeployment StackSetOrganizationsAutoDeployment `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
}

// Deploys the sourceArtifact to Amazon S3.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   	actionName: jsii.String("S3Deploy"),
//   	bucket: targetBucket,
//   	input: sourceOutput,
//   })
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type S3DeployAction interface {
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

// The jsii proxy struct for S3DeployAction
type jsiiProxy_S3DeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_S3DeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewS3DeployAction(props *S3DeployActionProps) S3DeployAction {
	_init_.Initialize()

	j := jsiiProxy_S3DeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3DeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3DeployAction_Override(s S3DeployAction, props *S3DeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3DeployAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3DeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeployAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link S3DeployAction S3 deploy Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   	actionName: jsii.String("S3Deploy"),
//   	bucket: targetBucket,
//   	input: sourceOutput,
//   })
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type S3DeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 bucket that is the deploy target.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The input Artifact to deploy to Amazon S3.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The specified canned ACL to objects deployed to Amazon S3.
	//
	// This overwrites any existing ACL that was applied to the object.
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// The caching behavior for requests/responses for objects in the bucket.
	//
	// The final cache control property will be the result of joining all of the provided array elements with a comma
	// (plus a space after the comma).
	CacheControl *[]CacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Should the deploy action extract the artifact before deploying to Amazon S3.
	Extract *bool `field:"optional" json:"extract" yaml:"extract"`
	// The key of the target object.
	//
	// This is required if extract is false.
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

// Source that is provided by a specific Amazon S3 object.
//
// Will trigger the pipeline as soon as the S3 object changes, but only if there is
// a CloudTrail Trail in the account that captures the S3 event.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.addS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		bucket: sourceBucket,
//   		objectPrefix: key,
//   	},
//   }, &addEventSelectorOptions{
//   	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucketKey: key,
//   	bucket: sourceBucket,
//   	output: sourceOutput,
//   	trigger: codepipeline_actions.s3Trigger_EVENTS,
//   })
//
type S3SourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *S3SourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
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
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link S3SourceAction S3 source Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.addS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		bucket: sourceBucket,
//   		objectPrefix: key,
//   	},
//   }, &addEventSelectorOptions{
//   	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucketKey: key,
//   	bucket: sourceBucket,
//   	output: sourceOutput,
//   	trigger: codepipeline_actions.s3Trigger_EVENTS,
//   })
//
type S3SourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 bucket that stores the source code.
	//
	// If you import an encrypted bucket in your stack, please specify
	// the encryption key at import time by using `Bucket.fromBucketAttributes()` method.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key within the S3 bucket that stores the source code.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "path/to/file.zip"
	//
	BucketKey *string `field:"required" json:"bucketKey" yaml:"bucketKey"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	Trigger S3Trigger `field:"optional" json:"trigger" yaml:"trigger"`
}

// The CodePipeline variables emitted by the S3 source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceVariables := &s3SourceVariables{
//   	eTag: jsii.String("eTag"),
//   	versionId: jsii.String("versionId"),
//   }
//
type S3SourceVariables struct {
	// The e-tag of the S3 version of the object that triggered the build.
	ETag *string `field:"required" json:"eTag" yaml:"eTag"`
	// The identifier of the S3 version of the object that triggered the build.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
}

// How should the S3 Action detect changes.
//
// This is the type of the {@link S3SourceAction.trigger} property.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.addS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		bucket: sourceBucket,
//   		objectPrefix: key,
//   	},
//   }, &addEventSelectorOptions{
//   	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucketKey: key,
//   	bucket: sourceBucket,
//   	output: sourceOutput,
//   	trigger: codepipeline_actions.s3Trigger_EVENTS,
//   })
//
type S3Trigger string

const (
	// The Action will never detect changes - the Pipeline it's part of will only begin a run when explicitly started.
	S3Trigger_NONE S3Trigger = "NONE"
	// CodePipeline will poll S3 to detect changes.
	//
	// This is the default method of detecting changes.
	S3Trigger_POLL S3Trigger = "POLL"
	// CodePipeline will use CloudWatch Events to be notified of changes.
	//
	// Note that the Bucket that the Action uses needs to be part of a CloudTrail Trail
	// for the events to be delivered.
	S3Trigger_EVENTS S3Trigger = "EVENTS"
)

// Properties for configuring self-managed permissions.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   existingAdminRole := iam.role.fromRoleName(this, jsii.String("AdminRole"), jsii.String("AWSCloudFormationStackSetAdministrationRole"))
//
//   deploymentModel := codepipeline_actions.stackSetDeploymentModel.selfManaged(&selfManagedDeploymentProps{
//   	// Use an existing Role. Leave this out to create a new Role.
//   	administrationRole: existingAdminRole,
//   })
//
type SelfManagedDeploymentProps struct {
	// The IAM role in the administrator account used to assume execution roles in the target accounts.
	//
	// You must create this role before using the StackSet action.
	//
	// The role needs to be assumable by CloudFormation, and it needs to be able
	// to `sts:AssumeRole` each of the execution roles (whose names are specified
	// in the `executionRoleName` parameter) in each of the target accounts.
	//
	// If you do not specify the role, we assume you have created a role named
	// `AWSCloudFormationStackSetAdministrationRole`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
	//
	AdministrationRole awsiam.IRole `field:"optional" json:"administrationRole" yaml:"administrationRole"`
	// The name of the IAM role in the target accounts used to perform stack set operations.
	//
	// You must create these roles in each of the target accounts before using the
	// StackSet action.
	//
	// The roles need to be assumable by by the `administrationRole`, and need to
	// have the permissions necessary to successfully create and modify the
	// resources that the subsequent CloudFormation deployments need.
	// Administrator permissions would be commonly granted to these, but if you can
	// scope the permissions down frome there you would be safer.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
	//
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
}

// CodePipeline action to connect to an existing ServiceCatalog product.
//
// **Note**: this class is still experimental, and may have breaking changes in the future!
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   cdkBuildOutput := codepipeline.NewArtifact()
//   serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&serviceCatalogDeployActionBeta1Props{
//   	actionName: jsii.String("ServiceCatalogDeploy"),
//   	templatePath: cdkBuildOutput.atPath(jsii.String("Sample.template.json")),
//   	productVersionName: jsii.String("Version - " + date.now.toString),
//   	productVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
//   	productId: jsii.String("prod-XXXXXXXX"),
//   })
//
type ServiceCatalogDeployActionBeta1 interface {
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

// The jsii proxy struct for ServiceCatalogDeployActionBeta1
type jsiiProxy_ServiceCatalogDeployActionBeta1 struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_ServiceCatalogDeployActionBeta1) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCatalogDeployActionBeta1) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewServiceCatalogDeployActionBeta1(props *ServiceCatalogDeployActionBeta1Props) ServiceCatalogDeployActionBeta1 {
	_init_.Initialize()

	j := jsiiProxy_ServiceCatalogDeployActionBeta1{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewServiceCatalogDeployActionBeta1_Override(s ServiceCatalogDeployActionBeta1, props *ServiceCatalogDeployActionBeta1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link ServiceCatalogDeployActionBeta1 ServiceCatalog deploy CodePipeline Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   cdkBuildOutput := codepipeline.NewArtifact()
//   serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&serviceCatalogDeployActionBeta1Props{
//   	actionName: jsii.String("ServiceCatalogDeploy"),
//   	templatePath: cdkBuildOutput.atPath(jsii.String("Sample.template.json")),
//   	productVersionName: jsii.String("Version - " + date.now.toString),
//   	productVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
//   	productId: jsii.String("prod-XXXXXXXX"),
//   })
//
type ServiceCatalogDeployActionBeta1Props struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The identifier of the product in the Service Catalog.
	//
	// This product must already exist.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The name of the version of the Service Catalog product to be deployed.
	ProductVersionName *string `field:"required" json:"productVersionName" yaml:"productVersionName"`
	// The path to the cloudformation artifact.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The optional description of this version of the Service Catalog product.
	ProductVersionDescription *string `field:"optional" json:"productVersionDescription" yaml:"productVersionDescription"`
}

// Where Stack Instances will be created from the StackSet.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type StackInstances interface {
}

// The jsii proxy struct for StackInstances
type jsiiProxy_StackInstances struct {
	_ byte // padding
}

func NewStackInstances_Override(s StackInstances) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		nil, // no parameters
		s,
	)
}

// Create stack instances in a set of accounts or organizational units taken from the pipeline artifacts, and a set of regions  The file must be a JSON file containing a list of strings.
//
// For example:
//
// ```json
// [
//    "111111111111",
//    "222222222222",
//    "333333333333"
// ]
// ```
//
// Stack Instances will be created in every combination of region and account, or region and
// Organizational Units (OUs).
//
// If this is set of Organizational Units, you must have selected `StackSetDeploymentModel.organizations()`
// as deployment model.
func StackInstances_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath, regions *[]*string) StackInstances {
	_init_.Initialize()

	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"fromArtifactPath",
		[]interface{}{artifactPath, regions},
		&returns,
	)

	return returns
}

// Create stack instances in a set of accounts and regions passed as literal lists.
//
// Stack Instances will be created in every combination of region and account.
//
// > NOTE: `StackInstances.inAccounts()` and `StackInstances.inOrganizationalUnits()`
// > have exactly the same behavior, and you can use them interchangeably if you want.
// > The only difference between them is that your code clearly indicates what entity
// > it's working with.
func StackInstances_InAccounts(accounts *[]*string, regions *[]*string) StackInstances {
	_init_.Initialize()

	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"inAccounts",
		[]interface{}{accounts, regions},
		&returns,
	)

	return returns
}

// Create stack instances in all accounts in a set of Organizational Units (OUs) and regions passed as literal lists.
//
// If you want to deploy to Organization Units, you must choose have created the StackSet
// with `deploymentModel: DeploymentModel.organizations()`.
//
// Stack Instances will be created in every combination of region and account.
//
// > NOTE: `StackInstances.inAccounts()` and `StackInstances.inOrganizationalUnits()`
// > have exactly the same behavior, and you can use them interchangeably if you want.
// > The only difference between them is that your code clearly indicates what entity
// > it's working with.
func StackInstances_InOrganizationalUnits(ous *[]*string, regions *[]*string) StackInstances {
	_init_.Initialize()

	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"inOrganizationalUnits",
		[]interface{}{ous, regions},
		&returns,
	)

	return returns
}

// Determines how IAM roles are created and managed.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type StackSetDeploymentModel interface {
}

// The jsii proxy struct for StackSetDeploymentModel
type jsiiProxy_StackSetDeploymentModel struct {
	_ byte // padding
}

func NewStackSetDeploymentModel_Override(s StackSetDeploymentModel) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		nil, // no parameters
		s,
	)
}

// Deploy to AWS Organizations accounts.
//
// AWS CloudFormation StackSets automatically creates the IAM roles required
// to deploy to accounts managed by AWS Organizations. This requires an
// account to be a member of an Organization.
//
// Using this deployment model, you can specify either AWS Account Ids or
// Organization Unit Ids in the `stackInstances` parameter.
func StackSetDeploymentModel_Organizations(props *OrganizationsDeploymentProps) StackSetDeploymentModel {
	_init_.Initialize()

	var returns StackSetDeploymentModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		"organizations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Deploy to AWS Accounts not managed by AWS Organizations.
//
// You are responsible for creating Execution Roles in every account you will
// be deploying to in advance to create the actual stack instances. Unless you
// specify overrides, StackSets expects the execution roles you create to have
// the default name `AWSCloudFormationStackSetExecutionRole`. See the [Grant
// self-managed
// permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html)
// section of the CloudFormation documentation.
//
// The CDK will automatically create the central Administration Role in the
// Pipeline account which will be used to assume the Execution Role in each of
// the target accounts.
//
// If you wish to use a pre-created Administration Role, use `Role.fromRoleName()`
// or `Role.fromRoleArn()` to import it, and pass it to this function:
//
// ```ts
// const existingAdminRole = iam.Role.fromRoleName(this, 'AdminRole', 'AWSCloudFormationStackSetAdministrationRole');
//
// const deploymentModel = codepipeline_actions.StackSetDeploymentModel.selfManaged({
//    // Use an existing Role. Leave this out to create a new Role.
//    administrationRole: existingAdminRole,
// });
// ```
//
// Using this deployment model, you can only specify AWS Account Ids in the
// `stackInstances` parameter.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
//
func StackSetDeploymentModel_SelfManaged(props *SelfManagedDeploymentProps) StackSetDeploymentModel {
	_init_.Initialize()

	var returns StackSetDeploymentModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		"selfManaged",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Describes whether AWS CloudFormation StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
type StackSetOrganizationsAutoDeployment string

const (
	// StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, AWS CloudFormation StackSets
	// deletes stack instances from the account in the specified Regions.
	StackSetOrganizationsAutoDeployment_ENABLED StackSetOrganizationsAutoDeployment = "ENABLED"
	// StackSets does not automatically deploy additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	StackSetOrganizationsAutoDeployment_DISABLED StackSetOrganizationsAutoDeployment = "DISABLED"
	// Stack resources are retained when an account is removed from a target organization or OU.
	StackSetOrganizationsAutoDeployment_ENABLED_WITH_STACK_RETENTION StackSetOrganizationsAutoDeployment = "ENABLED_WITH_STACK_RETENTION"
)

// Base parameters for the StackSet.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   parameters := codepipeline_actions.stackSetParameters.fromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
type StackSetParameters interface {
}

// The jsii proxy struct for StackSetParameters
type jsiiProxy_StackSetParameters struct {
	_ byte // padding
}

func NewStackSetParameters_Override(s StackSetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		nil, // no parameters
		s,
	)
}

// Read the parameters from a JSON file from one of the pipeline's artifacts.
//
// The file needs to contain a list of `{ ParameterKey, ParameterValue, UsePreviousValue }` objects, like
// this:
//
// ```
// [
//      {
//          "ParameterKey": "BucketName",
//          "ParameterValue": "my-bucket"
//      },
//      {
//          "ParameterKey": "Asset1",
//          "ParameterValue": "true"
//      },
//      {
//          "ParameterKey": "Asset2",
//          "UsePreviousValue": true
//      }
// ]
// ```
//
// You must specify all template parameters. Parameters you don't specify will revert
// to their `Default` values as specified in the template.
//
// For of parameters you want to retain their existing values
// without specifying what those values are, set `UsePreviousValue: true`.
// Use of this feature is discouraged. CDK is for
// specifying desired-state infrastructure, and use of this feature makes the
// parameter values unmanaged.
func StackSetParameters_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath) StackSetParameters {
	_init_.Initialize()

	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		"fromArtifactPath",
		[]interface{}{artifactPath},
		&returns,
	)

	return returns
}

// A list of template parameters for your stack set.
//
// You must specify all template parameters. Parameters you don't specify will revert
// to their `Default` values as specified in the template.
//
// Specify the names of parameters you want to retain their existing values,
// without specifying what those values are, in an array in the second
// argument to this function. Use of this feature is discouraged. CDK is for
// specifying desired-state infrastructure, and use of this feature makes the
// parameter values unmanaged.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   parameters := codepipeline_actions.stackSetParameters.fromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
func StackSetParameters_FromLiteral(parameters *map[string]*string, usePreviousValues *[]*string) StackSetParameters {
	_init_.Initialize()

	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		"fromLiteral",
		[]interface{}{parameters, usePreviousValues},
		&returns,
	)

	return returns
}

// The source of a StackSet template.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type StackSetTemplate interface {
}

// The jsii proxy struct for StackSetTemplate
type jsiiProxy_StackSetTemplate struct {
	_ byte // padding
}

func NewStackSetTemplate_Override(s StackSetTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetTemplate",
		nil, // no parameters
		s,
	)
}

// Use a file in an artifact as Stack Template.
func StackSetTemplate_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath) StackSetTemplate {
	_init_.Initialize()

	var returns StackSetTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetTemplate",
		"fromArtifactPath",
		[]interface{}{artifactPath},
		&returns,
	)

	return returns
}

// Represents the input for the StateMachine.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StateMachineInput interface {
	// When InputType is set to Literal (default), the Input field is used directly as the input for the state machine execution.
	//
	// Otherwise, the state machine is invoked with an empty JSON object {}.
	//
	// When InputType is set to FilePath, this field is required.
	// An input artifact is also required when InputType is set to FilePath.
	Input() interface{}
	// The optional input Artifact of the Action.
	//
	// If InputType is set to FilePath, this artifact is required
	// and is used to source the input for the state machine execution.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-StepFunctions.html#action-reference-StepFunctions-example
	//
	InputArtifact() awscodepipeline.Artifact
	// Optional StateMachine InputType InputType can be Literal or FilePath.
	InputType() *string
}

// The jsii proxy struct for StateMachineInput
type jsiiProxy_StateMachineInput struct {
	_ byte // padding
}

func (j *jsiiProxy_StateMachineInput) Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineInput) InputArtifact() awscodepipeline.Artifact {
	var returns awscodepipeline.Artifact
	_jsii_.Get(
		j,
		"inputArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineInput) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}


// When the input type is FilePath, input artifact and filepath must be specified.
func StateMachineInput_FilePath(inputFile awscodepipeline.ArtifactPath) StateMachineInput {
	_init_.Initialize()

	var returns StateMachineInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StateMachineInput",
		"filePath",
		[]interface{}{inputFile},
		&returns,
	)

	return returns
}

// When the input type is Literal, input value is passed directly to the state machine input.
func StateMachineInput_Literal(object *map[string]interface{}) StateMachineInput {
	_init_.Initialize()

	var returns StateMachineInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StateMachineInput",
		"literal",
		[]interface{}{object},
		&returns,
	)

	return returns
}

// StepFunctionInvokeAction that is provided by an AWS CodePipeline.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StepFunctionInvokeAction interface {
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

// The jsii proxy struct for StepFunctionInvokeAction
type jsiiProxy_StepFunctionInvokeAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_StepFunctionInvokeAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionInvokeAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewStepFunctionInvokeAction(props *StepFunctionsInvokeActionProps) StepFunctionInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionInvokeAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StepFunctionInvokeAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStepFunctionInvokeAction_Override(s StepFunctionInvokeAction, props *StepFunctionsInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StepFunctionInvokeAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionInvokeAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionInvokeAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Construction properties of the {@link StepFunctionsInvokeAction StepFunction Invoke Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StepFunctionsInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The state machine to invoke.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Prefix (optional).
	//
	// By default, the action execution ID is used as the state machine execution name.
	// If a prefix is provided, it is prepended to the action execution ID with a hyphen and
	// together used as the state machine execution name.
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
	// The optional output Artifact of the Action.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// Represents the input to the StateMachine.
	//
	// This includes input artifact, input type and the statemachine input.
	StateMachineInput StateMachineInput `field:"optional" json:"stateMachineInput" yaml:"stateMachineInput"`
}

