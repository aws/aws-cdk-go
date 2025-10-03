package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Resource.
// Experimental.
type IResourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceRef
type jsiiProxy_IResourceRef struct {
	internal.Type__constructsIConstruct
}

