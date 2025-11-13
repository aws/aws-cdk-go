package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBInstance.
// Experimental.
type IDBInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBInstance resource.
	// Experimental.
	DbInstanceRef() *DBInstanceReference
}

// The jsii proxy for IDBInstanceRef
type jsiiProxy_IDBInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDBInstanceRef) DbInstanceRef() *DBInstanceReference {
	var returns *DBInstanceReference
	_jsii_.Get(
		j,
		"dbInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

