package interfacesawsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Script.
// Experimental.
type IScriptRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Script resource.
	// Experimental.
	ScriptRef() *ScriptReference
}

// The jsii proxy for IScriptRef
type jsiiProxy_IScriptRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IScriptRef) ScriptRef() *ScriptReference {
	var returns *ScriptReference
	_jsii_.Get(
		j,
		"scriptRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScriptRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScriptRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

