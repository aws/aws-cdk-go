package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DelegatedAdmin.
// Experimental.
type IDelegatedAdminRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DelegatedAdmin resource.
	// Experimental.
	DelegatedAdminRef() *DelegatedAdminReference
}

// The jsii proxy for IDelegatedAdminRef
type jsiiProxy_IDelegatedAdminRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDelegatedAdminRef) DelegatedAdminRef() *DelegatedAdminReference {
	var returns *DelegatedAdminReference
	_jsii_.Get(
		j,
		"delegatedAdminRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegatedAdminRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegatedAdminRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

