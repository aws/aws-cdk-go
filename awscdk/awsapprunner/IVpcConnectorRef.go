package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcConnector.
// Experimental.
type IVpcConnectorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcConnectorRef
type jsiiProxy_IVpcConnectorRef struct {
	internal.Type__constructsIConstruct
}

