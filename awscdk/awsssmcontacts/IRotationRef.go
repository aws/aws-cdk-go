package awsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rotation.
// Experimental.
type IRotationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRotationRef
type jsiiProxy_IRotationRef struct {
	internal.Type__constructsIConstruct
}

