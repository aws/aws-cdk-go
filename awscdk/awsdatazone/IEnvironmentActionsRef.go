package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentActions.
// Experimental.
type IEnvironmentActionsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentActionsRef
type jsiiProxy_IEnvironmentActionsRef struct {
	internal.Type__constructsIConstruct
}

