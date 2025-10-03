package awswafregional

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SqlInjectionMatchSet.
// Experimental.
type ISqlInjectionMatchSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISqlInjectionMatchSetRef
type jsiiProxy_ISqlInjectionMatchSetRef struct {
	internal.Type__constructsIConstruct
}

