package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// A hosted configuration represents configuration stored in the AWS AppConfig hosted configuration store.
//
// Example:
//   var application application
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   })
//
type HostedConfiguration interface {
	constructs.Construct
	IConfiguration
	IExtensible
	// The application associated with the configuration.
	Application() IApplication
	ApplicationId() *string
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the configuration profile.
	ConfigurationProfileArn() *string
	// The ID of the configuration profile.
	ConfigurationProfileId() *string
	// The content of the hosted configuration.
	Content() *string
	// The content type of the hosted configuration.
	ContentType() *string
	// The deployment key for the configuration.
	DeploymentKey() awskms.IKey
	// The deployment strategy for the configuration.
	DeploymentStrategy() IDeploymentStrategy
	// The environments to deploy to.
	DeployTo() *[]IEnvironment
	// The description of the configuration.
	Description() *string
	Extensible() ExtensibleBase
	SetExtensible(val ExtensibleBase)
	// The Amazon Resource Name (ARN) of the hosted configuration version.
	HostedConfigurationVersionArn() *string
	// The latest version number of the hosted configuration.
	LatestVersionNumber() *float64
	// The name of the configuration.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// The configuration type.
	Type() ConfigurationType
	// The validators for the configuration.
	Validators() *[]IValidator
	// The version label of the hosted configuration.
	VersionLabel() *string
	// The version number of the hosted configuration.
	VersionNumber() *string
	AddExistingEnvironmentsToApplication()
	// Adds an extension association to the configuration profile.
	AddExtension(extension IExtension)
	DeployConfigToEnvironments()
	// Adds an extension defined by the action point and event destination and also creates an extension association to the configuration profile.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to the configuration profile.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to the configuration profile.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to the configuration profile.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to the configuration profile.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to the configuration profile.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to the configuration profile.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to the configuration profile.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for HostedConfiguration
type jsiiProxy_HostedConfiguration struct {
	internal.Type__constructsConstruct
	jsiiProxy_IConfiguration
	jsiiProxy_IExtensible
}

func (j *jsiiProxy_HostedConfiguration) Application() IApplication {
	var returns IApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) ConfigurationProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) DeploymentKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"deploymentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) DeploymentStrategy() IDeploymentStrategy {
	var returns IDeploymentStrategy
	_jsii_.Get(
		j,
		"deploymentStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) DeployTo() *[]IEnvironment {
	var returns *[]IEnvironment
	_jsii_.Get(
		j,
		"deployTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Extensible() ExtensibleBase {
	var returns ExtensibleBase
	_jsii_.Get(
		j,
		"extensible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) HostedConfigurationVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedConfigurationVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) LatestVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Type() ConfigurationType {
	var returns ConfigurationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) Validators() *[]IValidator {
	var returns *[]IValidator
	_jsii_.Get(
		j,
		"validators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) VersionLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedConfiguration) VersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}


func NewHostedConfiguration(scope constructs.Construct, id *string, props *HostedConfigurationProps) HostedConfiguration {
	_init_.Initialize()

	if err := validateNewHostedConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostedConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.HostedConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHostedConfiguration_Override(h HostedConfiguration, scope constructs.Construct, id *string, props *HostedConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.HostedConfiguration",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HostedConfiguration)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_HostedConfiguration)SetExtensible(val ExtensibleBase) {
	if err := j.validateSetExtensibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensible",
		val,
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
func HostedConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostedConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.HostedConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostedConfiguration) AddExistingEnvironmentsToApplication() {
	_jsii_.InvokeVoid(
		h,
		"addExistingEnvironmentsToApplication",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostedConfiguration) AddExtension(extension IExtension) {
	if err := h.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addExtension",
		[]interface{}{extension},
	)
}

func (h *jsiiProxy_HostedConfiguration) DeployConfigToEnvironments() {
	_jsii_.InvokeVoid(
		h,
		"deployConfigToEnvironments",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostedConfiguration) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := h.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

func (h *jsiiProxy_HostedConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

