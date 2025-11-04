package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyTargetGroup.
// Experimental.
type IDBProxyTargetGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBProxyTargetGroup resource.
	// Experimental.
	DbProxyTargetGroupRef() *DBProxyTargetGroupReference
}

// The jsii proxy for IDBProxyTargetGroupRef
type jsiiProxy_IDBProxyTargetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDBProxyTargetGroupRef) DbProxyTargetGroupRef() *DBProxyTargetGroupReference {
	var returns *DBProxyTargetGroupReference
	_jsii_.Get(
		j,
		"dbProxyTargetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBProxyTargetGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBProxyTargetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

