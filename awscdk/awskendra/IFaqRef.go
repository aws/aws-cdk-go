package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Faq.
// Experimental.
type IFaqRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFaqRef
type jsiiProxy_IFaqRef struct {
	internal.Type__constructsIConstruct
}

