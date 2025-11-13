package interfacesawsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAccessor.
// Experimental.
type IDataAccessorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataAccessor resource.
	// Experimental.
	DataAccessorRef() *DataAccessorReference
}

// The jsii proxy for IDataAccessorRef
type jsiiProxy_IDataAccessorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataAccessorRef) DataAccessorRef() *DataAccessorReference {
	var returns *DataAccessorReference
	_jsii_.Get(
		j,
		"dataAccessorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAccessorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAccessorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

