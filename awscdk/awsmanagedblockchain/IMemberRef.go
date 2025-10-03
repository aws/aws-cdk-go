package awsmanagedblockchain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Member.
// Experimental.
type IMemberRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMemberRef
type jsiiProxy_IMemberRef struct {
	internal.Type__constructsIConstruct
}

