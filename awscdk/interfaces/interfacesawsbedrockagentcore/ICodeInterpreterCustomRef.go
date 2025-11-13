package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeInterpreterCustom.
// Experimental.
type ICodeInterpreterCustomRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CodeInterpreterCustom resource.
	// Experimental.
	CodeInterpreterCustomRef() *CodeInterpreterCustomReference
}

// The jsii proxy for ICodeInterpreterCustomRef
type jsiiProxy_ICodeInterpreterCustomRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICodeInterpreterCustomRef) CodeInterpreterCustomRef() *CodeInterpreterCustomReference {
	var returns *CodeInterpreterCustomReference
	_jsii_.Get(
		j,
		"codeInterpreterCustomRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustomRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustomRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

