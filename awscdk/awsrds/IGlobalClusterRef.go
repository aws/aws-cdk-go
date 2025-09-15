package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalCluster.
// Experimental.
type IGlobalClusterRef interface {
	constructs.IConstruct
	// A reference to a GlobalCluster resource.
	// Experimental.
	GlobalClusterRef() *GlobalClusterReference
}

// The jsii proxy for IGlobalClusterRef
type jsiiProxy_IGlobalClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGlobalClusterRef) GlobalClusterRef() *GlobalClusterReference {
	var returns *GlobalClusterReference
	_jsii_.Get(
		j,
		"globalClusterRef",
		&returns,
	)
	return returns
}

