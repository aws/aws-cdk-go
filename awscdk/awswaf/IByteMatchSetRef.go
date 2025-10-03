package awswaf

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ByteMatchSet.
// Experimental.
type IByteMatchSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IByteMatchSetRef
type jsiiProxy_IByteMatchSetRef struct {
	internal.Type__constructsIConstruct
}

