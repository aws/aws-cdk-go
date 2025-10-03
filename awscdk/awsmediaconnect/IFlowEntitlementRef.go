package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowEntitlement.
// Experimental.
type IFlowEntitlementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowEntitlementRef
type jsiiProxy_IFlowEntitlementRef struct {
	internal.Type__constructsIConstruct
}

