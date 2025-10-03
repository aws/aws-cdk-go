package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Guardrail.
// Experimental.
type IGuardrailRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGuardrailRef
type jsiiProxy_IGuardrailRef struct {
	internal.Type__constructsIConstruct
}

