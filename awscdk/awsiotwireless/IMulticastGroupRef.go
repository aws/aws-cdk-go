package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MulticastGroup.
// Experimental.
type IMulticastGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMulticastGroupRef
type jsiiProxy_IMulticastGroupRef struct {
	internal.Type__constructsIConstruct
}

