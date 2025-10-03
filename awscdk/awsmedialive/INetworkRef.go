package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Network.
// Experimental.
type INetworkRef interface {
	constructs.IConstruct
}

// The jsii proxy for INetworkRef
type jsiiProxy_INetworkRef struct {
	internal.Type__constructsIConstruct
}

