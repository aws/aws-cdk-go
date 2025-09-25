package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyEndpoint.
// Experimental.
type IDBProxyEndpointRef interface {
	constructs.IConstruct
	// A reference to a DBProxyEndpoint resource.
	// Experimental.
	DbProxyEndpointRef() *DBProxyEndpointReference
}

// The jsii proxy for IDBProxyEndpointRef
type jsiiProxy_IDBProxyEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBProxyEndpointRef) DbProxyEndpointRef() *DBProxyEndpointReference {
	var returns *DBProxyEndpointReference
	_jsii_.Get(
		j,
		"dbProxyEndpointRef",
		&returns,
	)
	return returns
}

