package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNGateway.
// Experimental.
type IVPNGatewayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVPNGatewayRef
type jsiiProxy_IVPNGatewayRef struct {
	internal.Type__constructsIConstruct
}

