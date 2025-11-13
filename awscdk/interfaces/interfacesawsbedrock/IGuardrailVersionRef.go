package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GuardrailVersion.
// Experimental.
type IGuardrailVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GuardrailVersion resource.
	// Experimental.
	GuardrailVersionRef() *GuardrailVersionReference
}

// The jsii proxy for IGuardrailVersionRef
type jsiiProxy_IGuardrailVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IGuardrailVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrailVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

