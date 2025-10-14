package awscassandra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Keyspace.
// Experimental.
type IKeyspaceRef interface {
	constructs.IConstruct
	// A reference to a Keyspace resource.
	// Experimental.
	KeyspaceRef() *KeyspaceReference
}

// The jsii proxy for IKeyspaceRef
type jsiiProxy_IKeyspaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKeyspaceRef) KeyspaceRef() *KeyspaceReference {
	var returns *KeyspaceReference
	_jsii_.Get(
		j,
		"keyspaceRef",
		&returns,
	)
	return returns
}

