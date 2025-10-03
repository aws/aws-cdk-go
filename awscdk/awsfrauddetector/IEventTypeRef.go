package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventType.
// Experimental.
type IEventTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventTypeRef
type jsiiProxy_IEventTypeRef struct {
	internal.Type__constructsIConstruct
}

