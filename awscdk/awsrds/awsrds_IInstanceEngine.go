package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Interface representing a database instance (as opposed to cluster) engine.
// Experimental.
type IInstanceEngine interface {
	IEngine
	// Method called when the engine is used to create a new instance.
	// Experimental.
	BindToInstance(scope awscdk.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	// Experimental.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	// Experimental.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// Whether this engine supports automatic backups of a read replica instance.
	// Experimental.
	SupportsReadReplicaBackups() *bool
}

// The jsii proxy for IInstanceEngine
type jsiiProxy_IInstanceEngine struct {
	jsiiProxy_IEngine
}

func (i *jsiiProxy_IInstanceEngine) BindToInstance(scope awscdk.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig {
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

