package interfacesawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueryDefinition.
// Experimental.
type IQueryDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QueryDefinition resource.
	// Experimental.
	QueryDefinitionRef() *QueryDefinitionReference
}

// The jsii proxy for IQueryDefinitionRef
type jsiiProxy_IQueryDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQueryDefinitionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQueryDefinitionRef) QueryDefinitionRef() *QueryDefinitionReference {
	var returns *QueryDefinitionReference
	_jsii_.Get(
		j,
		"queryDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueryDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueryDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

