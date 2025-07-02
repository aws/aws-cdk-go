package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type IEnvironment interface {
	awscdk.IResource
	// Creates a deployment of the supplied configuration to this environment.
	//
	// Note that you can only deploy one configuration at a time to an environment.
	// However, you can deploy one configuration each to different environments at the same time.
	// If more than one deployment is requested for this environment, they will occur in the same order they were provided.
	AddDeployment(configuration IConfiguration)
	// Creates a deployment for each of the supplied configurations to this environment.
	//
	// These configurations will be deployed in the same order as the input array.
	AddDeployments(configurations ...IConfiguration)
	// Adds an extension association to the environment.
	AddExtension(extension IExtension)
	// Adds an AT_DEPLOYMENT_TICK extension with the provided event destination and also creates an extension association to an application.
	AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an IAM policy statement associated with this environment to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM principal to perform read operations on this environment's configurations.
	//
	// Actions: GetLatestConfiguration, StartConfigurationSession.
	GrantReadConfig(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an extension defined by the action point and event destination and also creates an extension association to the environment.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to the environment.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to the environment.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to the environment.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to the environment.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to the environment.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to the environment.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to the environment.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
	// The application associated with the environment.
	Application() IApplication
	// The ID of the application associated to the environment.
	ApplicationId() *string
	// The description of the environment.
	Description() *string
	// The Amazon Resource Name (ARN) of the environment.
	EnvironmentArn() *string
	// The ID of the environment.
	EnvironmentId() *string
	// The monitors for the environment.
	Monitors() *[]Monitor
	// The name of the environment.
	Name() *string
}

// The jsii proxy for IEnvironment
type jsiiProxy_IEnvironment struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEnvironment) AddDeployment(configuration IConfiguration) {
	if err := i.validateAddDeploymentParameters(configuration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDeployment",
		[]interface{}{configuration},
	)
}

func (i *jsiiProxy_IEnvironment) AddDeployments(configurations ...IConfiguration) {
	args := []interface{}{}
	for _, a := range configurations {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addDeployments",
		args,
	)
}

func (i *jsiiProxy_IEnvironment) AddExtension(extension IExtension) {
	if err := i.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addExtension",
		[]interface{}{extension},
	)
}

func (i *jsiiProxy_IEnvironment) AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateAtDeploymentTickParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"atDeploymentTick",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironment) GrantReadConfig(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadConfigParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadConfig",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironment) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IEnvironment) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

func (j *jsiiProxy_IEnvironment) Application() IApplication {
	var returns IApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Monitors() *[]Monitor {
	var returns *[]Monitor
	_jsii_.Get(
		j,
		"monitors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

