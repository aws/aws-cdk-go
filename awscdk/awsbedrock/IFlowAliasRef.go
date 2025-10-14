package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowAlias.
// Experimental.
type IFlowAliasRef interface {
	constructs.IConstruct
	// A reference to a FlowAlias resource.
	// Experimental.
	FlowAliasRef() *FlowAliasReference
}

// The jsii proxy for IFlowAliasRef
type jsiiProxy_IFlowAliasRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFlowAliasRef) FlowAliasRef() *FlowAliasReference {
	var returns *FlowAliasReference
	_jsii_.Get(
		j,
		"flowAliasRef",
		&returns,
	)
	return returns
}

