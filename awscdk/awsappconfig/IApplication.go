package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
)

type IApplication interface {
	awscdk.IResource
	// Adds an environment.
	AddEnvironment(id *string, options *EnvironmentOptions) IEnvironment
	// Adds an existing environment.
	AddExistingEnvironment(environment IEnvironment)
	// Adds an extension association to the application.
	AddExtension(extension IExtension)
	// Adds a hosted configuration.
	AddHostedConfiguration(id *string, options *HostedConfigurationOptions) HostedConfiguration
	// Adds a sourced configuration.
	AddSourcedConfiguration(id *string, options *SourcedConfigurationOptions) SourcedConfiguration
	// Adds an AT_DEPLOYMENT_TICK extension with the provided event destination and also creates an extension association to an application.
	AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions)
	// Returns the list of associated environments.
	Environments() *[]IEnvironment
	// Adds an extension defined by the action point and event destination and also creates an extension association to an application.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to an application.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to an application.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to an application.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to an application.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to an application.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to an application.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to an application.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
	// The Amazon Resource Name (ARN) of the application.
	ApplicationArn() *string
	// The ID of the application.
	ApplicationId() *string
	// The description of the application.
	Description() *string
	// The name of the application.
	Name() *string
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AddEnvironment(id *string, options *EnvironmentOptions) IEnvironment {
	if err := i.validateAddEnvironmentParameters(id, options); err != nil {
		panic(err)
	}
	var returns IEnvironment

	_jsii_.Invoke(
		i,
		"addEnvironment",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) AddExistingEnvironment(environment IEnvironment) {
	if err := i.validateAddExistingEnvironmentParameters(environment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addExistingEnvironment",
		[]interface{}{environment},
	)
}

func (i *jsiiProxy_IApplication) AddExtension(extension IExtension) {
	if err := i.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addExtension",
		[]interface{}{extension},
	)
}

func (i *jsiiProxy_IApplication) AddHostedConfiguration(id *string, options *HostedConfigurationOptions) HostedConfiguration {
	if err := i.validateAddHostedConfigurationParameters(id, options); err != nil {
		panic(err)
	}
	var returns HostedConfiguration

	_jsii_.Invoke(
		i,
		"addHostedConfiguration",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) AddSourcedConfiguration(id *string, options *SourcedConfigurationOptions) SourcedConfiguration {
	if err := i.validateAddSourcedConfigurationParameters(id, options); err != nil {
		panic(err)
	}
	var returns SourcedConfiguration

	_jsii_.Invoke(
		i,
		"addSourcedConfiguration",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateAtDeploymentTickParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"atDeploymentTick",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) Environments() *[]IEnvironment {
	var returns *[]IEnvironment

	_jsii_.Invoke(
		i,
		"environments",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IApplication) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

func (j *jsiiProxy_IApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

