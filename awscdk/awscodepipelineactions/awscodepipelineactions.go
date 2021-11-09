package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation"
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
// Experimental.
type Action interface {
	awscodepipeline.Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
func NewAction_Override(a Action, actionProperties *awscodepipeline.ActionProperties) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Action",
		[]interface{}{actionProperties},
		a,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type AlexaSkillDeployAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewAlexaSkillDeployAction_Override(a AlexaSkillDeployAction, props *AlexaSkillDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.AlexaSkillDeployAction",
		[]interface{}{props},
		a,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type AlexaSkillDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The client id of the developer console token.
	// Experimental.
	ClientId *string `json:"clientId"`
	// The client secret of the developer console token.
	// Experimental.
	ClientSecret awscdk.SecretValue `json:"clientSecret"`
	// The source artifact containing the voice model and skill manifest.
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
	// The refresh token of the developer console token.
	// Experimental.
	RefreshToken awscdk.SecretValue `json:"refreshToken"`
	// The Alexa skill id.
	// Experimental.
	SkillId *string `json:"skillId"`
	// An optional artifact containing overrides for the skill manifest.
	// Experimental.
	ParameterOverridesArtifact awscodepipeline.Artifact `json:"parameterOverridesArtifact"`
}

// Experimental.
type BaseJenkinsProvider interface {
	constructs.Construct
	IJenkinsProvider
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
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


// Experimental.
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Returns a string representation of this construct.
// Experimental.
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

// A CodePipeline source action for BitBucket.
// Deprecated: use CodeStarConnectionsSourceAction instead
type BitBucketSourceAction interface {
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for BitBucketSourceAction
type jsiiProxy_BitBucketSourceAction struct {
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_BitBucketSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}


// Deprecated: use CodeStarConnectionsSourceAction instead
func NewBitBucketSourceAction(props *BitBucketSourceActionProps) BitBucketSourceAction {
	_init_.Initialize()

	j := jsiiProxy_BitBucketSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.BitBucketSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: use CodeStarConnectionsSourceAction instead
func NewBitBucketSourceAction_Override(b BitBucketSourceAction, props *BitBucketSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.BitBucketSourceAction",
		[]interface{}{props},
		b,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Deprecated: use CodeStarConnectionsSourceAction instead
func (b *jsiiProxy_BitBucketSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// Creates an Event that will be triggered whenever the state of this Action changes.
// Deprecated: use CodeStarConnectionsSourceAction instead
func (b *jsiiProxy_BitBucketSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Construction properties for {@link BitBucketSourceAction}.
// Deprecated: use CodeStarConnectionsSourceActionProps instead
type BitBucketSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	Role awsiam.IRole `json:"role"`
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// TODO: EXAMPLE
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	ConnectionArn *string `json:"connectionArn"`
	// The output artifact that this action produces.
	//
	// Can be used as input for further pipeline actions.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	Output awscodepipeline.Artifact `json:"output"`
	// The owning user or organization of the repository.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	Owner *string `json:"owner"`
	// The name of the repository.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	Repo *string `json:"repo"`
	// The branch to build.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	Branch *string `json:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead
	TriggerOnPush *bool `json:"triggerOnPush"`
}

// Used for HTTP cache-control header, which influences downstream caches.
//
// Use the provided static factory methods to construct instances of this class.
// Used in the {@link S3DeployActionProps.cacheControl} property.
// See: https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
//
// Experimental.
type CacheControl interface {
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type CloudFormationCreateReplaceChangeSetAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCloudFormationCreateReplaceChangeSetAction_Override(c CloudFormationCreateReplaceChangeSetAction, props *CloudFormationCreateReplaceChangeSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		[]interface{}{props},
		c,
	)
}

// Add statement to the service role assumed by CloudFormation while executing this action.
// Experimental.
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

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CloudFormationCreateReplaceChangeSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
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
	// Experimental.
	AdminPermissions *bool `json:"adminPermissions"`
	// Name of the change set to create or update.
	// Experimental.
	ChangeSetName *string `json:"changeSetName"`
	// The name of the stack to apply this action to.
	// Experimental.
	StackName *string `json:"stackName"`
	// Input artifact with the ChangeSet's CloudFormation template.
	// Experimental.
	TemplatePath awscodepipeline.ArtifactPath `json:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	// Experimental.
	Account *string `json:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources, explicit acknowledgement that AWS CloudFormation
	// might create or update those resources. For example, you must specify `AnonymousIAM` or `NamedIAM`
	// if your stack template contains AWS Identity and Access Management (IAM) resources. For more
	// information see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Deprecated: use {@link cfnCapabilities} instead
	Capabilities *[]awscloudformation.CloudFormationCapabilities `json:"capabilities"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Experimental.
	CfnCapabilities *[]awscdk.CfnCapabilities `json:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	// Experimental.
	DeploymentRole awsiam.IRole `json:"deploymentRole"`
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
	// Experimental.
	ExtraInputs *[]awscodepipeline.Artifact `json:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
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
	// Experimental.
	ParameterOverrides *map[string]interface{} `json:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	Region *string `json:"region"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	// Experimental.
	TemplateConfiguration awscodepipeline.ArtifactPath `json:"templateConfiguration"`
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
// Experimental.
type CloudFormationCreateUpdateStackAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCloudFormationCreateUpdateStackAction_Override(c CloudFormationCreateUpdateStackAction, props *CloudFormationCreateUpdateStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		c,
	)
}

// Add statement to the service role assumed by CloudFormation while executing this action.
// Experimental.
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

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CloudFormationCreateUpdateStackActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
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
	// Experimental.
	AdminPermissions *bool `json:"adminPermissions"`
	// The name of the stack to apply this action to.
	// Experimental.
	StackName *string `json:"stackName"`
	// Input artifact with the CloudFormation template to deploy.
	// Experimental.
	TemplatePath awscodepipeline.ArtifactPath `json:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	// Experimental.
	Account *string `json:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources, explicit acknowledgement that AWS CloudFormation
	// might create or update those resources. For example, you must specify `AnonymousIAM` or `NamedIAM`
	// if your stack template contains AWS Identity and Access Management (IAM) resources. For more
	// information see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Deprecated: use {@link cfnCapabilities} instead
	Capabilities *[]awscloudformation.CloudFormationCapabilities `json:"capabilities"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Experimental.
	CfnCapabilities *[]awscdk.CfnCapabilities `json:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	// Experimental.
	DeploymentRole awsiam.IRole `json:"deploymentRole"`
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
	// Experimental.
	ExtraInputs *[]awscodepipeline.Artifact `json:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
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
	// Experimental.
	ParameterOverrides *map[string]interface{} `json:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	Region *string `json:"region"`
	// Replace the stack if it's in a failed state.
	//
	// If this is set to true and the stack is in a failed state (one of
	// ROLLBACK_COMPLETE, ROLLBACK_FAILED, CREATE_FAILED, DELETE_FAILED, or
	// UPDATE_ROLLBACK_FAILED), AWS CloudFormation deletes the stack and then
	// creates a new stack.
	//
	// If this is not set to true and the stack is in a failed state,
	// the deployment fails.
	// Experimental.
	ReplaceOnFailure *bool `json:"replaceOnFailure"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	// Experimental.
	TemplateConfiguration awscodepipeline.ArtifactPath `json:"templateConfiguration"`
}

// CodePipeline action to delete a stack.
//
// Deletes a stack. If you specify a stack that doesn't exist, the action completes successfully
// without deleting a stack.
// Experimental.
type CloudFormationDeleteStackAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCloudFormationDeleteStackAction_Override(c CloudFormationDeleteStackAction, props *CloudFormationDeleteStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		c,
	)
}

// Add statement to the service role assumed by CloudFormation while executing this action.
// Experimental.
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

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CloudFormationDeleteStackActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
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
	// Experimental.
	AdminPermissions *bool `json:"adminPermissions"`
	// The name of the stack to apply this action to.
	// Experimental.
	StackName *string `json:"stackName"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	// Experimental.
	Account *string `json:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources, explicit acknowledgement that AWS CloudFormation
	// might create or update those resources. For example, you must specify `AnonymousIAM` or `NamedIAM`
	// if your stack template contains AWS Identity and Access Management (IAM) resources. For more
	// information see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Deprecated: use {@link cfnCapabilities} instead
	Capabilities *[]awscloudformation.CloudFormationCapabilities `json:"capabilities"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Experimental.
	CfnCapabilities *[]awscdk.CfnCapabilities `json:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	// Experimental.
	DeploymentRole awsiam.IRole `json:"deploymentRole"`
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
	// Experimental.
	ExtraInputs *[]awscodepipeline.Artifact `json:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
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
	// Experimental.
	ParameterOverrides *map[string]interface{} `json:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	Region *string `json:"region"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	// Experimental.
	TemplateConfiguration awscodepipeline.ArtifactPath `json:"templateConfiguration"`
}

// CodePipeline action to execute a prepared change set.
// Experimental.
type CloudFormationExecuteChangeSetAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCloudFormationExecuteChangeSetAction_Override(c CloudFormationExecuteChangeSetAction, props *CloudFormationExecuteChangeSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationExecuteChangeSetAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CloudFormationExecuteChangeSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Name of the change set to execute.
	// Experimental.
	ChangeSetName *string `json:"changeSetName"`
	// The name of the stack to apply this action to.
	// Experimental.
	StackName *string `json:"stackName"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	// Experimental.
	Account *string `json:"account"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	Region *string `json:"region"`
}

// CodePipeline build action that uses AWS CodeBuild.
// Experimental.
type CodeBuildAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
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


// Experimental.
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

// Experimental.
func NewCodeBuildAction_Override(c CodeBuildAction, props *CodeBuildActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeBuildAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Reference a CodePipeline variable defined by the CodeBuild project this action points to.
//
// Variables in CodeBuild actions are defined using the 'exported-variables' subsection of the 'env'
// section of the buildspec.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-syntax
//
// Experimental.
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

// Experimental.
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
// Experimental.
type CodeBuildActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The source to use as input for this action.
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
	// The action's Project.
	// Experimental.
	Project awscodebuild.IProject `json:"project"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	// Experimental.
	CheckSecretsInPlainTextEnvVariables *bool `json:"checkSecretsInPlainTextEnvVariables"`
	// Combine the build artifacts for a batch builds.
	//
	// Enabling this will combine the build artifacts into the same location for batch builds.
	// If `executeBatchBuild` is not set to `true`, this property is ignored.
	// Experimental.
	CombineBatchBuildArtifacts *bool `json:"combineBatchBuildArtifacts"`
	// The environment variables to pass to the CodeBuild project when this action executes.
	//
	// If a variable with the same name was set both on the project level, and here,
	// this value will take precedence.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Trigger a batch build.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	// Experimental.
	ExecuteBatchBuild *bool `json:"executeBatchBuild"`
	// The list of additional input Artifacts for this action.
	//
	// The directories the additional inputs will be available at are available
	// during the project's build in the CODEBUILD_SRC_DIR_<artifact-name> environment variables.
	// The project's build always starts in the directory with the primary input artifact checked out,
	// the one pointed to by the {@link input} property.
	// For more information,
	// see https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html .
	// Experimental.
	ExtraInputs *[]awscodepipeline.Artifact `json:"extraInputs"`
	// The list of output Artifacts for this action.
	//
	// **Note**: if you specify more than one output Artifact here,
	// you cannot use the primary 'artifacts' section of the buildspec;
	// you have to use the 'secondary-artifacts' section instead.
	// See https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	// for details.
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `json:"outputs"`
	// The type of the action that determines its CodePipeline Category - Build, or Test.
	// Experimental.
	Type CodeBuildActionType `json:"type"`
}

// The type of the CodeBuild action that determines its CodePipeline Category - Build, or Test.
//
// The default is Build.
// Experimental.
type CodeBuildActionType string

const (
	CodeBuildActionType_BUILD CodeBuildActionType = "BUILD"
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
// Experimental.
type CodeCommitSourceAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Variables() *CodeCommitSourceVariables
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCodeCommitSourceAction_Override(c CodeCommitSourceAction, props *CodeCommitSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeCommitSourceAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CodeCommitSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// The CodeCommit repository.
	// Experimental.
	Repository awscodecommit.IRepository `json:"repository"`
	// Experimental.
	Branch *string `json:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	// Experimental.
	Trigger CodeCommitTrigger `json:"trigger"`
}

// The CodePipeline variables emitted by the CodeCommit source Action.
// Experimental.
type CodeCommitSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	// Experimental.
	AuthorDate *string `json:"authorDate"`
	// The name of the branch this action tracks.
	// Experimental.
	BranchName *string `json:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	// Experimental.
	CommitId *string `json:"commitId"`
	// The message of the currently last commit on the tracked branch.
	// Experimental.
	CommitMessage *string `json:"commitMessage"`
	// The date the currently last commit on the tracked branch was committed, in ISO-8601 format.
	// Experimental.
	CommitterDate *string `json:"committerDate"`
	// The name of the repository this action points to.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
}

// How should the CodeCommit Action detect changes.
//
// This is the type of the {@link CodeCommitSourceAction.trigger} property.
// Experimental.
type CodeCommitTrigger string

const (
	CodeCommitTrigger_NONE CodeCommitTrigger = "NONE"
	CodeCommitTrigger_POLL CodeCommitTrigger = "POLL"
	CodeCommitTrigger_EVENTS CodeCommitTrigger = "EVENTS"
)

// Configuration for replacing a placeholder string in the ECS task definition template file with an image URI.
// Experimental.
type CodeDeployEcsContainerImageInput struct {
	// The artifact that contains an `imageDetails.json` file with the image URI.
	//
	// The artifact's `imageDetails.json` file must be a JSON file containing an
	// `ImageURI` property.  For example:
	// `{ "ImageURI": "ACCOUNTID.dkr.ecr.us-west-2.amazonaws.com/dk-image-repo@sha256:example3" }`
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
	// The placeholder string in the ECS task definition template file that will be replaced with the image URI.
	//
	// The placeholder string must be surrounded by angle brackets in the template file.
	// For example, if the task definition template file contains a placeholder like
	// `"image": "<PLACEHOLDER>"`, then the `taskDefinitionPlaceholder` value should
	// be `PLACEHOLDER`.
	// Experimental.
	TaskDefinitionPlaceholder *string `json:"taskDefinitionPlaceholder"`
}

// Experimental.
type CodeDeployEcsDeployAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCodeDeployEcsDeployAction_Override(c CodeDeployEcsDeployAction, props *CodeDeployEcsDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CodeDeployEcsDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The CodeDeploy ECS Deployment Group to deploy to.
	// Experimental.
	DeploymentGroup awscodedeploy.IEcsDeploymentGroup `json:"deploymentGroup"`
	// The name of the CodeDeploy AppSpec file.
	//
	// During deployment, a new task definition will be registered
	// with ECS, and the new task definition ID will be inserted into
	// the CodeDeploy AppSpec file.  The AppSpec file contents will be
	// provided to CodeDeploy for the deployment.
	//
	// Use this property if you want to use a different name for this file than the default 'appspec.yaml'.
	// If you use this property, you don't need to specify the `appSpecTemplateInput` property.
	// Experimental.
	AppSpecTemplateFile awscodepipeline.ArtifactPath `json:"appSpecTemplateFile"`
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
	// Experimental.
	AppSpecTemplateInput awscodepipeline.Artifact `json:"appSpecTemplateInput"`
	// Configuration for dynamically updated images in the task definition.
	//
	// Provide pairs of an image details input artifact and a placeholder string
	// that will be used to dynamically update the ECS task definition template
	// file prior to deployment. A maximum of 4 images can be given.
	// Experimental.
	ContainerImageInputs *[]*CodeDeployEcsContainerImageInput `json:"containerImageInputs"`
	// The name of the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// Use this property if you want to use a different name for this file than the default 'taskdef.json'.
	// If you use this property, you don't need to specify the `taskDefinitionTemplateInput` property.
	// Experimental.
	TaskDefinitionTemplateFile awscodepipeline.ArtifactPath `json:"taskDefinitionTemplateFile"`
	// The artifact containing the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// If you use this property, it's assumed the file is called 'taskdef.json'.
	// If your task definition template uses a different filename, leave this property empty,
	// and use the `taskDefinitionTemplateFile` property instead.
	// Experimental.
	TaskDefinitionTemplateInput awscodepipeline.Artifact `json:"taskDefinitionTemplateInput"`
}

// Experimental.
type CodeDeployServerDeployAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCodeDeployServerDeployAction_Override(c CodeDeployServerDeployAction, props *CodeDeployServerDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployServerDeployAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CodeDeployServerDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The CodeDeploy server Deployment Group to deploy to.
	// Experimental.
	DeploymentGroup awscodedeploy.IServerDeploymentGroup `json:"deploymentGroup"`
	// The source to use as input for deployment.
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
}

// A CodePipeline source action for the CodeStar Connections source, which allows connecting to GitHub and BitBucket.
// Experimental.
type CodeStarConnectionsSourceAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewCodeStarConnectionsSourceAction_Override(c CodeStarConnectionsSourceAction, props *CodeStarConnectionsSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeStarConnectionsSourceAction",
		[]interface{}{props},
		c,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type CodeStarConnectionsSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// TODO: EXAMPLE
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	// Experimental.
	ConnectionArn *string `json:"connectionArn"`
	// The output artifact that this action produces.
	//
	// Can be used as input for further pipeline actions.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// The owning user or organization of the repository.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the repository.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Repo *string `json:"repo"`
	// The branch to build.
	// Experimental.
	Branch *string `json:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Experimental.
	TriggerOnPush *bool `json:"triggerOnPush"`
}

// The ECR Repository source CodePipeline Action.
//
// Will trigger the pipeline as soon as the target tag in the repository
// changes, but only if there is a CloudTrail Trail in the account that
// captures the ECR event.
// Experimental.
type EcrSourceAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Variables() *EcrSourceVariables
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewEcrSourceAction_Override(e EcrSourceAction, props *EcrSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcrSourceAction",
		[]interface{}{props},
		e,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type EcrSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// The repository that will be watched for changes.
	// Experimental.
	Repository awsecr.IRepository `json:"repository"`
	// The image tag that will be checked for changes.
	// Experimental.
	ImageTag *string `json:"imageTag"`
}

// The CodePipeline variables emitted by the ECR source Action.
// Experimental.
type EcrSourceVariables struct {
	// The digest of the current image, in the form '<digest type>:<digest value>'.
	// Experimental.
	ImageDigest *string `json:"imageDigest"`
	// The Docker tag of the current image.
	// Experimental.
	ImageTag *string `json:"imageTag"`
	// The full ECR Docker URI of the current image.
	// Experimental.
	ImageUri *string `json:"imageUri"`
	// The identifier of the registry.
	//
	// In ECR, this is usually the ID of the AWS account owning it.
	// Experimental.
	RegistryId *string `json:"registryId"`
	// The physical name of the repository that this action tracks.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
}

// CodePipeline Action to deploy an ECS Service.
// Experimental.
type EcsDeployAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewEcsDeployAction_Override(e EcsDeployAction, props *EcsDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.EcsDeployAction",
		[]interface{}{props},
		e,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type EcsDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The ECS Service to deploy.
	// Experimental.
	Service awsecs.IBaseService `json:"service"`
	// Timeout for the ECS deployment in minutes.
	//
	// Value must be between 1-60.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-ECS.html
	//
	// Experimental.
	DeploymentTimeout awscdk.Duration `json:"deploymentTimeout"`
	// The name of the JSON image definitions file to use for deployments.
	//
	// The JSON file is a list of objects,
	// each with 2 keys: `name` is the name of the container in the Task Definition,
	// and `imageUri` is the Docker image URI you want to update your service with.
	// Use this property if you want to use a different name for this file than the default 'imagedefinitions.json'.
	// If you use this property, you don't need to specify the `input` property.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions
	//
	// Experimental.
	ImageFile awscodepipeline.ArtifactPath `json:"imageFile"`
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
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
}

// Source that is provided by a GitHub repository.
// Experimental.
type GitHubSourceAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Variables() *GitHubSourceVariables
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewGitHubSourceAction_Override(g GitHubSourceAction, props *GitHubSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.GitHubSourceAction",
		[]interface{}{props},
		g,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type GitHubSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
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
	// * **admin:repo_hook** - if you plan to use webhooks (true by default)
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/appendix-github-oauth.html#GitHub-create-personal-token-CLI
	//
	// Experimental.
	OauthToken awscdk.SecretValue `json:"oauthToken"`
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// The GitHub account/user that owns the repo.
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the repo, without the username.
	// Experimental.
	Repo *string `json:"repo"`
	// The branch to use.
	// Experimental.
	Branch *string `json:"branch"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action
	// With "POLL", CodePipeline periodically checks the source for changes
	// With "None", the action is not triggered through changes in the source
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	// Experimental.
	Trigger GitHubTrigger `json:"trigger"`
}

// The CodePipeline variables emitted by GitHub source Action.
// Experimental.
type GitHubSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	// Experimental.
	AuthorDate *string `json:"authorDate"`
	// The name of the branch this action tracks.
	// Experimental.
	BranchName *string `json:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	// Experimental.
	CommitId *string `json:"commitId"`
	// The message of the currently last commit on the tracked branch.
	// Experimental.
	CommitMessage *string `json:"commitMessage"`
	// The date the currently last commit on the tracked branch was committed, in ISO-8601 format.
	// Experimental.
	CommitterDate *string `json:"committerDate"`
	// The GitHub API URL of the currently last commit on the tracked branch.
	// Experimental.
	CommitUrl *string `json:"commitUrl"`
	// The name of the repository this action points to.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
}

// If and how the GitHub source action should be triggered.
// Experimental.
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
// Experimental.
type IJenkinsProvider interface {
	constructs.IConstruct
	// Experimental.
	ProviderName() *string
	// Experimental.
	ServerUrl() *string
	// Experimental.
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
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-four-stage-pipeline.html
//
// Experimental.
type JenkinsAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewJenkinsAction_Override(j JenkinsAction, props *JenkinsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsAction",
		[]interface{}{props},
		j,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type JenkinsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Jenkins Provider for this Action.
	// Experimental.
	JenkinsProvider IJenkinsProvider `json:"jenkinsProvider"`
	// The name of the project (sometimes also called job, or task) on your Jenkins installation that will be invoked by this Action.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ProjectName *string `json:"projectName"`
	// The type of the Action - Build, or Test.
	// Experimental.
	Type JenkinsActionType `json:"type"`
	// The source to use as input for this build.
	// Experimental.
	Inputs *[]awscodepipeline.Artifact `json:"inputs"`
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `json:"outputs"`
}

// The type of the Jenkins Action that determines its CodePipeline Category - Build, or Test.
//
// Note that a Jenkins provider, even if it has the same name,
// must be separately registered for each type.
// Experimental.
type JenkinsActionType string

const (
	JenkinsActionType_BUILD JenkinsActionType = "BUILD"
	JenkinsActionType_TEST JenkinsActionType = "TEST"
)

// A class representing Jenkins providers.
// See: #import
//
// Experimental.
type JenkinsProvider interface {
	BaseJenkinsProvider
	Node() constructs.Node
	ProviderName() *string
	ServerUrl() *string
	Version() *string
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


// Experimental.
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

// Experimental.
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
// Returns: a new Construct representing a reference to an existing Jenkins provider
// Experimental.
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Returns a string representation of this construct.
// Experimental.
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
// Experimental.
type JenkinsProviderAttributes struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ProviderName *string `json:"providerName"`
	// The base URL of your Jenkins server.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ServerUrl *string `json:"serverUrl"`
	// The version of your provider.
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type JenkinsProviderProps struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ProviderName *string `json:"providerName"`
	// The base URL of your Jenkins server.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ServerUrl *string `json:"serverUrl"`
	// Whether to immediately register a Jenkins Provider for the build category.
	//
	// The Provider will always be registered if you create a {@link JenkinsAction}.
	// Experimental.
	ForBuild *bool `json:"forBuild"`
	// Whether to immediately register a Jenkins Provider for the test category.
	//
	// The Provider will always be registered if you create a {@link JenkinsTestAction}.
	// Experimental.
	ForTest *bool `json:"forTest"`
	// The version of your provider.
	// Experimental.
	Version *string `json:"version"`
}

// CodePipeline invoke Action that is provided by an AWS Lambda function.
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html
//
// Experimental.
type LambdaInvokeAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
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


// Experimental.
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

// Experimental.
func NewLambdaInvokeAction_Override(l LambdaInvokeAction, props *LambdaInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.LambdaInvokeAction",
		[]interface{}{props},
		l,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Reference a CodePipeline variable defined by the Lambda function this action points to.
//
// Variables in Lambda invoke actions are defined by calling the PutJobSuccessResult CodePipeline API call
// with the 'outputVariables' property filled.
// See: https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_PutJobSuccessResult.html
//
// Experimental.
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

// Experimental.
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
// Experimental.
type LambdaInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The lambda function to invoke.
	// Experimental.
	Lambda awslambda.IFunction `json:"lambda"`
	// The optional input Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 inputs.
	// The inputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.inputArtifacts` path.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	// Experimental.
	Inputs *[]awscodepipeline.Artifact `json:"inputs"`
	// The optional names of the output Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 outputs.
	// The outputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.outputArtifacts` path.
	// It is the responsibility of the Lambda to upload ZIP files with the Artifact contents to the provided locations.
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `json:"outputs"`
	// A set of key-value pairs that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParameters` or `userParametersString` can be specified.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	// Experimental.
	UserParameters *map[string]interface{} `json:"userParameters"`
	// The string representation of the user parameters that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParametersString` or `userParameters` can be specified.
	// Experimental.
	UserParametersString *string `json:"userParametersString"`
}

// Manual approval action.
// Experimental.
type ManualApprovalAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	NotificationTopic() awssns.ITopic
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	GrantManualApproval(grantable awsiam.IGrantable)
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


// Experimental.
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

// Experimental.
func NewManualApprovalAction_Override(m ManualApprovalAction, props *ManualApprovalActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ManualApprovalAction",
		[]interface{}{props},
		m,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// grant the provided principal the permissions to approve or reject this manual approval action.
//
// For more info see:
// https://docs.aws.amazon.com/codepipeline/latest/userguide/approvals-iam-permissions.html
// Experimental.
func (m *jsiiProxy_ManualApprovalAction) GrantManualApproval(grantable awsiam.IGrantable) {
	_jsii_.InvokeVoid(
		m,
		"grantManualApproval",
		[]interface{}{grantable},
	)
}

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type ManualApprovalActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Any additional information that you want to include in the notification email message.
	// Experimental.
	AdditionalInformation *string `json:"additionalInformation"`
	// URL you want to provide to the reviewer as part of the approval request.
	// Experimental.
	ExternalEntityLink *string `json:"externalEntityLink"`
	// Optional SNS topic to send notifications to when an approval is pending.
	// Experimental.
	NotificationTopic awssns.ITopic `json:"notificationTopic"`
	// A list of email addresses to subscribe to notifications when this Action is pending approval.
	//
	// If this has been provided, but not `notificationTopic`,
	// a new Topic will be created.
	// Experimental.
	NotifyEmails *[]*string `json:"notifyEmails"`
}

// Deploys the sourceArtifact to Amazon S3.
// Experimental.
type S3DeployAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewS3DeployAction_Override(s S3DeployAction, props *S3DeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3DeployAction",
		[]interface{}{props},
		s,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type S3DeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Amazon S3 bucket that is the deploy target.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// The input Artifact to deploy to Amazon S3.
	// Experimental.
	Input awscodepipeline.Artifact `json:"input"`
	// The specified canned ACL to objects deployed to Amazon S3.
	//
	// This overwrites any existing ACL that was applied to the object.
	// Experimental.
	AccessControl awss3.BucketAccessControl `json:"accessControl"`
	// The caching behavior for requests/responses for objects in the bucket.
	//
	// The final cache control property will be the result of joining all of the provided array elements with a comma
	// (plus a space after the comma).
	// Experimental.
	CacheControl *[]CacheControl `json:"cacheControl"`
	// Should the deploy action extract the artifact before deploying to Amazon S3.
	// Experimental.
	Extract *bool `json:"extract"`
	// The key of the target object.
	//
	// This is required if extract is false.
	// Experimental.
	ObjectKey *string `json:"objectKey"`
}

// Source that is provided by a specific Amazon S3 object.
//
// Will trigger the pipeline as soon as the S3 object changes, but only if there is
// a CloudTrail Trail in the account that captures the S3 event.
// Experimental.
type S3SourceAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Variables() *S3SourceVariables
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewS3SourceAction_Override(s S3SourceAction, props *S3SourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.S3SourceAction",
		[]interface{}{props},
		s,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type S3SourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Amazon S3 bucket that stores the source code.
	//
	// If you import an encrypted bucket in your stack, please specify
	// the encryption key at import time by using `Bucket.fromBucketAttributes()` method.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// The key within the S3 bucket that stores the source code.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	BucketKey *string `json:"bucketKey"`
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	// Experimental.
	Trigger S3Trigger `json:"trigger"`
}

// The CodePipeline variables emitted by the S3 source Action.
// Experimental.
type S3SourceVariables struct {
	// The e-tag of the S3 version of the object that triggered the build.
	// Experimental.
	ETag *string `json:"eTag"`
	// The identifier of the S3 version of the object that triggered the build.
	// Experimental.
	VersionId *string `json:"versionId"`
}

// How should the S3 Action detect changes.
//
// This is the type of the {@link S3SourceAction.trigger} property.
// Experimental.
type S3Trigger string

const (
	S3Trigger_NONE S3Trigger = "NONE"
	S3Trigger_POLL S3Trigger = "POLL"
	S3Trigger_EVENTS S3Trigger = "EVENTS"
)

// CodePipeline action to connect to an existing ServiceCatalog product.
//
// **Note**: this class is still experimental, and may have breaking changes in the future!
// Experimental.
type ServiceCatalogDeployActionBeta1 interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewServiceCatalogDeployActionBeta1_Override(s ServiceCatalogDeployActionBeta1, props *ServiceCatalogDeployActionBeta1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		s,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type ServiceCatalogDeployActionBeta1Props struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The identifier of the product in the Service Catalog.
	//
	// This product must already exist.
	// Experimental.
	ProductId *string `json:"productId"`
	// The name of the version of the Service Catalog product to be deployed.
	// Experimental.
	ProductVersionName *string `json:"productVersionName"`
	// The path to the cloudformation artifact.
	// Experimental.
	TemplatePath awscodepipeline.ArtifactPath `json:"templatePath"`
	// The optional description of this version of the Service Catalog product.
	// Experimental.
	ProductVersionDescription *string `json:"productVersionDescription"`
}

// Represents the input for the StateMachine.
// Experimental.
type StateMachineInput interface {
	Input() interface{}
	InputArtifact() awscodepipeline.Artifact
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
// Experimental.
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
// Experimental.
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
// Experimental.
type StepFunctionInvokeAction interface {
	Action
	ActionProperties() *awscodepipeline.ActionProperties
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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


// Experimental.
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

// Experimental.
func NewStepFunctionInvokeAction_Override(s StepFunctionInvokeAction, props *StepFunctionsInvokeActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StepFunctionInvokeAction",
		[]interface{}{props},
		s,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
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

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
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

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
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

// Experimental.
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
// Experimental.
type StepFunctionsInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The state machine to invoke.
	// Experimental.
	StateMachine awsstepfunctions.IStateMachine `json:"stateMachine"`
	// Prefix (optional).
	//
	// By default, the action execution ID is used as the state machine execution name.
	// If a prefix is provided, it is prepended to the action execution ID with a hyphen and
	// together used as the state machine execution name.
	// Experimental.
	ExecutionNamePrefix *string `json:"executionNamePrefix"`
	// The optional output Artifact of the Action.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// Represents the input to the StateMachine.
	//
	// This includes input artifact, input type and the statemachine input.
	// Experimental.
	StateMachineInput StateMachineInput `json:"stateMachineInput"`
}

