package awspaymentcryptography

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspaymentcryptography/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Key.
// Experimental.
type IKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKeyRef
type jsiiProxy_IKeyRef struct {
	internal.Type__constructsIConstruct
}

