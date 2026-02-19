package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailIdentity.
// Experimental.
type IEmailIdentityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EmailIdentity resource.
	// Experimental.
	EmailIdentityRef() *EmailIdentityReference
}

// The jsii proxy for IEmailIdentityRef
type jsiiProxy_IEmailIdentityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEmailIdentityRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEmailIdentityRef) EmailIdentityRef() *EmailIdentityReference {
	var returns *EmailIdentityReference
	_jsii_.Get(
		j,
		"emailIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

