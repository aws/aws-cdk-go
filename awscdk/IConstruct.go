package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a construct.
// Experimental.
type IConstruct interface {
	constructs.IConstruct
	IDependable
	// The construct tree node for this construct.
	// Experimental.
	Node() ConstructNode
}

// The jsii proxy for IConstruct
type jsiiProxy_IConstruct struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IDependable
}

func (j *jsiiProxy_IConstruct) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

