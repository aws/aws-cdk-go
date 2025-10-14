package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowSource.
// Experimental.
type IFlowSourceRef interface {
	constructs.IConstruct
	// A reference to a FlowSource resource.
	// Experimental.
	FlowSourceRef() *FlowSourceReference
}

// The jsii proxy for IFlowSourceRef
type jsiiProxy_IFlowSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFlowSourceRef) FlowSourceRef() *FlowSourceReference {
	var returns *FlowSourceReference
	_jsii_.Get(
		j,
		"flowSourceRef",
		&returns,
	)
	return returns
}

