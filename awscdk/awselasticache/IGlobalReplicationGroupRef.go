package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalReplicationGroup.
// Experimental.
type IGlobalReplicationGroupRef interface {
	constructs.IConstruct
	// A reference to a GlobalReplicationGroup resource.
	// Experimental.
	GlobalReplicationGroupRef() *GlobalReplicationGroupReference
}

// The jsii proxy for IGlobalReplicationGroupRef
type jsiiProxy_IGlobalReplicationGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGlobalReplicationGroupRef) GlobalReplicationGroupRef() *GlobalReplicationGroupReference {
	var returns *GlobalReplicationGroupReference
	_jsii_.Get(
		j,
		"globalReplicationGroupRef",
		&returns,
	)
	return returns
}

