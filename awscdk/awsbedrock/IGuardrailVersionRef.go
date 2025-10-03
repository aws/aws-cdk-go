package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GuardrailVersion.
// Experimental.
type IGuardrailVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGuardrailVersionRef
type jsiiProxy_IGuardrailVersionRef struct {
	internal.Type__constructsIConstruct
}

