package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowLog.
// Experimental.
type IFlowLogRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowLogRef
type jsiiProxy_IFlowLogRef struct {
	internal.Type__constructsIConstruct
}

