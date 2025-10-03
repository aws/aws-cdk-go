package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCluster.
// Experimental.
type IServerlessClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServerlessClusterRef
type jsiiProxy_IServerlessClusterRef struct {
	internal.Type__constructsIConstruct
}

