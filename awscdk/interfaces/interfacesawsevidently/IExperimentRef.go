package interfacesawsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Experiment.
// Experimental.
type IExperimentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Experiment resource.
	// Experimental.
	ExperimentRef() *ExperimentReference
}

// The jsii proxy for IExperimentRef
type jsiiProxy_IExperimentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IExperimentRef) ExperimentRef() *ExperimentReference {
	var returns *ExperimentReference
	_jsii_.Get(
		j,
		"experimentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

