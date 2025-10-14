package awspcs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspcs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeNodeGroup.
// Experimental.
type IComputeNodeGroupRef interface {
	constructs.IConstruct
	// A reference to a ComputeNodeGroup resource.
	// Experimental.
	ComputeNodeGroupRef() *ComputeNodeGroupReference
}

// The jsii proxy for IComputeNodeGroupRef
type jsiiProxy_IComputeNodeGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IComputeNodeGroupRef) ComputeNodeGroupRef() *ComputeNodeGroupReference {
	var returns *ComputeNodeGroupReference
	_jsii_.Get(
		j,
		"computeNodeGroupRef",
		&returns,
	)
	return returns
}

