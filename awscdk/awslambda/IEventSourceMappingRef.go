package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventSourceMapping.
// Experimental.
type IEventSourceMappingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventSourceMappingRef
type jsiiProxy_IEventSourceMappingRef struct {
	internal.Type__constructsIConstruct
}

