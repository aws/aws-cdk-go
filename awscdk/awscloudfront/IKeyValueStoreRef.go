package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyValueStore.
// Experimental.
type IKeyValueStoreRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a KeyValueStore resource.
	// Experimental.
	KeyValueStoreRef() *KeyValueStoreReference
}

// The jsii proxy for IKeyValueStoreRef
type jsiiProxy_IKeyValueStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IKeyValueStoreRef) KeyValueStoreRef() *KeyValueStoreReference {
	var returns *KeyValueStoreReference
	_jsii_.Get(
		j,
		"keyValueStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStoreRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

