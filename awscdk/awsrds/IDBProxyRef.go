package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxy.
// Experimental.
type IDBProxyRef interface {
	constructs.IConstruct
	// A reference to a DBProxy resource.
	// Experimental.
	DbProxyRef() *DBProxyReference
}

// The jsii proxy for IDBProxyRef
type jsiiProxy_IDBProxyRef struct {
	internal.Type__constructsIConstruct
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

