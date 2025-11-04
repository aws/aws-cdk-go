package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLSchema.
// Experimental.
type IGraphQLSchemaRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GraphQLSchema resource.
	// Experimental.
	GraphQlSchemaRef() *GraphQLSchemaReference
}

// The jsii proxy for IGraphQLSchemaRef
type jsiiProxy_IGraphQLSchemaRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGraphQLSchemaRef) GraphQlSchemaRef() *GraphQLSchemaReference {
	var returns *GraphQLSchemaReference
	_jsii_.Get(
		j,
		"graphQlSchemaRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphQLSchemaRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphQLSchemaRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

