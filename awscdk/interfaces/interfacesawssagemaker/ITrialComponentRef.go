package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrialComponent.
// Experimental.
type ITrialComponentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrialComponent resource.
	// Experimental.
	TrialComponentRef() *TrialComponentReference
}

// The jsii proxy for ITrialComponentRef
type jsiiProxy_ITrialComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITrialComponentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITrialComponentRef) TrialComponentRef() *TrialComponentReference {
	var returns *TrialComponentReference
	_jsii_.Get(
		j,
		"trialComponentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrialComponentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrialComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

