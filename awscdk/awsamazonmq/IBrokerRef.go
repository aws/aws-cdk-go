package awsamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Broker.
// Experimental.
type IBrokerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBrokerRef
type jsiiProxy_IBrokerRef struct {
	internal.Type__constructsIConstruct
}

