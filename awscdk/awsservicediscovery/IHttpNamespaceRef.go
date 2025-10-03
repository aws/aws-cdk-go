package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HttpNamespace.
// Experimental.
type IHttpNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHttpNamespaceRef
type jsiiProxy_IHttpNamespaceRef struct {
	internal.Type__constructsIConstruct
}

