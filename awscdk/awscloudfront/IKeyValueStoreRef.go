package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyValueStore.
// Experimental.
type IKeyValueStoreRef interface {
	constructs.IConstruct
	// A reference to a KeyValueStore resource.
	// Experimental.
	KeyValueStoreRef() *KeyValueStoreReference
}

// The jsii proxy for IKeyValueStoreRef
type jsiiProxy_IKeyValueStoreRef struct {
	internal.Type__constructsIConstruct
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

