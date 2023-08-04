package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface representing a database instance (as opposed to cluster) engine.
type IInstanceEngine interface {
	IEngine
	// Method called when the engine is used to create a new instance.
	BindToInstance(scope constructs.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// Whether this engine supports automatic backups of a read replica instance.
	// Default: false.
	//
	SupportsReadReplicaBackups() *bool
}

// The jsii proxy for IInstanceEngine
type jsiiProxy_IInstanceEngine struct {
	jsiiProxy_IEngine
}

func (i *jsiiProxy_IInstanceEngine) BindToInstance(scope constructs.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig {
	if err := i.validateBindToInstanceParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *InstanceEngineConfig

	_jsii_.Invoke(
		i,
		"bindToInstance",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInstanceEngine) MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"multiUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceEngine) SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"singleUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceEngine) SupportsReadReplicaBackups() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsReadReplicaBackups",
		&returns,
	)
	return returns
}

