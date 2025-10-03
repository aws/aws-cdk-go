package awsmanagedblockchain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Accessor.
// Experimental.
type IAccessorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessorRef
type jsiiProxy_IAccessorRef struct {
	internal.Type__constructsIConstruct
}

