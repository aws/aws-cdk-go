package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSecurityGroupIngress.
// Experimental.
type IDBSecurityGroupIngressRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBSecurityGroupIngress resource.
	// Experimental.
	DbSecurityGroupIngressRef() *DBSecurityGroupIngressReference
}

// The jsii proxy for IDBSecurityGroupIngressRef
type jsiiProxy_IDBSecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDBSecurityGroupIngressRef) DbSecurityGroupIngressRef() *DBSecurityGroupIngressReference {
	var returns *DBSecurityGroupIngressReference
	_jsii_.Get(
		j,
		"dbSecurityGroupIngressRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSecurityGroupIngressRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSecurityGroupIngressRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

