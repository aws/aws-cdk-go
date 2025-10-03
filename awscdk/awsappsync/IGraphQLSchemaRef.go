package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLSchema.
// Experimental.
type IGraphQLSchemaRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGraphQLSchemaRef
type jsiiProxy_IGraphQLSchemaRef struct {
	internal.Type__constructsIConstruct
}

