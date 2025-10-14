package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Macro.
// Experimental.
type IMacroRef interface {
	constructs.IConstruct
	// A reference to a Macro resource.
	// Experimental.
	MacroRef() *MacroReference
}

// The jsii proxy for IMacroRef
type jsiiProxy_IMacroRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMacroRef) MacroRef() *MacroReference {
	var returns *MacroReference
	_jsii_.Get(
		j,
		"macroRef",
		&returns,
	)
	return returns
}

