package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Integration.
// Experimental.
type IIntegrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIntegrationRef
type jsiiProxy_IIntegrationRef struct {
	internal.Type__constructsIConstruct
}

