package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachineVersion.
// Experimental.
type IStateMachineVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStateMachineVersionRef
type jsiiProxy_IStateMachineVersionRef struct {
	internal.Type__constructsIConstruct
}

