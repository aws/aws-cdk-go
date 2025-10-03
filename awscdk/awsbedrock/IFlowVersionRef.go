package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowVersion.
// Experimental.
type IFlowVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowVersionRef
type jsiiProxy_IFlowVersionRef struct {
	internal.Type__constructsIConstruct
}

