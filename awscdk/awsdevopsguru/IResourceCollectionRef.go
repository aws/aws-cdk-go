package awsdevopsguru

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevopsguru/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceCollection.
// Experimental.
type IResourceCollectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceCollectionRef
type jsiiProxy_IResourceCollectionRef struct {
	internal.Type__constructsIConstruct
}

