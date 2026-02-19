package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSecurityGroupIngress.
// Experimental.
type IDBSecurityGroupIngressRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBSecurityGroupIngress resource.
	// Experimental.
	DbSecurityGroupIngressRef() *DBSecurityGroupIngressReference
}

// The jsii proxy for IDBSecurityGroupIngressRef
type jsiiProxy_IDBSecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDBSecurityGroupIngressRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDBSecurityGroupIngressRef) DbSecurityGroupIngressRef() *DBSecurityGroupIngressReference {
	var returns *DBSecurityGroupIngressReference
	_jsii_.Get(
		j,
		"dbSecurityGroupIngressRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSecurityGroupIngressRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

