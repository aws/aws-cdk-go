package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GuardrailVersion.
// Experimental.
type IGuardrailVersionRef interface {
	constructs.IConstruct
	// A reference to a GuardrailVersion resource.
	// Experimental.
	GuardrailVersionRef() *GuardrailVersionReference
}

// The jsii proxy for IGuardrailVersionRef
type jsiiProxy_IGuardrailVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGuardrailVersionRef) GuardrailVersionRef() *GuardrailVersionReference {
	var returns *GuardrailVersionReference
	_jsii_.Get(
		j,
		"guardrailVersionRef",
		&returns,
	)
	return returns
}

