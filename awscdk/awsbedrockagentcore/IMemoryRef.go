package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Memory.
// Experimental.
type IMemoryRef interface {
	constructs.IConstruct
	// A reference to a Memory resource.
	// Experimental.
	MemoryRef() *MemoryReference
}

// The jsii proxy for IMemoryRef
type jsiiProxy_IMemoryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMemoryRef) MemoryRef() *MemoryReference {
	var returns *MemoryReference
	_jsii_.Get(
		j,
		"memoryRef",
		&returns,
	)
	return returns
}

