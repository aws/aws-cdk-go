package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowAlias.
// Experimental.
type IFlowAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowAliasRef
type jsiiProxy_IFlowAliasRef struct {
	internal.Type__constructsIConstruct
}

