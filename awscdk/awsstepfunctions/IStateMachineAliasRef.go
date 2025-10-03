package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachineAlias.
// Experimental.
type IStateMachineAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStateMachineAliasRef
type jsiiProxy_IStateMachineAliasRef struct {
	internal.Type__constructsIConstruct
}

