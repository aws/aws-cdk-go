package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBShardGroup.
// Experimental.
type IDBShardGroupRef interface {
	constructs.IConstruct
	// A reference to a DBShardGroup resource.
	// Experimental.
	DbShardGroupRef() *DBShardGroupReference
}

// The jsii proxy for IDBShardGroupRef
type jsiiProxy_IDBShardGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBShardGroupRef) DbShardGroupRef() *DBShardGroupReference {
	var returns *DBShardGroupReference
	_jsii_.Get(
		j,
		"dbShardGroupRef",
		&returns,
	)
	return returns
}

