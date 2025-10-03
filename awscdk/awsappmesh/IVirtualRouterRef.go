package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualRouter.
// Experimental.
type IVirtualRouterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualRouterRef
type jsiiProxy_IVirtualRouterRef struct {
	internal.Type__constructsIConstruct
}

