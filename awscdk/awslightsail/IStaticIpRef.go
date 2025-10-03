package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StaticIp.
// Experimental.
type IStaticIpRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStaticIpRef
type jsiiProxy_IStaticIpRef struct {
	internal.Type__constructsIConstruct
}

