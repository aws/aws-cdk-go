package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxy.
// Experimental.
type IDBProxyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBProxy resource.
	// Experimental.
	DbProxyRef() *DBProxyReference
}

// The jsii proxy for IDBProxyRef
type jsiiProxy_IDBProxyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDBProxyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

