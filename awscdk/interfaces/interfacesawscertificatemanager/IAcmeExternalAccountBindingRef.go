package interfacesawscertificatemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscertificatemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AcmeExternalAccountBinding.
// Experimental.
type IAcmeExternalAccountBindingRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AcmeExternalAccountBinding resource.
	// Experimental.
	AcmeExternalAccountBindingRef() *AcmeExternalAccountBindingReference
}

// The jsii proxy for IAcmeExternalAccountBindingRef
type jsiiProxy_IAcmeExternalAccountBindingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAcmeExternalAccountBindingRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAcmeExternalAccountBindingRef) AcmeExternalAccountBindingRef() *AcmeExternalAccountBindingReference {
	var returns *AcmeExternalAccountBindingReference
	_jsii_.Get(
		j,
		"acmeExternalAccountBindingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcmeExternalAccountBindingRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcmeExternalAccountBindingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

