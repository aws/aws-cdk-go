package interfacesawsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLApi.
// Experimental.
type IGraphQLApiRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GraphQLApi resource.
	// Experimental.
	GraphQlApiRef() *GraphQLApiReference
}

// The jsii proxy for IGraphQLApiRef
type jsiiProxy_IGraphQLApiRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGraphQLApiRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGraphQLApiRef) GraphQlApiRef() *GraphQLApiReference {
	var returns *GraphQLApiReference
	_jsii_.Get(
		j,
		"graphQlApiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphQLApiRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphQLApiRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

