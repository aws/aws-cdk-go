package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetwork.
// Experimental.
type IServiceNetworkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceNetworkRef
type jsiiProxy_IServiceNetworkRef struct {
	internal.Type__constructsIConstruct
}

