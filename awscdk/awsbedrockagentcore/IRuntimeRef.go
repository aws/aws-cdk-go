package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Runtime.
// Experimental.
type IRuntimeRef interface {
	constructs.IConstruct
	// A reference to a Runtime resource.
	// Experimental.
	RuntimeRef() *RuntimeReference
}

// The jsii proxy for IRuntimeRef
type jsiiProxy_IRuntimeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRuntimeRef) RuntimeRef() *RuntimeReference {
	var returns *RuntimeReference
	_jsii_.Get(
		j,
		"runtimeRef",
		&returns,
	)
	return returns
}

