package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCluster.
// Experimental.
type IServerlessClusterRef interface {
	constructs.IConstruct
	// A reference to a ServerlessCluster resource.
	// Experimental.
	ServerlessClusterRef() *ServerlessClusterReference
}

// The jsii proxy for IServerlessClusterRef
type jsiiProxy_IServerlessClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServerlessClusterRef) ServerlessClusterRef() *ServerlessClusterReference {
	var returns *ServerlessClusterReference
	_jsii_.Get(
		j,
		"serverlessClusterRef",
		&returns,
	)
	return returns
}

