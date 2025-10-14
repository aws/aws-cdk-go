package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLApi.
// Experimental.
type IGraphQLApiRef interface {
	constructs.IConstruct
	// A reference to a GraphQLApi resource.
	// Experimental.
	GraphQlApiRef() *GraphQLApiReference
}

// The jsii proxy for IGraphQLApiRef
type jsiiProxy_IGraphQLApiRef struct {
	internal.Type__constructsIConstruct
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

