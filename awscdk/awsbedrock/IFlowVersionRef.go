package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowVersion.
// Experimental.
type IFlowVersionRef interface {
	constructs.IConstruct
	// A reference to a FlowVersion resource.
	// Experimental.
	FlowVersionRef() *FlowVersionReference
}

// The jsii proxy for IFlowVersionRef
type jsiiProxy_IFlowVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFlowVersionRef) FlowVersionRef() *FlowVersionReference {
	var returns *FlowVersionReference
	_jsii_.Get(
		j,
		"flowVersionRef",
		&returns,
	)
	return returns
}

