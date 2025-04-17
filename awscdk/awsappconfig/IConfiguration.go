package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

type IConfiguration interface {
	constructs.IConstruct
	// The application associated with the configuration.
	Application() IApplication
	// The ID of the configuration profile.
	ConfigurationProfileId() *string
	// The deployment key for the configuration.
	DeploymentKey() awskms.IKey
	// The deployment strategy for the configuration.
	DeploymentStrategy() IDeploymentStrategy
	// The environments to deploy to.
	DeployTo() *[]IEnvironment
	// The description of the configuration.
	Description() *string
	// The name of the configuration.
	Name() *string
	// The configuration type.
	Type() ConfigurationType
	// The validators for the configuration.
	Validators() *[]IValidator
	// The configuration version number.
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

