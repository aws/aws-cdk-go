package awsappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorProfile.
// Experimental.
type IConnectorProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectorProfileRef
type jsiiProxy_IConnectorProfileRef struct {
	internal.Type__constructsIConstruct
}

