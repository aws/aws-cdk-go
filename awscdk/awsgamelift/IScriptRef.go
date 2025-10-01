package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Script.
// Experimental.
type IScriptRef interface {
	constructs.IConstruct
	// A reference to a Script resource.
	// Experimental.
	ScriptRef() *ScriptReference
}

// The jsii proxy for IScriptRef
type jsiiProxy_IScriptRef struct {
	internal.Type__constructsIConstruct
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

