package awswafregional

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a XssMatchSet.
// Experimental.
type IXssMatchSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IXssMatchSetRef
type jsiiProxy_IXssMatchSetRef struct {
	internal.Type__constructsIConstruct
}

