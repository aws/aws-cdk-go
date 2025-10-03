package awssimspaceweaver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssimspaceweaver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Simulation.
// Experimental.
type ISimulationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISimulationRef
type jsiiProxy_ISimulationRef struct {
	internal.Type__constructsIConstruct
}

