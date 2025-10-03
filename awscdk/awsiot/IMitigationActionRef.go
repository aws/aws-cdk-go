package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MitigationAction.
// Experimental.
type IMitigationActionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMitigationActionRef
type jsiiProxy_IMitigationActionRef struct {
	internal.Type__constructsIConstruct
}

