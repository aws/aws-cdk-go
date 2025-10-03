package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalNetwork.
// Experimental.
type IGlobalNetworkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGlobalNetworkRef
type jsiiProxy_IGlobalNetworkRef struct {
	internal.Type__constructsIConstruct
}

