package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBusPolicy.
// Experimental.
type IEventBusPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventBusPolicyRef
type jsiiProxy_IEventBusPolicyRef struct {
	internal.Type__constructsIConstruct
}

