package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Gateway.
// Experimental.
type IGatewayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGatewayRef
type jsiiProxy_IGatewayRef struct {
	internal.Type__constructsIConstruct
}

