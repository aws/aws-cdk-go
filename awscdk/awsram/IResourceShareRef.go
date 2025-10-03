package awsram

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceShare.
// Experimental.
type IResourceShareRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceShareRef
type jsiiProxy_IResourceShareRef struct {
	internal.Type__constructsIConstruct
}

