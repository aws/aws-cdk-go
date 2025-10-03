package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowVpcInterface.
// Experimental.
type IFlowVpcInterfaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowVpcInterfaceRef
type jsiiProxy_IFlowVpcInterfaceRef struct {
	internal.Type__constructsIConstruct
}

