package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBParameterGroup.
// Experimental.
type IDBParameterGroupRef interface {
	constructs.IConstruct
	// A reference to a DBParameterGroup resource.
	// Experimental.
	DbParameterGroupRef() *DBParameterGroupReference
}

// The jsii proxy for IDBParameterGroupRef
type jsiiProxy_IDBParameterGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBParameterGroupRef) DbParameterGroupRef() *DBParameterGroupReference {
	var returns *DBParameterGroupReference
	_jsii_.Get(
		j,
		"dbParameterGroupRef",
		&returns,
	)
	return returns
}

