package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowOutput.
// Experimental.
type IFlowOutputRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowOutputRef
type jsiiProxy_IFlowOutputRef struct {
	internal.Type__constructsIConstruct
}

