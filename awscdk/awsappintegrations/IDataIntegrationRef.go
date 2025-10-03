package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataIntegration.
// Experimental.
type IDataIntegrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataIntegrationRef
type jsiiProxy_IDataIntegrationRef struct {
	internal.Type__constructsIConstruct
}

