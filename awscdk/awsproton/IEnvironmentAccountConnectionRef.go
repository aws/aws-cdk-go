package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentAccountConnection.
// Experimental.
type IEnvironmentAccountConnectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentAccountConnectionRef
type jsiiProxy_IEnvironmentAccountConnectionRef struct {
	internal.Type__constructsIConstruct
}

