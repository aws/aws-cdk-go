package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolSessionConfigurationAttachment.
// Experimental.
type IUserPoolSessionConfigurationAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserPoolSessionConfigurationAttachment resource.
	// Experimental.
	UserPoolSessionConfigurationAttachmentRef() *UserPoolSessionConfigurationAttachmentReference
}

// The jsii proxy for IUserPoolSessionConfigurationAttachmentRef
type jsiiProxy_IUserPoolSessionConfigurationAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserPoolSessionConfigurationAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserPoolSessionConfigurationAttachmentRef) UserPoolSessionConfigurationAttachmentRef() *UserPoolSessionConfigurationAttachmentReference {
	var returns *UserPoolSessionConfigurationAttachmentReference
	_jsii_.Get(
		j,
		"userPoolSessionConfigurationAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolSessionConfigurationAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolSessionConfigurationAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

