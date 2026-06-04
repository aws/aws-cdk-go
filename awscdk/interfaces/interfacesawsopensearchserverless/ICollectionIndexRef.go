package interfacesawsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CollectionIndex.
// Experimental.
type ICollectionIndexRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CollectionIndex resource.
	// Experimental.
	CollectionIndexRef() *CollectionIndexReference
}

// The jsii proxy for ICollectionIndexRef
type jsiiProxy_ICollectionIndexRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICollectionIndexRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICollectionIndexRef) CollectionIndexRef() *CollectionIndexReference {
	var returns *CollectionIndexReference
	_jsii_.Get(
		j,
		"collectionIndexRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICollectionIndexRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICollectionIndexRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

