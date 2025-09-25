package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalTable.
// Experimental.
type IGlobalTableRef interface {
	constructs.IConstruct
	// A reference to a GlobalTable resource.
	// Experimental.
	GlobalTableRef() *GlobalTableReference
}

// The jsii proxy for IGlobalTableRef
type jsiiProxy_IGlobalTableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGlobalTableRef) GlobalTableRef() *GlobalTableReference {
	var returns *GlobalTableReference
	_jsii_.Get(
		j,
		"globalTableRef",
		&returns,
	)
	return returns
}

