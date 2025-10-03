package awspinpointemail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DedicatedIpPool.
// Experimental.
type IDedicatedIpPoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDedicatedIpPoolRef
type jsiiProxy_IDedicatedIpPoolRef struct {
	internal.Type__constructsIConstruct
}

