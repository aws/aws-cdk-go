package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSubnetGroup.
// Experimental.
type IDBSubnetGroupRef interface {
	constructs.IConstruct
	// A reference to a DBSubnetGroup resource.
	// Experimental.
	DbSubnetGroupRef() *DBSubnetGroupReference
}

// The jsii proxy for IDBSubnetGroupRef
type jsiiProxy_IDBSubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBSubnetGroupRef) DbSubnetGroupRef() *DBSubnetGroupReference {
	var returns *DBSubnetGroupReference
	_jsii_.Get(
		j,
		"dbSubnetGroupRef",
		&returns,
	)
	return returns
}

