package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackSet.
// Experimental.
type IStackSetRef interface {
	constructs.IConstruct
	// A reference to a StackSet resource.
	// Experimental.
	StackSetRef() *StackSetReference
}

// The jsii proxy for IStackSetRef
type jsiiProxy_IStackSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStackSetRef) StackSetRef() *StackSetReference {
	var returns *StackSetReference
	_jsii_.Get(
		j,
		"stackSetRef",
		&returns,
	)
	return returns
}

