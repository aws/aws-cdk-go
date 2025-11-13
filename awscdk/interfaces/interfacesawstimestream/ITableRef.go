package interfacesawstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Table.
// Experimental.
type ITableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Table resource.
	// Experimental.
	TableRef() *TableReference
}

// The jsii proxy for ITableRef
type jsiiProxy_ITableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITableRef) TableRef() *TableReference {
	var returns *TableReference
	_jsii_.Get(
		j,
		"tableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

