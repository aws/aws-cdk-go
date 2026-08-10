package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExperimentTrialComponent.
// Experimental.
type IExperimentTrialComponentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ExperimentTrialComponent resource.
	// Experimental.
	ExperimentTrialComponentRef() *ExperimentTrialComponentReference
}

// The jsii proxy for IExperimentTrialComponentRef
type jsiiProxy_IExperimentTrialComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IExperimentTrialComponentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IExperimentTrialComponentRef) ExperimentTrialComponentRef() *ExperimentTrialComponentReference {
	var returns *ExperimentTrialComponentReference
	_jsii_.Get(
		j,
		"experimentTrialComponentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentTrialComponentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentTrialComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

