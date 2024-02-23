package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated.
type IConfiguration interface {
	constructs.IConstruct
	// The application associated with the configuration.
	// Deprecated.
	Application() IApplication
	// The ID of the configuration profile.
	// Deprecated.
	ConfigurationProfileId() *string
	// The deployment key for the configuration.
	// Deprecated.
	DeploymentKey() awskms.IKey
	// The deployment strategy for the configuration.
	// Deprecated.
	DeploymentStrategy() IDeploymentStrategy
	// The environments to deploy to.
	// Deprecated.
	DeployTo() *[]IEnvironment
	// The description of the configuration.
	// Deprecated.
	Description() *string
	// The name of the configuration.
	// Deprecated.
	Name() *string
	// The configuration type.
	// Deprecated.
	Type() ConfigurationType
	// The validators for the configuration.
	// Deprecated.
	Validators() *[]IValidator
	// The configuration version number.
	// Deprecated.
	VersionNumber() *string
}

// The jsii proxy for IConfiguration
type jsiiProxy_IConfiguration struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfiguration) Application() IApplication {
	var returns IApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) DeploymentKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"deploymentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) DeploymentStrategy() IDeploymentStrategy {
	var returns IDeploymentStrategy
	_jsii_.Get(
		j,
		"deploymentStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) DeployTo() *[]IEnvironment {
	var returns *[]IEnvironment
	_jsii_.Get(
		j,
		"deployTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) Type() ConfigurationType {
	var returns ConfigurationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) Validators() *[]IValidator {
	var returns *[]IValidator
	_jsii_.Get(
		j,
		"validators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration) VersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}

