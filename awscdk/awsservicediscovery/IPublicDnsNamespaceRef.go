package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicDnsNamespace.
// Experimental.
type IPublicDnsNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPublicDnsNamespaceRef
type jsiiProxy_IPublicDnsNamespaceRef struct {
	internal.Type__constructsIConstruct
}

