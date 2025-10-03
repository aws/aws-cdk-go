package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyEndpoint.
// Experimental.
type IDBProxyEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBProxyEndpointRef
type jsiiProxy_IDBProxyEndpointRef struct {
	internal.Type__constructsIConstruct
}

