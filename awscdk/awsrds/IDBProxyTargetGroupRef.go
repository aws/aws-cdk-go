package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyTargetGroup.
// Experimental.
type IDBProxyTargetGroupRef interface {
	constructs.IConstruct
	// A reference to a DBProxyTargetGroup resource.
	// Experimental.
	DbProxyTargetGroupRef() *DBProxyTargetGroupReference
}

// The jsii proxy for IDBProxyTargetGroupRef
type jsiiProxy_IDBProxyTargetGroupRef struct {
	internal.Type__constructsIConstruct
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

