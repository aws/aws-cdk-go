package awssimspaceweaver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssimspaceweaver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Simulation.
// Experimental.
type ISimulationRef interface {
	constructs.IConstruct
	// A reference to a Simulation resource.
	// Experimental.
	SimulationRef() *SimulationReference
}

// The jsii proxy for ISimulationRef
type jsiiProxy_ISimulationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISimulationRef) SimulationRef() *SimulationReference {
	var returns *SimulationReference
	_jsii_.Get(
		j,
		"simulationRef",
		&returns,
	)
	return returns
}

