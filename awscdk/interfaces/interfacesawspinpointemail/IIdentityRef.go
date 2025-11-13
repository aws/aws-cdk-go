package interfacesawspinpointemail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Identity.
// Experimental.
type IIdentityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Identity resource.
	// Experimental.
	IdentityRef() *IdentityReference
}

// The jsii proxy for IIdentityRef
type jsiiProxy_IIdentityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIdentityRef) IdentityRef() *IdentityReference {
	var returns *IdentityReference
	_jsii_.Get(
		j,
		"identityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

