package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBus.
// Experimental.
type IEventBusRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventBusRef
type jsiiProxy_IEventBusRef struct {
	internal.Type__constructsIConstruct
}

