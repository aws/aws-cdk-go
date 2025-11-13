package interfacesawsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBCluster.
// Experimental.
type IDBClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBCluster resource.
	// Experimental.
	DbClusterRef() *DBClusterReference
}

// The jsii proxy for IDBClusterRef
type jsiiProxy_IDBClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDBClusterRef) DbClusterRef() *DBClusterReference {
	var returns *DBClusterReference
	_jsii_.Get(
		j,
		"dbClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

