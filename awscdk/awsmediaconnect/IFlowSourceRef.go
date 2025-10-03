package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowSource.
// Experimental.
type IFlowSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowSourceRef
type jsiiProxy_IFlowSourceRef struct {
	internal.Type__constructsIConstruct
}

