package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailAddress.
// Experimental.
type IEmailAddressRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EmailAddress resource.
	// Experimental.
	EmailAddressRef() *EmailAddressReference
}

// The jsii proxy for IEmailAddressRef
type jsiiProxy_IEmailAddressRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEmailAddressRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEmailAddressRef) EmailAddressRef() *EmailAddressReference {
	var returns *EmailAddressReference
	_jsii_.Get(
		j,
		"emailAddressRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailAddressRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailAddressRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

