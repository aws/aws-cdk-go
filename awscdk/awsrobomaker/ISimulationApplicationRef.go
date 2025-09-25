package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimulationApplication.
// Experimental.
type ISimulationApplicationRef interface {
	constructs.IConstruct
	// A reference to a SimulationApplication resource.
	// Experimental.
	SimulationApplicationRef() *SimulationApplicationReference
}

// The jsii proxy for ISimulationApplicationRef
type jsiiProxy_ISimulationApplicationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISimulationApplicationRef) SimulationApplicationRef() *SimulationApplicationReference {
	var returns *SimulationApplicationReference
	_jsii_.Get(
		j,
		"simulationApplicationRef",
		&returns,
	)
	return returns
}

