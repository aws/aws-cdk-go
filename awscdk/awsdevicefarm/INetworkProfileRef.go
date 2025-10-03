package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkProfile.
// Experimental.
type INetworkProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for INetworkProfileRef
type jsiiProxy_INetworkProfileRef struct {
	internal.Type__constructsIConstruct
}

