package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventIntegration.
// Experimental.
type IEventIntegrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventIntegrationRef
type jsiiProxy_IEventIntegrationRef struct {
	internal.Type__constructsIConstruct
}

