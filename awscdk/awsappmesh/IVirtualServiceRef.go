package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualService.
// Experimental.
type IVirtualServiceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualServiceRef
type jsiiProxy_IVirtualServiceRef struct {
	internal.Type__constructsIConstruct
}

