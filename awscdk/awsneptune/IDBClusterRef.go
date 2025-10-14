package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBCluster.
// Experimental.
type IDBClusterRef interface {
	constructs.IConstruct
	// A reference to a DBCluster resource.
	// Experimental.
	DbClusterRef() *DBClusterReference
}

// The jsii proxy for IDBClusterRef
type jsiiProxy_IDBClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBClusterRef) DbClusterRef() *DBClusterReference {
	var returns *DBClusterReference
	_jsii_.Get(
		j,
		"dbClusterRef",
		&returns,
	)
	return returns
}

