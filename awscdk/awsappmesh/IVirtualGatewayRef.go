package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualGateway.
// Experimental.
type IVirtualGatewayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualGatewayRef
type jsiiProxy_IVirtualGatewayRef struct {
	internal.Type__constructsIConstruct
}

