package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Retriever.
// Experimental.
type IRetrieverRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRetrieverRef
type jsiiProxy_IRetrieverRef struct {
	internal.Type__constructsIConstruct
}

