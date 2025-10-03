package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CRL.
// Experimental.
type ICRLRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICRLRef
type jsiiProxy_ICRLRef struct {
	internal.Type__constructsIConstruct
}

