package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Memory.
// Experimental.
type IMemoryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Memory resource.
	// Experimental.
	MemoryRef() *MemoryReference
}

// The jsii proxy for IMemoryRef
type jsiiProxy_IMemoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IMemoryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

