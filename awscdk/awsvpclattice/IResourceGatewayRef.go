package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceGateway.
// Experimental.
type IResourceGatewayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceGatewayRef
type jsiiProxy_IResourceGatewayRef struct {
	internal.Type__constructsIConstruct
}

