package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Collection.
// Experimental.
type ICollectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICollectionRef
type jsiiProxy_ICollectionRef struct {
	internal.Type__constructsIConstruct
}

