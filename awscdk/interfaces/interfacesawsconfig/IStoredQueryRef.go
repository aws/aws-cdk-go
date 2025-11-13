package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StoredQuery.
// Experimental.
type IStoredQueryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StoredQuery resource.
	// Experimental.
	StoredQueryRef() *StoredQueryReference
}

// The jsii proxy for IStoredQueryRef
type jsiiProxy_IStoredQueryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStoredQueryRef) StoredQueryRef() *StoredQueryReference {
	var returns *StoredQueryReference
	_jsii_.Get(
		j,
		"storedQueryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStoredQueryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStoredQueryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

