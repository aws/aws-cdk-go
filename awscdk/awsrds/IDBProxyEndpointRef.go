package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyEndpoint.
// Experimental.
type IDBProxyEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBProxyEndpoint resource.
	// Experimental.
	DbProxyEndpointRef() *DBProxyEndpointReference
}

// The jsii proxy for IDBProxyEndpointRef
type jsiiProxy_IDBProxyEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDBProxyEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBProxyEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

