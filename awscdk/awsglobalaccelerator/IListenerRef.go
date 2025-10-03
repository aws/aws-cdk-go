package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Listener.
// Experimental.
type IListenerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IListenerRef
type jsiiProxy_IListenerRef struct {
	internal.Type__constructsIConstruct
}

