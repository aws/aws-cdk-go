package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphQLApi.
// Experimental.
type IGraphQLApiRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGraphQLApiRef
type jsiiProxy_IGraphQLApiRef struct {
	internal.Type__constructsIConstruct
}

