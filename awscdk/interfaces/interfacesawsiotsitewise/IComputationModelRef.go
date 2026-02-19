package interfacesawsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputationModel.
// Experimental.
type IComputationModelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ComputationModel resource.
	// Experimental.
	ComputationModelRef() *ComputationModelReference
}

// The jsii proxy for IComputationModelRef
type jsiiProxy_IComputationModelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IComputationModelRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IComputationModelRef) ComputationModelRef() *ComputationModelReference {
	var returns *ComputationModelReference
	_jsii_.Get(
		j,
		"computationModelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputationModelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputationModelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

