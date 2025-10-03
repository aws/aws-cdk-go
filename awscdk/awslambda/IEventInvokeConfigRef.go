package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventInvokeConfig.
// Experimental.
type IEventInvokeConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventInvokeConfigRef
type jsiiProxy_IEventInvokeConfigRef struct {
	internal.Type__constructsIConstruct
}

