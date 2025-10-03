package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TypeActivation.
// Experimental.
type ITypeActivationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITypeActivationRef
type jsiiProxy_ITypeActivationRef struct {
	internal.Type__constructsIConstruct
}

