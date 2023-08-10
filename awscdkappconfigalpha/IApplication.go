package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
)

// Experimental.
type IApplication interface {
	awscdk.IResource
	// Adds an environment.
	// Experimental.
	AddEnvironment(id *string, options *EnvironmentOptions) IEnvironment
	// Adds an existing environment.
	// Experimental.
	AddExistingEnvironment(environment IEnvironment)
	// Adds a hosted configuration.
	// Experimental.
	AddHostedConfiguration(id *string, options *HostedConfigurationOptions) HostedConfiguration
	// Adds a sourced configuration.
	// Experimental.
	AddSourcedConfiguration(id *string, options *SourcedConfigurationOptions) SourcedConfiguration
	// The Amazon Resource Name (ARN) of the application.
	// Experimental.
	ApplicationArn() *string
	// The ID of the application.
	// Experimental.
	ApplicationId() *string
	// The description of the application.
	// Experimental.
	Description() *string
	// Returns the list of associated environments.
	// Experimental.
	Environments() *[]IEnvironment
	// The name of the application.
	// Experimental.
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

func (j *jsiiProxy_IApplication) Environments() *[]IEnvironment {
	var returns *[]IEnvironment
	_jsii_.Get(
		j,
		"environments",
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

