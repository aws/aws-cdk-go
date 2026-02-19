package interfacesawsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Retriever.
// Experimental.
type IRetrieverRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Retriever resource.
	// Experimental.
	RetrieverRef() *RetrieverReference
}

// The jsii proxy for IRetrieverRef
type jsiiProxy_IRetrieverRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRetrieverRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRetrieverRef) RetrieverRef() *RetrieverReference {
	var returns *RetrieverReference
	_jsii_.Get(
		j,
		"retrieverRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrieverRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrieverRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

