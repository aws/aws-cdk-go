package awsneptunegraph

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateGraphEndpoint.
// Experimental.
type IPrivateGraphEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPrivateGraphEndpointRef
type jsiiProxy_IPrivateGraphEndpointRef struct {
	internal.Type__constructsIConstruct
}

