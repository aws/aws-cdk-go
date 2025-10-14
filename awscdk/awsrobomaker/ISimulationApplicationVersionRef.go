package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimulationApplicationVersion.
// Experimental.
type ISimulationApplicationVersionRef interface {
	constructs.IConstruct
	// A reference to a SimulationApplicationVersion resource.
	// Experimental.
	SimulationApplicationVersionRef() *SimulationApplicationVersionReference
}

// The jsii proxy for ISimulationApplicationVersionRef
type jsiiProxy_ISimulationApplicationVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISimulationApplicationVersionRef) SimulationApplicationVersionRef() *SimulationApplicationVersionReference {
	var returns *SimulationApplicationVersionReference
	_jsii_.Get(
		j,
		"simulationApplicationVersionRef",
		&returns,
	)
	return returns
}

