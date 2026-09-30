package interfacesawssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistrationAttachment.
// Experimental.
type IRegistrationAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RegistrationAttachment resource.
	// Experimental.
	RegistrationAttachmentRef() *RegistrationAttachmentReference
}

// The jsii proxy for IRegistrationAttachmentRef
type jsiiProxy_IRegistrationAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRegistrationAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRegistrationAttachmentRef) RegistrationAttachmentRef() *RegistrationAttachmentReference {
	var returns *RegistrationAttachmentReference
	_jsii_.Get(
		j,
		"registrationAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistrationAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistrationAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

