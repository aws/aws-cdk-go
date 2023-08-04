package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// The interface representing a database cluster (as opposed to instance) engine.
type IClusterEngine interface {
	IEngine
	// Method called when the engine is used to create a new cluster.
	BindToCluster(scope constructs.Construct, options *ClusterEngineBindOptions) *ClusterEngineConfig
	// Whether the IAM Roles used for data importing and exporting need to be combined for this Engine, or can they be kept separate.
	// Default: false.
	//
	CombineImportAndExportRoles() *bool
	// The application used by this engine to perform rotation for a multi-user scenario.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The log types that are available with this engine type.
	SupportedLogTypes() *[]*string
}

// The jsii proxy for IClusterEngine
type jsiiProxy_IClusterEngine struct {
	jsiiProxy_IEngine
}

func (i *jsiiProxy_IClusterEngine) BindToCluster(scope constructs.Construct, options *ClusterEngineBindOptions) *ClusterEngineConfig {
	if err := i.validateBindToClusterParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *ClusterEngineConfig

	_jsii_.Invoke(
		i,
		"bindToCluster",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IClusterEngine) CombineImportAndExportRoles() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"combineImportAndExportRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterEngine) MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"multiUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterEngine) SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"singleUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterEngine) SupportedLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLogTypes",
		&returns,
	)
	return returns
}

