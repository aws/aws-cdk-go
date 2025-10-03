package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OdbNetwork.
// Experimental.
type IOdbNetworkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOdbNetworkRef
type jsiiProxy_IOdbNetworkRef struct {
	internal.Type__constructsIConstruct
}

