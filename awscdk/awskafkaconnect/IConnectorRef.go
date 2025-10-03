package awskafkaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connector.
// Experimental.
type IConnectorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectorRef
type jsiiProxy_IConnectorRef struct {
	internal.Type__constructsIConstruct
}

