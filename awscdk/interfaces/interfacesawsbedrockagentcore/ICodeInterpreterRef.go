package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeInterpreter.
// Experimental.
type ICodeInterpreterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CodeInterpreter resource.
	// Experimental.
	CodeInterpreterRef() *CodeInterpreterReference
}

// The jsii proxy for ICodeInterpreterRef
type jsiiProxy_ICodeInterpreterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICodeInterpreterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICodeInterpreterRef) CodeInterpreterRef() *CodeInterpreterReference {
	var returns *CodeInterpreterReference
	_jsii_.Get(
		j,
		"codeInterpreterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

