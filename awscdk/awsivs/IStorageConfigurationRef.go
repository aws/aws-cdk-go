package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageConfiguration.
// Experimental.
type IStorageConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StorageConfiguration resource.
	// Experimental.
	StorageConfigurationRef() *StorageConfigurationReference
}

// The jsii proxy for IStorageConfigurationRef
type jsiiProxy_IStorageConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStorageConfigurationRef) StorageConfigurationRef() *StorageConfigurationReference {
	var returns *StorageConfigurationReference
	_jsii_.Get(
		j,
		"storageConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

