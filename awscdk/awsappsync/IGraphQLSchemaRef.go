package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLSchema.
// Experimental.
type IGraphQLSchemaRef interface {
	constructs.IConstruct
	// A reference to a GraphQLSchema resource.
	// Experimental.
	GraphQlSchemaRef() *GraphQLSchemaReference
}

// The jsii proxy for IGraphQLSchemaRef
type jsiiProxy_IGraphQLSchemaRef struct {
	internal.Type__constructsIConstruct
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

