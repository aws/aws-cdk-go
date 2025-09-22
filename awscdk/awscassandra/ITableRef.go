package awscassandra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Table.
// Experimental.
type ITableRef interface {
	constructs.IConstruct
	// A reference to a Table resource.
	// Experimental.
	TableRef() *TableReference
}

// The jsii proxy for ITableRef
type jsiiProxy_ITableRef struct {
	internal.Type__constructsIConstruct
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

