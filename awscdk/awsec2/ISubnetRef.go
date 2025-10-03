package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subnet.
// Experimental.
type ISubnetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubnetRef
type jsiiProxy_ISubnetRef struct {
	internal.Type__constructsIConstruct
}

