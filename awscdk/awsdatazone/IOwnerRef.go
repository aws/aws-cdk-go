package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Owner.
// Experimental.
type IOwnerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOwnerRef
type jsiiProxy_IOwnerRef struct {
	internal.Type__constructsIConstruct
}

