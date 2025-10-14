package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Prompt.
// Experimental.
type IPromptRef interface {
	constructs.IConstruct
	// A reference to a Prompt resource.
	// Experimental.
	PromptRef() *PromptReference
}

// The jsii proxy for IPromptRef
type jsiiProxy_IPromptRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPromptRef) PromptRef() *PromptReference {
	var returns *PromptReference
	_jsii_.Get(
		j,
		"promptRef",
		&returns,
	)
	return returns
}

