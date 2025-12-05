package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataTable.
// Experimental.
type IDataTableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataTable resource.
	// Experimental.
	DataTableRef() *DataTableReference
}

// The jsii proxy for IDataTableRef
type jsiiProxy_IDataTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataTableRef) DataTableRef() *DataTableReference {
	var returns *DataTableReference
	_jsii_.Get(
		j,
		"dataTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataTableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

