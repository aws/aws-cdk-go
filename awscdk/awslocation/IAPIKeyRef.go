package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APIKey.
// Experimental.
type IAPIKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAPIKeyRef
type jsiiProxy_IAPIKeyRef struct {
	internal.Type__constructsIConstruct
}

