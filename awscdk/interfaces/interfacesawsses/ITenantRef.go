package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tenant.
// Experimental.
type ITenantRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Tenant resource.
	// Experimental.
	TenantRef() *TenantReference
}

// The jsii proxy for ITenantRef
type jsiiProxy_ITenantRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITenantRef) TenantRef() *TenantReference {
	var returns *TenantReference
	_jsii_.Get(
		j,
		"tenantRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITenantRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITenantRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

