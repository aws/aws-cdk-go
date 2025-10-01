package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowLog.
// Experimental.
type IFlowLogRef interface {
	constructs.IConstruct
	// A reference to a FlowLog resource.
	// Experimental.
	FlowLogRef() *FlowLogReference
}

// The jsii proxy for IFlowLogRef
type jsiiProxy_IFlowLogRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFlowLogRef) FlowLogRef() *FlowLogReference {
	var returns *FlowLogReference
	_jsii_.Get(
		j,
		"flowLogRef",
		&returns,
	)
	return returns
}

