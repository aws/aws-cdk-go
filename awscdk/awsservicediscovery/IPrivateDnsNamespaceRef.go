package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateDnsNamespace.
// Experimental.
type IPrivateDnsNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPrivateDnsNamespaceRef
type jsiiProxy_IPrivateDnsNamespaceRef struct {
	internal.Type__constructsIConstruct
}

