package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Transformer.
// Experimental.
type ITransformerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITransformerRef
type jsiiProxy_ITransformerRef struct {
	internal.Type__constructsIConstruct
}

