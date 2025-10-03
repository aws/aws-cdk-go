package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BatchScramSecret.
// Experimental.
type IBatchScramSecretRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBatchScramSecretRef
type jsiiProxy_IBatchScramSecretRef struct {
	internal.Type__constructsIConstruct
}

