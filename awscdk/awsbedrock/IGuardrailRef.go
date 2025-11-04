package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Guardrail.
// Experimental.
type IGuardrailRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Guardrail resource.
	// Experimental.
	GuardrailRef() *GuardrailReference
}

// The jsii proxy for IGuardrailRef
type jsiiProxy_IGuardrailRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGuardrailRef) GuardrailRef() *GuardrailReference {
	var returns *GuardrailReference
	_jsii_.Get(
		j,
		"guardrailRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrailRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrailRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

