package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyEndpoint.
// Experimental.
type IDBProxyEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBProxyEndpoint resource.
	// Experimental.
	DbProxyEndpointRef() *DBProxyEndpointReference
}

// The jsii proxy for IDBProxyEndpointRef
type jsiiProxy_IDBProxyEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDBProxyEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IDBProxyEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

