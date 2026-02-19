package interfacesawssimspaceweaver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssimspaceweaver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Simulation.
// Experimental.
type ISimulationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Simulation resource.
	// Experimental.
	SimulationRef() *SimulationReference
}

// The jsii proxy for ISimulationRef
type jsiiProxy_ISimulationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISimulationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_ISimulationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimulationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

