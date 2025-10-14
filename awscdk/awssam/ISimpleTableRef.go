package awssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimpleTable.
// Experimental.
type ISimpleTableRef interface {
	constructs.IConstruct
	// A reference to a SimpleTable resource.
	// Experimental.
	SimpleTableRef() *SimpleTableReference
}

// The jsii proxy for ISimpleTableRef
type jsiiProxy_ISimpleTableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISimpleTableRef) SimpleTableRef() *SimpleTableReference {
	var returns *SimpleTableReference
	_jsii_.Get(
		j,
		"simpleTableRef",
		&returns,
	)
	return returns
}

