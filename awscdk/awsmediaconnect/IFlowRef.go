package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Flow.
// Experimental.
type IFlowRef interface {
	constructs.IConstruct
	// A reference to a Flow resource.
	// Experimental.
	FlowRef() *FlowReference
}

// The jsii proxy for IFlowRef
type jsiiProxy_IFlowRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFlowRef) FlowRef() *FlowReference {
	var returns *FlowReference
	_jsii_.Get(
		j,
		"flowRef",
		&returns,
	)
	return returns
}

