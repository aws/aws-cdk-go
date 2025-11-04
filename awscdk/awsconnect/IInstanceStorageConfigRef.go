package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceStorageConfig.
// Experimental.
type IInstanceStorageConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InstanceStorageConfig resource.
	// Experimental.
	InstanceStorageConfigRef() *InstanceStorageConfigReference
}

// The jsii proxy for IInstanceStorageConfigRef
type jsiiProxy_IInstanceStorageConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInstanceStorageConfigRef) InstanceStorageConfigRef() *InstanceStorageConfigReference {
	var returns *InstanceStorageConfigReference
	_jsii_.Get(
		j,
		"instanceStorageConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceStorageConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceStorageConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

