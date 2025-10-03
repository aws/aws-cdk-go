package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeySigningKey.
// Experimental.
type IKeySigningKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKeySigningKeyRef
type jsiiProxy_IKeySigningKeyRef struct {
	internal.Type__constructsIConstruct
}

