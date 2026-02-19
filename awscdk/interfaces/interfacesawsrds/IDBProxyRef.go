package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxy.
// Experimental.
type IDBProxyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBProxy resource.
	// Experimental.
	DbProxyRef() *DBProxyReference
}

// The jsii proxy for IDBProxyRef
type jsiiProxy_IDBProxyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDBProxyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDBProxyRef) DbProxyRef() *DBProxyReference {
	var returns *DBProxyReference
	_jsii_.Get(
		j,
		"dbProxyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBProxyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBProxyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

