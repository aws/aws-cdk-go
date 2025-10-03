package awsopsworkscm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworkscm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Server.
// Experimental.
type IServerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServerRef
type jsiiProxy_IServerRef struct {
	internal.Type__constructsIConstruct
}

