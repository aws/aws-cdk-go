package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimulationApplication.
// Experimental.
type ISimulationApplicationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SimulationApplication resource.
	// Experimental.
	SimulationApplicationRef() *SimulationApplicationReference
}

// The jsii proxy for ISimulationApplicationRef
type jsiiProxy_ISimulationApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISimulationApplicationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimulationApplicationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

